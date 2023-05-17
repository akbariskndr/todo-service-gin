package main

import (
	"akbariskndr/todo-service-gin/database"
	"akbariskndr/todo-service-gin/routes"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	config := database.Config{
		Host:     fmt.Sprintf("%s:%s", os.Getenv("DB_HOST"), os.Getenv("DB_PORT")),
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		DB:       os.Getenv("DB_DATABASE"),
	}

	if err := database.Connect(database.GetConnectionString(config)); err != nil {
		panic(err.Error())
	}

	router := routes.InitRoutes()
	router.Run("localhost:8080")
}
