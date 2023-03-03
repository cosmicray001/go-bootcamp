package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/cosmicray001/book-app/services"
)

func BookDetail(ctx *gin.Context) {
	id := ctx.Param("id")
	book, err := services.GetBookById(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, book)
}
