package repository

import (
	"database/sql"
	"fmt"

	"klui/clean-arch/internal/utils"

	"klui/clean-arch/internal/models"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(user models.User) error {
	var exists bool

	if err := r.db.QueryRow(
		`SELECT EXISTS(SELECT 1 FROM tb_users WHERE email = $1)`, user.Email,
	).Scan(&exists); err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("email already exists")
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	_, err = r.db.Exec(
		`
			INSERT INTO tb_users (
				email, 
				password, 
				first_name, 
				last_name,
				address
			) 
			VALUES ($1, $2, $3, $4, $5)
		`, user.Email, hashedPassword, user.FirstName, user.LastName, user.Address,
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) GetUserByEmail(email string, password string) (*models.User, error) {
	var userData models.User
	if err := r.db.QueryRow(
		`
			SELECT 
				user_id,
				email,
				password,
				first_name,
				last_name,
				address,
				is_seller
			FROM tb_users
			WHERE 1=1
				AND email = $1
		`, email,
	).Scan(&userData.UserId, &userData.Email, &userData.Password, &userData.FirstName, &userData.LastName, &userData.Address, &userData.IsSeller); err != nil {
		return &models.User{}, fmt.Errorf("user not found")
	}
	if err := utils.CheckPasswordHash(password, userData.Password); err != nil {
		return &models.User{}, err
	}
	return &userData, nil
}
