package repository

import (
	"database/sql"
	"fmt"
	"product-service/internal/models"

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
			INSERT INTO tb_product(
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
		product.ProductImage,
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

func (r *ProductRepository) DeleteProductById(productID string) error {
	if _, err := r.db.Exec(
		`
			DELETE FROM tb_products
			WHERE product_id = $1
		`, productID,
	); err != nil {
		return err
	}
	return nil
}

func (r *ProductRepository) InactivateProductById(productID string) error {
	if _, err := r.db.Exec(
		`
			UPDATE tb_products
			SET is_active = false
			WHERE product_id = $1
		`, productID,
	); err != nil {
		return err
	}
	return nil
}
