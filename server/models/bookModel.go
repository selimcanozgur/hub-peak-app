package models

import (
	"database/sql"
	"time"

	"github.com/selimcanozgur/hub-peak-app/db"
)

type Book struct {
	ID              int64        `json:"id"`
	Title           string       `json:"title"`
	Author          string       `json:"author"`
	PublishingHouse string       `json:"publishing_house"`
	PublishingYear  int          `json:"publishing_year"`
	Price           string       `json:"price"`
	ImgPath         string       `json:"img_path"`
	Pages           int          `json:"pages"`
	Summary         string       `json:"summary"`
	Lang            string       `json:"lang"`
	BookCover       string       `json:"book_cover"`
	Stock           int          `json:"stock"`
	CreatedAt       time.Time    `json:"created_at"`
	UpdatedAt       time.Time    `json:"updated_at"`
}

// Get All Book
func ListBooks(title string, order string) ([]Book, error) {
	query := "SELECT * FROM books"

	if title != "" {
		query += " WHERE title LIKE ?"
		title = "%" + title + "%"
	}

	if order == "asc" {
		query += " ORDER BY price ASC"
	} else if order == "desc" {
		query += " ORDER BY price DESC"
	}


	var rows *sql.Rows
	var err error
	if title != "" {
		rows, err = db.DB.Query(query, title)
	} else {
		rows, err = db.DB.Query(query)
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
			&book.Lang, &book.BookCover, &book.Stock, &book.CreatedAt, &book.UpdatedAt); err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}


// Create Book
func (b *Book) SaveBook () error {
	query := `
	INSERT INTO books (title, author, publishing_house, publishing_year, price, img_path, pages, summary, lang, book_cover, stock)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(b.Title, b.Author, b.PublishingHouse, b.PublishingYear, b.Price, b.ImgPath, b.Pages, b.Summary, b.Lang, b.BookCover,  b.Stock)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	b.ID = id
	return err
}


// Get Book Detail
func GetBookById(id int64) (*Book, error) {
	query := "SELECT * FROM books WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var book Book

	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.PublishingHouse, 
		&book.PublishingYear, &book.Price, &book.ImgPath, &book.Pages, &book.Summary, 
		&book.Lang, &book.BookCover, &book.Stock, &book.CreatedAt, &book.UpdatedAt)

		if err != nil {
			return nil, err
		}
	
		return &book, nil
}


// Update Book
func (book Book) Update ()  error {
	query := `
	UPDATE books
	SET title = ?, author = ?, publishing_house = ?, publishing_year= ?, price = ?,
	img_path = ?, pages = ?, summary = ?, lang = ?, book_cover = ?, stock = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(book.Title, book.Author, book.PublishingHouse, book.PublishingYear, 
	book.Price, book.ImgPath, book.Pages, book.Summary, book.Lang, book.BookCover, book.Stock, book.ID)

	return err
}


// Delete Book
func (book Book) Delete () error {
	query := "DELETE FROM books WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(book.ID)
	return err
}