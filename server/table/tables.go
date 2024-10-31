package table

var UserTable string

func CreateTable() {
	UserTable = `
	CREATE TABLE IF NOT EXISTS users (
	    id INT AUTO_INCREMENT PRIMARY KEY,
		email VARCHAR(100) NOT NULL UNIQUE, 
		password VARCHAR(255) NOT NULL,
		role ENUM('user', 'admin') DEFAULT 'user',
		is_active BOOLEAN DEFAULT FALSE,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	)
	`

}