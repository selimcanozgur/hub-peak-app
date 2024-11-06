package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/middlewares"
)

func AllRoutes(server *gin.Engine) {
	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/profile", getUser)
	server.GET("/books",getAllBooks)
	server.GET("/books/:id",getOneBook)
	
	adminRoutesBook := server.Group("/books")
	adminRoutesBook.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware()) 

	{
		adminRoutesBook.POST("/admin", createBook)
		adminRoutesBook.PUT("/admin/:id", deleteBook)
		adminRoutesBook.DELETE("/admin/:id", updateBook)
	}

	adminRoutesUser := server.Group("/users")
	adminRoutesUser.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())

	{
		adminRoutesUser.GET("/admin", getAllUser)

	}
}