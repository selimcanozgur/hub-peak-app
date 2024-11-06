package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/selimcanozgur/hub-peak-app/tables"
)

var DB *sql.DB

func Connect() {

	var err error

	err = godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DB_URI := os.Getenv("DB_URI")

	DB, err = sql.Open("mysql", DB_URI)

	if err != nil {
		log.Fatalf("Veritabanı açılırken hata oluştu: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Veritabanına bağlanırken hata oluştu: %v", err)
	}

	log.Println("Veritabanı bağlantısı başarılı")

	tables.CreateTable()

	_, err = DB.Exec(tables.UserTable)

	if err != nil {
		panic("Kullanıcı tablosu oluşturulamadı")
	}

	_, err = DB.Exec(tables.BookTable)

	if err != nil {
		panic("Kitap tablosu oluşturulamadı")
	}

}