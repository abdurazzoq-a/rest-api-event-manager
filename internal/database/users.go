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


func (u *UserModel) Get(userId int) (*User, error) {
	// ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	// defer cancel()

	// query := "SELECT * FROM users WHERE id = $1"

	return nil, nil
}