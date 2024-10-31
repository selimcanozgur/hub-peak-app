package main

import (
	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/db"
)

func main() {
	db.Connect()
	server := gin.Default()

	server.Run(":8080")

}

