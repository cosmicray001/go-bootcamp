package services

import (
	"errors"

	"github.com/cosmicray001/book-app/db"
	"github.com/cosmicray001/book-app/models"
)

func GetBookById(id string) (*models.Book, error) {
	books := db.GetDB()
	for _, book := range *books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, errors.New("Book not found")
}
