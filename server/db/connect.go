package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_URI := os.Getenv("DB_URI")

	DB, err := sql.Open("mysql", DB_URI)

	if err != nil {
		log.Fatal("Veritabanı açılırken hata oluştu: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Veritabanına bağlanırken hata oluştu: %v", err)
	}

	log.Println("Veritabanı bağlantısı başarılı")


}