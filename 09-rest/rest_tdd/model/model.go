// model.go

package model

import (
	"database/sql"
	"fmt"
)

type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
	Age      int    `json:"age,omitempty"`
	Active   bool   `json:"active,omitempty"`
}

func (u *User) getUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT name, age FROM users WHERE id=%d", u.ID)
	return db.QueryRow(statement).Scan(&u.Name, &u.Age)
	// return errors.New("Not implemented")
}

func (u *User) updateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE users SET name='%s', age=%d WHERE id=%d", u.Name, u.Age, u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *User) deleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM users WHERE id=%d", u.ID)
	_, err := db.Exec(statement)
	return err
}

func (u *User) createUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO users(name, email, age) VALUES('%s', %d)", u.Name, u.Email, u.Age)
	_, err := db.Exec(statement)

	if err != nil {
		return err
	}

	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func getUsers(db *sql.DB, start, count int) ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, password, age FROM users LIMIT ? OFFSET ?", count, start)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []User{}

	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
	// return nil, errors.New("Not implemented")
}
