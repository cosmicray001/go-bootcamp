package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/cosmicray001/book-app/db"
)

func BookList(ctx *gin.Context) {
	books := db.GetDB()
	ctx.IndentedJSON(http.StatusOK, &books)
}