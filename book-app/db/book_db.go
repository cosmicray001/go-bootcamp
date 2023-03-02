package db

import (
	"fmt"

	"github.com/cosmicray001/book-app/models"
)

func InitDb() {
	var books = []models.Book{
		{"1", "In Search of Lost Time", "Marcel Proust", 1},
		{"2", "The Great Gatsby", "F. Scott", 2},
		{"3", "War and Peace", "Leo Tolstoy", 3},
	}
	fmt.Println("db created!", len(books))
}
