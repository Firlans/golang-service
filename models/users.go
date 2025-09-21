package models

import (
	"database/sql"
	"fmt"
	"go-api/entities"
	"strings"
)

func AddUser(db *sql.DB, name string, email string, password string) (int, error) {
	var id int

	query := `INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id`

	err := db.QueryRow(query, name, email, password).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("could not insert user: %w", err)
	}

	return id, nil
}
func GetAllUser(db *sql.DB) ([]*entities.User, error) {
	rows, err := db.Query("SELECT id, name, email, created_at, updated_at FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entities.User

	for rows.Next() {
		var user entities.User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return []*entities.User{}, nil
	}

	return users, nil
}
func GetUserById(db *sql.DB, id int) (*entities.User, error) {
	var user entities.User
	row := db.QueryRow("SELECT id, name, email, created_at, updated_at FROM users where id = $1", id)

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// UpdateUser memperbarui user dan mengembalikan data yang diperbarui.
func UpdateUser(db *sql.DB, id int, data map[string]interface{}) (*entities.User, error) {
	var updates []string
	var args []interface{}

	argCounter := 1

	if name, ok := data["name"]; ok {
		updates = append(updates, fmt.Sprintf("name = $%d", argCounter))
		args = append(args, name)
		argCounter++
	}

	if email, ok := data["email"]; ok {
		updates = append(updates, fmt.Sprintf("email = $%d", argCounter))
		args = append(args, email)
		argCounter++
	}

	if password, ok := data["password"]; ok {
		updates = append(updates, fmt.Sprintf("password = $%d", argCounter))
		args = append(args, password)
		argCounter++
	}

	if len(updates) == 0 {
		// Jika tidak ada data yang di-update, coba ambil user yang ada
		// untuk memastikan ID tersebut valid.
		return GetUserById(db, id)
	}

	// Tambahkan klausa WHERE dan RETURNING
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d RETURNING id, name, email, created_at, updated_at", strings.Join(updates, ", "), argCounter)
	args = append(args, id)

	// Gunakan db.QueryRow untuk mengeksekusi dan mengambil hasilnya
	var user entities.User
	err := db.QueryRow(query, args...).Scan(&user.Id, &user.Name, &user.Email, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// Jika ID tidak ditemukan, kembalikan error yang spesifik.
			return nil, fmt.Errorf("user with ID %d not found: %w", id, sql.ErrNoRows)
		}
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return &user, nil
}
func DeleteUser(db *sql.DB, id int) error {
	rows, err := db.Query("delete from users where id = $1", id)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
