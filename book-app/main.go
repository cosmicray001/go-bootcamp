package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/cosmicray001/book-app/models"
)

var books []models.Book

func init() {
	books = append(books, models.Book{"1", "In Search of Lost Time", "Marcel Proust", 1})
	books = append(books, models.Book{"2", "The Great Gatsby", "F. Scott", 2})
	books = append(books, models.Book{"3", "War and Peace", "Leo Tolstoy", 3})
	fmt.Println("db created!", len(books))
}

func GetDb() []models.Book {
	return books
}



func bookList(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, books)
}

func bookCreate(ctx *gin.Context) {
	var book Book
	if err := ctx.BindJSON(&book); err != nil {
		return
	}
	books = append(books, book)
	ctx.IndentedJSON(http.StatusCreated, book)
}

func bookDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := getBookById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, book)
}



func bookCheckOut(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id must be sent"})
		return
	}
	book, err := getBookById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "out of stock"})
		return
	}

	book.Quantity -= 1
	ctx.IndentedJSON(http.StatusOK, book)
}

func main() {
	r := gin.Default()
	r.GET("/books", bookList)
	r.POST("/book", bookCreate)
	r.GET("/book/:id", bookDetail)
	r.PATCH("/checkout", bookCheckOut)
	r.Run("127.0.0.1:8000")
}
