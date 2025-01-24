package service

import (
	"errors"
	"klui/clean-arch/internal/models"
	"klui/clean-arch/internal/repository"
	"klui/clean-arch/internal/utils"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserSerivce(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user models.User) error {
	if user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" || user.Address == "" {
		return errors.New("all fields must be filled")
	}
	if err := s.repo.CreateUser(user); err != nil {
		return err
	}
	return nil
}

func (s *UserService) Login(user models.User) (string, error) {

	if user.Email == "" || user.Password == "" {
		return "", errors.New("all fields must be filled")
	}

	userData, err := s.repo.GetUserByEmail(user.Email, user.Password)
	if err != nil {
		return "", err
	}

	token, err := utils.JwtSign(*userData)
	if err != nil {
		return "", err
	}

	return token, nil
}
