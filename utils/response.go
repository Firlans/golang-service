package utils

import (
	"encoding/json"
	"go-api/entities"
	"log"
	"net/http"
	"os"
)

func Starting(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	response := entities.Response{
		Status:  "success",
		Message: "Server is running",
		Data: map[string]string{
			"version": "1.0.0",
		},
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(response)
}

func ErrorResponse(message string, err string) entities.Response {
	if debug, ok := os.LookupEnv("DEBUG"); ok && debug == "true" {
		log.Println(err)
	}
	return entities.Response{
		Status:  "fail",
		Message: message,
	}
}

func SuccessResponse(message string, detail interface{}) entities.Response {
	return entities.Response{
		Status:  "success",
		Message: message,
		Data:    detail,
	}
}

func NotFoundResponse(message string) entities.Response {
	return entities.Response{
		Status:  "fail",
		Message: message,
	}
}

func AlreadyExistsResponse(object string) entities.Response {
	return entities.Response{
		Status:  "fail",
		Message: object + " already exists",
	}
}

func InvalidRequestResponse(message string, detail string) entities.Response {
	return entities.Response{
		Status:  "fail",
		Message: message,
		Data: map[string]string{
			"error": detail,
		},
	}
}

func RespondJSON(w http.ResponseWriter, statusCode int, response entities.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
