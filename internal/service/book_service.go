package service

import (
    "github.com/DLLenjoyer/books-api/models"
    "github.com/DLLenjoyer/books-api/internal/repository"
)

type BookService struct {
	repo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{
		repo: repo,
	}
}

func (s *BookService) GetAll() ([]models.Book, error) {
	return s.repo.GetAll()
}

func (s *BookService) GetByID(id string) (*models.Book, error) {
    return s.repo.GetByID(id)
}

func (s *BookService) Add(book *models.Book) error {
	return s.repo.Add(book)
}

func (s *BookService) Update(book *models.Book) error {
	return s.repo.Update(book)
}

func (s *BookService) Delete(id string) error {
	return s.repo.Delete(id)
}
