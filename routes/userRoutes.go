package routes

import (
	"database/sql"
	"go-api/controllers"
	"go-api/utils"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(db *sql.DB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", utils.Starting).Methods("GET")
	r.HandleFunc("/users", controllers.GetAllUser(db)).Methods("GET")
	r.HandleFunc("/user/{id}", controllers.GetUserById(db)).Methods("GET")
	r.HandleFunc("/user", controllers.CreateUser(db)).Methods("POST")
	r.HandleFunc("/user/{id}", controllers.UpdateUser(db)).Methods("PUT")
	r.HandleFunc("/user/{id}", controllers.DeleteUser(db)).Methods("DELETE")

	r.HandleFunc("/auth/login", controllers.Login(db)).Methods("POST")
	return r
}
