package models

import (
	"time"

	"github.com/selimcanozgur/hub-peak-app/database"
)

type Book struct {
	ID              int64     `json:"id"`               
	Title           string    `json:"title" binding:"required"`          
	Author          string    `json:"author" binding:"required"`          
	PublishingHouse string    `json:"publishing_house" binding:"required"` 
	PublishingYear  int       `json:"publishing_year" binding:"required"`  
	Price           float64   `json:"price" binding:"required"`
	ImgPath         string    `json:"img_path" binding:"required"`         
	Pages           int       `json:"pages" binding:"required"`            
	Summary         string    `json:"summary" binding:"required"`      
	Lang            string    `json:"lang" binding:"required"`
	BookCover       string    `json:"book_cover" binding:"required"`      
	Stock           int       `json:"stock" binding:"required"`            
	BookID          int       `json:"book_id"`                                                  
	CreatedAt       time.Time `json:"created_at"`                           
	UpdatedAt       time.Time `json:"updated_at"`
}

var books = []Book{}

func (b Book) Save () error {
	query := `
	INSERT INTO books (title, author, publishing_house, publishing_year, price, img_path, pages, summary, lang, book_cover, stock, book_id)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(b.Title, b.Author, b.PublishingHouse, b.PublishingYear, b.Price, b.ImgPath, b.Pages, b.Summary, b.Lang, b.BookCover,  b.Stock, b.BookID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	b.ID = id
	return err
}

func AllBookQuery () ([]Book, error) {
	query := "SELECT * FROM books"
	rows, err := database.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var books []Book

	for rows.Next(){

		var book Book
		
		err := rows.Scan(    
			&book.ID, &book.Title, &book.Author, &book.PublishingHouse, 
			&book.PublishingYear, &book.Price, &book.ImgPath, &book.Pages, &book.Summary, 
			&book.Lang, &book.BookCover, &book.Stock, &book.BookID, &book.CreatedAt, &book.UpdatedAt,)	

		if err != nil {
			return nil, err
		}
		

		books = append(books, book)
	}

	return books, nil
}
