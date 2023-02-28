package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
)

type Book struct {
  ID string `json:"id"`
  Title string `json:"title"`
  Author string `json:"author"`
  Quantity int `json:"quantity"`
}

var books = []Book{
  {"1", "In Search of Lost Time", "Marcel Proust", 1},
  {"2", "The Great Gatsby", "F. Scott", 2},
  {"3", "War and Peace", "Leo Tolstoy", 3},
}

func main() {
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
  })
  r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
