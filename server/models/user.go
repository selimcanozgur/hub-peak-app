package models

import (
	"time"

	"github.com/selimcanozgur/hub-peak-app/database"
	"github.com/selimcanozgur/hub-peak-app/utils"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}


func (u User) Save () error {

	if u.Role == "" {
		u.Role = "user"
	}


	query :=  "INSERT INTO users(email, password, role) VALUES (?, ?, ?)"
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword, u.Role)

	if err != nil {
		return err
	}

	userId, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = userId
	return err
}