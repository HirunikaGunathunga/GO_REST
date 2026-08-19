package model

import (
	"errors"

	db "github.com/HirunikaGunathunga/GO_REST/Database"
	utils "github.com/HirunikaGunathunga/GO_REST/Utils"
)

type Users struct {
	ID       int64  `json:"id"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u Users) Save() error {
	query := `INSERT INTO USERS_TAB (email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	u.ID = id
	return err

}

func (u *Users) ConfirmUser() error {
	query := `SELECT id, password FROM USERS_TAB WHERE email = ? `
	row := db.DB.QueryRow(query, u.Email)

	var retPassword string
	err := row.Scan(&u.ID, &retPassword)

	if err != nil {
		return errors.New("Credentials Invalid")
	}

	confirmUser := utils.CheckPasswords(retPassword, u.Password)
	if !confirmUser {
		return errors.New("Credentials Invalid")
	}
	return nil
}
