package models

import (
	"time"

	"github.com/selimcanozgur/hub-peak-app/database"
)

type CoverType string

type Book struct {
	ID              int64     `json:"id"`               
	Title           string    `json:"title" binding:"required"`          
	Author          string    `json:"author" binding:"required"`          
	PublishingHouse string    `json:"publishing_house" binding:"required"` 
	PublishingYear  int       `json:"publishing_year" binding:"required"`  
	Price           float64   `json:"price" binding:"required"`           
	Pages           int       `json:"pages" binding:"required"`            
	Summary         string    `json:"summary" binding:"required"`      
	Lang            string    `json:"lang" binding:"required"`             
	ImageUrl        string    `json:"image_url" binding:"required"`   
	Stock           int       `json:"stock" binding:"required"`            
	BookID          int       `json:"book_id"`                            
	CoverType       string    `json:"cover_type"`                          
	CreatedAt       time.Time `json:"created_at"`                           
	UpdatedAt       time.Time `json:"updated_at"`
}

var books = []Book{}

func (b Book) Save () error {
	query := `
	INSERT INTO books (title, author, publishing_house, publishing_year, price, pages, summary, lang, image_url, stock, book_id, cover_type)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(b.Title, b.Author, b.PublishingHouse, b.PublishingYear, b.Price, b.Pages, b.Summary, b.Lang, b.ImageUrl, b.Stock, b.BookID, b.CoverType)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	b.ID = id
	return err
}
