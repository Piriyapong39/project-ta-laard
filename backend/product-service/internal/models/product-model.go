package models

type Product struct {
	ProductID    string   `json:"product_id"`
	ProductName  string   `json:"product_name"`
	Price        float64  `json:"price"`
	Stock        int      `json:"stock"`
	Description  string   `json:"description"`
	MainCategory int      `json:"main_category"`
	SubCategory  int      `json:"sub_category"`
	ProductImage []string `json:"-"`
	UserID       int      `json:"user_id"`
}

type ResponseProduct struct {
	ProductID    string   `json:"product_id"`
	ProductName  string   `json:"productName"`
	Price        float64  `json:"price"`
	Stock        int      `json:"stock"`
	Description  string   `json:"description"`
	MainCategory int      `json:"mainCategory"`
	SubCategory  int      `json:"subCategory"`
	ProductImage []string `json:"product_image"`
}

type ProductFilter struct {
	ProductID    string `json:"productId"`
	ProductName  string `json:"productName"`
	MainCategory int    `json:"mainCategory"`
	SubCategory  int    `json:"subCategory"`
	Page         int    `json:"page"`
}
