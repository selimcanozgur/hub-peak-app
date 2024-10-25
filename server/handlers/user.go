package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/models"
	"github.com/selimcanozgur/hub-peak-app/utils"
)

func createUser(context *gin.Context) {
	var user models.User
	
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Veri ayrıştırılamadı."})
		return
	}

	err = user.Save()

	if err != nil {
		log.Println("Hata:", err) // Hata mesajını konsola yazdır
		context.JSON(http.StatusInternalServerError, gin.H{"message":"Kullanıcı kaydedilemedi",})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Kullanıcı kaydı başarıyla oluşturuldu"})


}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message":"Colud not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusUnauthorized, gin.H{"message":"Could not authenticate user."})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID, user.Role)
	fmt.Println("Role Display: ", user.Role)

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login successfully", "token": token})
}