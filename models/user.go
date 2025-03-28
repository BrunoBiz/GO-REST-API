package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users (email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	defer stmt.Close()
	result, err := stmt.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	userid, err := result.LastInsertId()

	u.ID = userid
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT ID, PASSWORD FROM USERS WHERE EMAIL = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("credentials invalid ")
	}

	validPassword := utils.CheckPasswordHash(u.Password, retrievedPassword)

	if !validPassword {
		return errors.New("credentials invalid ")
	}
	return nil
}
