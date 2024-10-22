package main

import (
	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/database"
)

func main() {
	database.Connect()
	server := gin.Default()
	server.Run(":8080")
}