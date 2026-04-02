package models

import (
	"errors"
	"rest_api/db"
	"rest_api/utils"
	"time"
)

type User struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
	DateTime time.Time `json:"date_time"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password, date_time) VALUES (?, ?, ?)"
	stmt, err := db.Db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword, u.DateTime)
	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}

	u.ID = userId
	return nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email=?"
	row := db.Db.QueryRow(query, u.Email)

	var retrievedPassword string

	if err := row.Scan(&u.ID, &retrievedPassword); err != nil {
		return errors.New("Invalid Credentials")
	}

	isPasswordValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !isPasswordValid {
		return errors.New("Invalid Credentials")
	}

	return nil
}
