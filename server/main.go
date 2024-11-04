package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/db"
	"github.com/selimcanozgur/hub-peak-app/handlers"
)

func main() {
	db.Connect()
	server := gin.Default()

	server.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true, 
	}))
	handlers.AllRoutes(server)
	server.Run(":8080")

}

