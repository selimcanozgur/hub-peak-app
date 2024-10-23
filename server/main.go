package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/database"
	"github.com/selimcanozgur/hub-peak-app/handlers"
)

func main() {
	database.Connect()
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowAllOrigins: true, // Tüm origin'lere izin ver
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, // Gerekirse kimlik bilgilerine izin ver
	}))

	handlers.AllRoutes(server)
	server.Run(":8080")
}	