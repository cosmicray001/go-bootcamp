package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/cosmicray001/book-app"
)

func BookList(ctx *gin.Context) {
	books := main.GetDb()
	ctx.IndentedJSON(http.StatusOK, books)
}