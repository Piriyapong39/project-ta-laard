package main

import (
	"klui/clean-arch/config"
	"klui/clean-arch/internal/handler"
	"klui/clean-arch/internal/repository"
	"klui/clean-arch/internal/service"
	"os"

	"github.com/joho/godotenv"

	"github.com/gofiber/fiber/v2"
)

func main() {

	// impport config
	if err := godotenv.Load("../config/.env"); err != nil {
		panic(err)
	}
	PORT := os.Getenv("PORT")

	// import db
	db, err := config.Connection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// initialize dependecies
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserSerivce(userRepository)

	// run the application
	app := fiber.New()
	handler.SetUpUserRoutes(app, userService)

	app.Listen(":" + PORT)
}
