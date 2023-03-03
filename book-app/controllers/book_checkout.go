package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/cosmicray001/book-app/services"
)

func BookCheckOut(ctx *gin.Context) {
	id, ok := ctx.GetQuery("id")
	if !ok {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id must be sent"})
		return
	}
	book, err := services.GetBookById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}

	if book.Quantity <= 0 {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"message": "out of stock"})
		return
	}
	book.Quantity -= 1

	ctx.IndentedJSON(http.StatusOK, &book)
}
