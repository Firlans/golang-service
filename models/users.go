package models

import (
	"database/sql"
	"go-api/entities"
)

func AddUser(db *sql.DB, name string, email string, password string) error {
	_, err := db.Exec("INSERT INTO USERS(name, email, password) VALUES($1,$2,$3)", name, email, password)
	return err
}
func GetUserById(db *sql.DB, id int) (*entities.User, error) {
	row := db.QueryRow("SELECT id, name, password,email, created_at, updated_at password FROM users WHERE id=$1", id)

	var user entities.User
	err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func Update(db *sql.DB) {
	// _, err := db.Exec('INSERT INTO USERS(username, email, password) VALUES($1,$2)', username, email, password)
	// return err
}
func Delete(db *sql.DB) {
	// _, err := db.Exec('INSERT INTO USERS(username, email, password) VALUES($1,$2)', username, email, password)
	// return err
}
