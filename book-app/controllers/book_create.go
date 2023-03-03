package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/cosmicray001/book-app/models"
	"github.com/cosmicray001/book-app/db"
)

func BookCreate(ctx *gin.Context) {
	var book models.Book
	if err := ctx.BindJSON(&book); err != nil {
		return
	}
	books := db.GetDB()
	*books = append(*books, book)
	ctx.IndentedJSON(http.StatusCreated, book)
}