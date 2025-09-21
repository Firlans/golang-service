package main

import (
	"go-api/config"
	"go-api/routes"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db := config.ConnectDB()
	defer db.Close()

	routes.RegisterUserRoutes(db)
	config.RunningApp()
}
