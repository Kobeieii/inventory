package main

import (
	"fmt"
	"inventory/internal/app/adapter/repository"
	"inventory/internal/app/adapter/handler"
	"inventory/internal/app/core/service"
	"inventory/internal/infrastructure/database"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {

	app := fiber.New()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable TimeZone=Asia/Bangkok",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := database.ConnectDB(dsn)
	if err != nil {
		panic(err)
	}

	productRepository := repository.NewGormProductRepository(db)
	productService := service.NewProductService(productRepository)
	productHandler := handler.NewHttpProductHandler(productService)

	api := app.Group("/api")
	productHandler.RegisterRoutes(api)

	app.Listen(":3000")
}
