package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/models"
	"github.com/selimcanozgur/hub-peak-app/utils"
)

func getAllUser(c *gin.Context){
	users, err := models.ListUsers()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
		return
	}

	c.JSON(http.StatusOK, users)
}

func getUser(c *gin.Context) {
	token, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Token bulunamadı"})
		return
	}

	claims, err := utils.VerifyToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Geçersiz token"})
		return
	}

	response := gin.H{
		"id":         claims["id"],
		"email":      claims["email"],
		"first_name": claims["first_name"],
		"last_name":  claims["last_name"],
		"birth_date": claims["birth_date"],
		"role":       claims["role"],
	}

	c.JSON(http.StatusOK, response)
}
