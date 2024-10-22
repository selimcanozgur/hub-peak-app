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

}
