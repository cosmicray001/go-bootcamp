package main

import (
	"github.com/gin-gonic/gin"
	"github.com/cosmicray001/book-app/controllers"
)


func main() {
	r := gin.Default()
	r.GET("/books", controllers.BookList)
	r.POST("/book", controllers.BookCreate)
	r.GET("/book/:id", controllers.BookDetail)
	r.PATCH("/checkout", controllers.BookCheckOut)
	r.Run("127.0.0.1:8000")
}
