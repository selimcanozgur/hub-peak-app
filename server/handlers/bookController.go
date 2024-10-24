package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/models"
)

func getAllBooks (context *gin.Context) {

	books, err := models.AllBookQuery()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Kitaplar yüklenemedi tekrar deneyiniz.", "err": err.Error()})
		return
	}

	context.JSON(http.StatusOK, books)
}

func createBook(context *gin.Context) {
	var book models.Book
	err := context.ShouldBindJSON(&book)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Geçersiz JSON formatı"})
		return
	}

	book.ID = 1
	book.BookID = 1

	err = book.Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Kitap eklenemedi, tekrar deneyiniz",
			"err": err.Error(),
		})
		return
	}
	

	context.JSON(http.StatusCreated, gin.H{"message": "Kitap oluşturuldu."})

}

