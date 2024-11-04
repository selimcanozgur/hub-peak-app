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


	books, err := models.ListBooks(title, order)


	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Kitaplar yüklenemedi tekrar deneyiniz."})
		return
	}


	if len(books) == 0 {
		context.JSON(http.StatusNotFound, gin.H{"message": "Aradığınız kitap bulunamadı."})
		return
	}

	context.JSON(http.StatusOK, books)
}


func createBook(c *gin.Context) {
	var book models.Book
	err := c.ShouldBindJSON(&book)

	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{"message":"Kitap verileri ayrıştırılamadı"})
		return
	}

	err = book.SaveBook()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"message":"Kitap oluşturulamadı"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kitap başarıyla oluşturuldu"})
}

func getOneBook(c *gin.Context) {
	bookId, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Kitap kimliği ayrıştırılamadı"})

	}

	book, err := models.GetBookById(bookId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Kitap verisi getirilemedi."})
		return
	}

	c.JSON(http.StatusOK, book)
}

func updateBook (c *gin.Context) {
	bookId, err := strconv.ParseInt(c.Param("id"),10,64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Kitap kimliği ayrıştırılamadı"})
	}

	_, err = models.GetBookById(bookId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Kitap verisi getirilemedi"})
		return
	}

	var updateBook models.Book
	err = c.ShouldBindJSON(&updateBook)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse"})
		return
	}

	updateBook.ID = bookId
	err = updateBook.Update()

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update boook."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Kitap başarıyla güncellendi"})

}

func deleteBook (c *gin.Context) {
	bookId, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse book id."})
	}

	book, err := models.GetBookById(bookId)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the book."})
		return
	}

	err = book.Delete()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the book."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Book deleted successfully"})
}
