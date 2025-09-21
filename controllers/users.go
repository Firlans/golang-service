package controllers

import (
	"database/sql"
	"encoding/json"
	"go-api/entities"
	"go-api/models"
	"go-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		var req entities.CreateUserRequest

		err := json.NewDecoder(r.Body).Decode(&req)
		if err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, utils.ErrorResponse("Server Error", err.Error()))
			return
		}

		if req.Name == "" || req.Email == "" || req.Password == "" {
			utils.RespondJSON(w, http.StatusBadRequest, utils.InvalidRequestResponse("Invalid input", "name, email, and password are required"))
			return
		}

		id, err := models.AddUser(db, req.Name, req.Email, req.Password)
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value") {
				utils.RespondJSON(w, http.StatusInternalServerError, utils.AlreadyExistsResponse("user"))
			} else {
				utils.RespondJSON(w, http.StatusInternalServerError, utils.ErrorResponse("Server error", err.Error()))
			}
			return
		}

		utils.RespondJSON(w, http.StatusCreated, entities.Response{
			Status:  "success",
			Message: "user created successfully",
			Data: struct {
				ID int `json:"id"`
			}{
				ID: id,
			},
		})
	}
}
func UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idStr, ok := vars["id"]
		if !ok {
			utils.RespondJSON(w, http.StatusBadRequest, utils.InvalidRequestResponse("Invalid input", "User ID is required"))
			return
		}

		id, err := strconv.Atoi(idStr)
		if err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, utils.InvalidRequestResponse("Invalid ID format", "ID must be a number"))
			return
		}

		var updateData map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, utils.ErrorResponse("Invalid JSON format", err.Error()))
			return
		}

		// Panggil model untuk update user dan dapatkan data yang diperbarui
		updatedUser, err := models.UpdateUser(db, id, updateData)
		if err != nil {
			if strings.Contains(err.Error(), sql.ErrNoRows.Error()) {
				utils.RespondJSON(w, http.StatusNotFound, utils.ErrorResponse("User not found", "No user found with the given ID"))
			} else if strings.Contains(err.Error(), "duplicate key value") {
				utils.RespondJSON(w, http.StatusConflict, utils.ErrorResponse("Data conflict", "Email already in use"))
			} else {
				utils.RespondJSON(w, http.StatusInternalServerError, utils.ErrorResponse("Server error", err.Error()))
			}
			return
		}

		// Kirim respons sukses dengan data user yang diperbarui
		utils.RespondJSON(w, http.StatusOK, utils.SuccessResponse("User updated successfully", updatedUser))
	}
}
func DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ids := vars["id"]
		id, err := strconv.Atoi(ids)
		if err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, utils.InvalidRequestResponse("Invalid input", "id must be numeric"))
			return
		}

		if err := models.DeleteUser(db, id); err != nil {
			utils.RespondJSON(w, http.StatusInternalServerError, utils.ErrorResponse("Server error", err.Error()))
			return
		}

		utils.RespondJSON(w, http.StatusOK, utils.SuccessResponse("User deleted successfully", id))
	}
}

func GetAllUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := models.GetAllUser(db)

		if err != nil {
			utils.RespondJSON(w, http.StatusInternalServerError, utils.ErrorResponse("Server error", err.Error()))
			return
		}

		utils.RespondJSON(w, http.StatusOK, utils.SuccessResponse("User retrieved successfully", users))
	}
}
func GetUserById(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		ids := vars["id"]
		id, err := strconv.Atoi(ids)
		if err != nil {
			utils.RespondJSON(w, http.StatusBadRequest, utils.InvalidRequestResponse("Invalid input", "id must be numeric"))
			return
		}
		user, err := models.GetUserById(db, id)
		if err != nil {
			utils.RespondJSON(w, http.StatusInternalServerError, utils.ErrorResponse("Server error", err.Error()))
			return
		}

		if user == nil {
			utils.RespondJSON(w, http.StatusNotFound, utils.NotFoundResponse("user", ids))
			return
		}

		utils.RespondJSON(w, http.StatusOK, utils.SuccessResponse("User retrieved successfully", user))
	}
}
