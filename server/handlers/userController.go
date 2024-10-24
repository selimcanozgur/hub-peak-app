package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/models"
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