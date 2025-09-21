package routes

import (
	"database/sql"
	"go-api/controllers"
	"go-api/utils"
	"net/http"
)

func RegisterUserRoutes(db *sql.DB) {
	http.HandleFunc("/", utils.Starting)
	http.HandleFunc("/users", controllers.GetUser(db))
	http.HandleFunc("/user/create", controllers.CreateUser(db))
}
