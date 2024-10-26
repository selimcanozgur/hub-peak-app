package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/utils"
)

func AllRoutes(server *gin.Engine) {

	server.POST("/signup", createUser)
	server.POST("/login", login)
	server.GET("/books",getAllBooks)
	server.GET("/books/:id", getOneBook)


	adminRoutes := server.Group("/books")
	adminRoutes.Use(utils.AuthMiddleware(), utils.AdminMiddleware()) 
	
	{
		adminRoutes.POST("/admin", createBook)
		adminRoutes.PUT("/admin/:id", updateBook)
		adminRoutes.DELETE("/admin/:id", deleteBook)
	}
}