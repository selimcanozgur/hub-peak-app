package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/models"
	"github.com/selimcanozgur/hub-peak-app/utils"
)

func signup(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)

	if err != nil {

		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Veri ayrıştırılamadı"})
		return
	}

	err = user.SaveUser()

	if err != nil {
		log.Println("Hata:", err) // Hata mesajını konsola yazdır
		c.JSON(http.StatusInternalServerError, gin.H{"message":"Kullanıcı kaydedilemedi",})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Kullanıcı kaydı başarıyla oluşturuldu"})

}

func login (c *gin.Context){
	var user models.User
	
	err := c.ShouldBindJSON(&user)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message": "Veri ayrıştırılamadı"})
		return
	}

	err = user.LoginUser()

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Kullanıcı kimliği doğrulanamadı"})
		return
	}

	token, err := utils.GenerateToken(user.ID, user.Role)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Kullanıcı kimliği doğrulanamadı, token."})
		return
	}

	c.SetCookie("token", token, 3600, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{"message": "Giriş başarılı"})

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
