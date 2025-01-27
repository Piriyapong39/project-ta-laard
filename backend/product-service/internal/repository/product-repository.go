package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"product-service/internal/models"

	"strconv"
	"strings"

	"github.com/lib/pq"
)

type ProductRepository struct {
	db *sql.DB
}

func NewProductRepository(db *sql.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) CreateProduct(product models.Product) error {

	if _, err := r.db.Exec(
		`
			INSERT INTO tb_products(
				product_id,
				product_name,
				product_desc,
				product_price,
				product_stock,
				main_cate,
				sub_cate,
				product_image,
				user_id
			)
				VAlUES($1, $2, $3, $4, $5, $6, $7, $8, $9)
		`, product.ProductID,
		product.ProductName,
		product.Description,
		product.Price,
		product.Stock,
		product.MainCategory,
		product.SubCategory,
		pq.Array(product.ProductImage),
		product.UserID,
	); err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) GetProducts(productFilter models.ProductFilter, page int, userId uint) ([]models.ResponseProduct, error) {
	limit := 50
	offset := (page - 1) * limit

	query := `
        SELECT
            product_id,
            product_name,
            product_desc,
            product_price,
            product_stock,
            main_cate,
            sub_cate,
            product_image
        FROM tb_products
        WHERE user_id = $1
    `

	params := []interface{}{userId, limit, offset}

	if productFilter.ProductID != "" {
		query += " AND product_id LIKE $4"
		params = append(params, "%"+productFilter.ProductID+"%")
	}
	if productFilter.ProductName != "" {
		query += fmt.Sprintf(" AND product_name LIKE $%d", len(params)+1)
		params = append(params, "%"+productFilter.ProductName+"%")
	}
	if productFilter.MainCategory != 0 {
		query += fmt.Sprintf(" AND main_cate = $%d", len(params)+1)
		params = append(params, productFilter.MainCategory)
	}
	if productFilter.SubCategory != 0 {
		query += fmt.Sprintf(" AND sub_cate = $%d", len(params)+1)
		params = append(params, productFilter.SubCategory)
	}

	query += " LIMIT $2 OFFSET $3"

	rows, err := r.db.Query(query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []models.ResponseProduct
	for rows.Next() {
		var product models.ResponseProduct
		err := rows.Scan(
			&product.ProductID,
			&product.ProductName,
			&product.Description,
			&product.Price,
			&product.Stock,
			&product.MainCategory,
			&product.SubCategory,
			pq.Array(&product.ProductImage),
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (r *ProductRepository) DeleteProductById(productID string, userId uint) error {
	results, err := r.db.Exec(
		`
			DELETE FROM tb_products
			WHERE 1=1
				AND product_id = $1
				AND user_id = $2
		`, productID, userId,
	)
	if err != nil {
		return err
	}
	rowsAffected, _ := results.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no affected row")
	}
	return nil
}

func (r *ProductRepository) InactivateProductById(productID string, userId uint) error {
	results, err := r.db.Exec(
		`
			UPDATE tb_products
			SET is_active = false
			WHERE 1=1
				AND product_id = $1
				AND user_id = $2
		`, productID, userId,
	)
	if err != nil {
		return err
	}

	affectedRows, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return errors.New("no affected row")
	}
	return nil
}

func (r *ProductRepository) ActivateProduct(productId string, userId uint) error {

	results, err := r.db.Exec(
		`
			UPDATE tb_products
			SET is_active = true
			WHERE 1=1
				AND product_id = $1
				AND user_id = $2
		`, productId, userId,
	)
	if err != nil {
		return err
	}
	affectedRows, err := results.RowsAffected()
	if err != nil {
		return err
	}
	if affectedRows == 0 {
		return errors.New("no affected row")
	}
	return nil
}

func (r *ProductRepository) UpdateProduct(product models.Product, userId uint) error {
	var setStatements []string
	params := []interface{}{userId}

	if product.ProductName != "" {
		params = append(params, product.ProductName)
		setStatements = append(setStatements, "product_name = $"+strconv.Itoa(len(params)))
	}

	if product.Description != "" {
		params = append(params, product.Description)
		setStatements = append(setStatements, "description = $"+strconv.Itoa(len(params)))
	}

	if product.Price != 0 {
		params = append(params, product.Price)
		setStatements = append(setStatements, "price = $"+strconv.Itoa(len(params)))
	}

	if product.Stock != 0 {
		params = append(params, product.Stock)
		setStatements = append(setStatements, "stock = $"+strconv.Itoa(len(params)))
	}

	if len(setStatements) == 0 {
		return errors.New("insert data to update")
	}

	query := `
        UPDATE tb_products
        SET ` + strings.Join(setStatements, ", ") + `
        WHERE user_id = $1
        RETURNING id`

	return r.db.QueryRow(query, params...).Scan(&product.ProductID)
}
