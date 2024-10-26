package models

import (
	"database/sql"
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

func AllBookQuery(title string, order string) ([]Book, error) {
	query := "SELECT * FROM books"

	// Filter by title if provided
	if title != "" {
		query += " WHERE title LIKE ?"
		title = "%" + title + "%"
	}

	// Order by price if order parameter is provided
	if order == "asc" {
		query += " ORDER BY price ASC"
	} else if order == "desc" {
		query += " ORDER BY price DESC"
	}

	// Execute the query
	var rows *sql.Rows
	var err error
	if title != "" {
		rows, err = database.DB.Query(query, title)
	} else {
		rows, err = database.DB.Query(query)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []Book
	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.PublishingHouse,
			&book.PublishingYear, &book.Price, &book.ImgPath, &book.Pages, &book.Summary,
			&book.Lang, &book.BookCover, &book.Stock, &book.BookID, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func GetBookById (id int64) (*Book, error) {
	query := "SELECT * FROM books WHERE id = ?"
	row := database.DB.QueryRow(query, id)

	var book Book

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.PublishingHouse, 
		&book.PublishingYear, &book.Price, &book.ImgPath, &book.Pages, &book.Summary, 
		&book.Lang, &book.BookCover, &book.Stock, &book.BookID, &book.CreatedAt, &book.UpdatedAt)

		if err != nil {
			return nil, err
		}
	
		return &book, nil
}

func (book Book) Update ()  error {
	query := `
	UPDATE books
	SET title = ?, author = ?, publishing_house = ?, publishing_year= ?, price = ?,
	img_path = ?, pages = ?, summary = ?, lang = ?, book_cover = ?, stock = ?, book_id = ?
	WHERE id = ?
	`
	stmt, err := database.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Author, book.PublishingHouse, book.PublishingYear, 
	book.Price, book.ImgPath, book.Pages, book.Summary, book.Lang, book.BookCover, book.Stock, book.BookID, book.ID)

	return err
}


func (book Book) Delete () error {
	query := "DELETE FROM books WHERE id = ?"
	stmt, err := database.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(book.ID)
	return err
}