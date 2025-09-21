package controllers

import (
	"database/sql"
	"encoding/json"
	"go-api/entities"
	"go-api/models"
	"go-api/utils"
	"log"
	"net/http"
	"strconv"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, utils.ErrorResponse("Invalid form data", err.Error()))
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")
		password := r.FormValue("password")

		if name == "" || email == "" || password == "" {
			utils.RespondJSON(w, http.StatusBadRequest, utils.InvalidRequestResponse("Invalid input", "name, email, and password are required"))
			return
		}

		err := models.AddUser(db, name, email, password)
		if err != nil {
			utils.RespondJSON(w, http.StatusInternalServerError, utils.ErrorResponse("Server error", err.Error()))
			return
		}

		utils.RespondJSON(w, http.StatusCreated, entities.Response{
			Status:  "success",
			Message: "user created successfully",
			Data: map[string]string{
				"name":  name,
				"email": email,
			},
		})
	}
}
func UpdateUser() {

}
func DeleteUser() {

}
func GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ids := r.URL.Query().Get("id")
		id, err := strconv.Atoi(ids)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		user, err := models.GetUserById(db, id)

		if err != nil {
			log.Println("error database:", err)
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
