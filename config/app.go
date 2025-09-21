package config

import (
	"go-api/utils"
	"log"
	"net/http"
)

func RunningApp() {
	appName := utils.GetEnv("APP_NAME", "go-app")
	appHost := utils.GetEnv("APP_HOST", "localhost")
	appPort := utils.GetEnv("APP_PORT", "8000")
	appSchema := utils.GetEnv("APP_SCHEMA", "http")
	log.Printf("%s running at %s://%s:%s\n", appName, appSchema, appHost, appPort)
	log.Fatal(http.ListenAndServe(":"+appPort, nil))
}
