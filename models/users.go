package models

import (
	"rest_api/db"
	"rest_api/utils"
	"time"
)

type User struct {
	ID       int64     `json:"id"`
	Email    string    `json:"email" binding:"required"`
	Password string    `json:"password" binding:"required"`
	DateTime time.Time `json:"date_time" binding:"required"`
}

func(u *User) Save() error {
	query := "INSERT INTO users(email, password, date_time) VALUES (?, ?, ?)"
	stmt, err := db.Db.Prepare(query)
	if err!=nil {
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
