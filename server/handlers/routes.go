package handlers

import "github.com/gin-gonic/gin"

func AllRoutes(server *gin.Engine) {
	server.POST("/books", createBook)
}