package main

import (
	"fmt"
	"os"
	"product-service/config"
	"product-service/internal/handler"
	"product-service/internal/repository"
	"product-service/internal/service"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load("../config/.env"); err != nil {
		fmt.Println(err)
	}

	PORT := os.Getenv("PORT")
	fmt.Println(PORT)
	// import db
	db, err := config.Connection()
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// initial dependencies
	productRepository := repository.NewProductRepository(db)
	userService := service.NewProductService(productRepository)

	// run the application
	app := fiber.New()
	handler.SetupProductRoute(app, userService)

	app.Listen(":" + PORT)
}
