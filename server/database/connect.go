package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatal(".env dosyası yüklenemedi")
	}

	DB_URI := os.Getenv("DB_URI")

	DB, err = sql.Open("mysql", DB_URI)

	if err != nil {
		log.Fatal("Veritabanı bağlantı yolu hatalı veya geçersiz:", err)
	}

	err = DB.Ping()

	if err != nil {
		log.Fatal("Veritabanına bağlanılamadı, sunucuya erişilemiyor:", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createBookTable()

}

func createBookTable () {
	bookTable := `
	CREATE TABLE IF NOT EXISTS books (
	 id INT AUTO_INCREMENT PRIMARY KEY,
     title VARCHAR(100) NOT NULL,
     author VARCHAR(100) NOT NULL,
     publishing_house VARCHAR(100) NOT NULL,
     publishing_year YEAR NOT NULL,
     price DECIMAL(8,2) NOT NULL,
     pages INT NOT NULL,
     summary LONGTEXT NOT NULL,
     lang VARCHAR(20) NOT NULL,
     cover_type ENUM('Ciltli', 'İnce Kapak') NOT NULL,
     image_url VARCHAR(255),
     isbn VARCHAR(13) UNIQUE,
     stock INT NOT NULL DEFAULT 0,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)`

	_, err := DB.Exec(bookTable)

	if err != nil {
		log.Fatal("Tablo oluşturulamadı.", err)
	}
}
