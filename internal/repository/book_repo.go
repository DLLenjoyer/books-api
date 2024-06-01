package repository

import (
	"github.com/DLLenjoyer/books-api/models"
)

type BookRepository interface {
	GetAll() ([]models.Book, error)
	GetByID(id string) (*models.Book, error)
	Add(book *models.Book) error
	Update(book *models.Book) error
	Delete(id string) error
}

type InMemoryBook struct {
	Books []models.Book
}

func NewInMemoryBook() *InMemoryBook {
	return &InMemoryBook{
		Books: []models.Book{},
	}
}

func (r *InMemoryBook) GetAll() ([]models.Book, error) {
	return r.Books, nil
}

func (r *InMemoryBook) GetByID(id string) (*models.Book, error) {
	for _, book := range r.Books {
		if book.ID == id {
			return &book, nil
		}
	}
	return nil, nil
}

func (r *InMemoryBook) Add(book *models.Book) error {
	r.Books = append(r.Books, *book)
	return nil
}

func (r *InMemoryBook) Update(book *models.Book) error {
	for i, b := range r.Books {
		if b.ID == book.ID {
			r.Books[i] = *book
			return nil
		}
	}
	return nil
}

func (r *InMemoryBook) Delete(id string) error {
	for i, book := range r.Books {
		if book.ID == id {
			r.Books = append(r.Books[:i], r.Books[i+1:]...)
			return nil
		}
	}
	return nil
}
