package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/middlewares"
)

func AllRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/user", getUser)
	server.GET("/books",getAllBooks)
	server.GET("/books/:id",getOneBook)
	
	adminRoutes := server.Group("/books")
	adminRoutes.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware()) 

	{
		adminRoutes.POST("/admin", createBook)
		adminRoutes.PUT("/admin/:id", deleteBook)
		adminRoutes.DELETE("/admin/:id", updateBook)
	}
}