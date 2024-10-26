package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/selimcanozgur/hub-peak-app/models"
)

func getAllBooks(context *gin.Context) {
	title := context.Query("title")
	order := context.Query("filter")

	books, err := models.AllBookQuery(title, order)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Kitaplar yüklenemedi tekrar deneyiniz.", "err": err.Error()})
		return
	}

	context.JSON(http.StatusOK, books)
}


func getOneBook (context *gin.Context) {
	bookId, err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Kitap kimliği ayrıştırılamadı"})

	}

	book, err := models.GetBookById(bookId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Kitap verisi getirilemedi."})
		return
	}

	context.JSON(http.StatusOK, book)
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

func updateBook (context *gin.Context) {
	bookId, err := strconv.ParseInt(context.Param("id"),10,64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Kitap kimliği ayrıştırılamadı"})
	}

	_, err = models.GetBookById(bookId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Kitap verisi getirilemedi"})
		return
	}

	var updateBook models.Book
	err = context.ShouldBindJSON(&updateBook)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse"})
		return
	}

	updateBook.ID = bookId
	err = updateBook.Update()

	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update boook."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Kitap başarıyla güncellendi"})

}

func deleteBook (context *gin.Context) {
	bookId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse book id."})
	}

	book, err := models.GetBookById(bookId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the book."})
		return
	}

	err = book.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the book."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}

