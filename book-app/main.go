package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []Book{
	{"1", "In Search of Lost Time", "Marcel Proust", 1},
	{"2", "The Great Gatsby", "F. Scott", 2},
	{"3", "War and Peace", "Leo Tolstoy", 3},
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

func main() {
	r := gin.Default()
	r.GET("/books", bookList)
	r.POST("/book", bookCreate)
	r.Run("127.0.0.1:8000")
}
