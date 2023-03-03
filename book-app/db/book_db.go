package db

import (
	"fmt"

	"github.com/cosmicray001/book-app/models"
)

var books []models.Book

func init() {
	books = append(books, models.Book{ID: "1", Title: "In Search of Lost Time", Author: "Marcel Proust", Quantity: 1})
	books = append(books, models.Book{ID: "2", Title: "The Great Gatsby", Author: "F. Scott", Quantity: 2})
	books = append(books, models.Book{ID: "3", Title: "War and Peace", Author: "Leo Tolstoy", Quantity: 3})
	fmt.Println("db created!", len(books))
}

func GetDB() (*[]models.Book) {
	return &books
}
