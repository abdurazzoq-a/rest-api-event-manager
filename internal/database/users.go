package database

import (
	"context"
	"database/sql"
	"time"
)

type UserModel struct {
	DB *sql.DB
}

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (u *UserModel) Insert(user *User) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	query := "INSERT INTO users (email, password, name) VALUES ($1, $2, $3) RETURNING id"

	return u.DB.QueryRowContext(ctx, query, user.Email, user.Password, user.Name).Scan(&user.Id)
}

func (u *UserModel) getUser(query string, args ...interface{}) (*User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	var user User

	err := u.DB.QueryRowContext(ctx, query, args...).Scan(&user.Id, &user.Email, &user.Name, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, err
	}

	return &user, nil
}

func (u *UserModel) GetUserById(id int) (*User, error) {
	query := "SELECT * FROM users WHERE id = $1"
	return u.getUser(query, id)
}

func (u *UserModel) GetUserByEmail(email string) (*User, error) {
	query := "SELECT * FROM users WHERE email = $1"
	return u.getUser(query, email)
}
