package services

import (
	"github.com/cosmicray001/book-app/models"
	"errors"
)

func getBookById(id string) (*models.Book, error) {
	for idx, book := range books {
		if book.ID == id {
			return &books[idx], nil
		}
	}
	return nil, errors.New("Book not found")
}