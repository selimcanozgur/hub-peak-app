package tables

var UserTable string
var BookTable string

func CreateTable() {
	UserTable = `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT PRIMARY KEY,
			email VARCHAR(100) NOT NULL UNIQUE, 
			password VARCHAR(255) NOT NULL,
			first_name VARCHAR(30) NOT NULL,
			last_name VARCHAR(30) NOT NULL,
			birth_date DATE,
			role ENUM('user', 'admin') DEFAULT 'user',
			is_active BOOLEAN DEFAULT FALSE,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
		`
	BookTable = `
	    CREATE TABLE IF NOT EXISTS books (
		  id INT AUTO_INCREMENT PRIMARY KEY,
          title VARCHAR(100) NOT NULL,
          author VARCHAR(100) NOT NULL,
          publishing_house VARCHAR(100) NOT NULL,
          publishing_year YEAR NOT NULL,
          price DECIMAL(6,2) NOT NULL,
	      img_path VARCHAR(255) NOT NULL,
          pages INT NOT NULL,
          summary LONGTEXT NOT NULL,
          lang VARCHAR(3) NOT NULL,
	      book_cover VARCHAR(30) NOT NULL,
          stock INT NOT NULL DEFAULT 0,
          created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
          updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
		)
	`
}