package config

import (
	"database/sql"
	"fmt"
	"go-api/utils"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	host := utils.GetEnv("DB_HOST", "localhost")
	port := utils.GetEnv("DB_PORT", "5432")
	user := utils.GetEnv("DB_USER", "postgres")
	pass := utils.GetEnv("DB_PASSWORD", "")
	name := utils.GetEnv("DB_NAME", "postgres")
	sslmode := utils.GetEnv("DB_SSL]", "disable")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		host, port, user, pass, name, sslmode,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed to open database:", err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("connection to database was failed:", err)
	}

	log.Println("Connection to PostgreSQL is success!")
	return db
}
