package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/selimcanozgur/hub-peak-app/db"
	"github.com/selimcanozgur/hub-peak-app/utils"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email" binding:"required"`
	Password  string    `json:"password" binding:"required"`
	FirstName string    `json:"first_name"` 
	LastName  string    `json:"last_name"`
	BirthDate time.Time `json:"birth_date`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) UnmarshalJSON(data []byte) error {
	type Alias User
	aux := &struct {
		BirthDate string `json:"birth_date"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	if aux.BirthDate != "" {
		parsedDate, err := time.Parse("2006-01-02", aux.BirthDate)
		if err != nil {
			return errors.New("birth_date formatı 'YYYY-MM-DD' olmalıdır")
		}
		u.BirthDate = parsedDate
	}

	return nil
}


func (u *User) SaveUser() error {

	if u.Role == "" {
        u.Role = "user"
    }

	query := "INSERT INTO users(email, password, first_name, last_name, birth_date, role, is_active) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)
	
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(u.Email, hashedPassword, u.FirstName, u.LastName, u.BirthDate, u.Role, u.IsActive)
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

func (u *User) LoginUser() error {
	
	query := "SELECT id, email, first_name, last_name, birth_date, password, role FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &u.Email, &u.FirstName, &u.LastName, &u.BirthDate, &retrievedPassword, &u.Role)

	if err != nil {
		fmt.Println(err)
		return errors.New("Kimlik bilgileri geçersiz")
	}

	passwordIsValid := utils.ComparePassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		fmt.Println(err)
		return errors.New("Kimlik bilgileri geçersiz")
	}

	return nil

}

func ListUsers() ([]User, error) {
	query := "SELECT id, email, is_active, first_name, last_name, birth_date, role, created_at, updated_at FROM users"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.IsActive,
			&user.FirstName,
			&user.LastName,
			&user.BirthDate,
			&user.Role,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}
