package book

import (
	"go-clean-arch/entity"
	"time"
)

type Service struct {
	repo Repository
}

// create NewService
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// create Book
func (s *Service) CreateBook(title, author, isbn string, pages, qty int) (entity.ID, error) {
	e, err := entity.NewBook(title, author, isbn, pages, qty)
	if err != nil {
		return e.ID, err
	}

	return s.repo.Create(e)
}

// GetBook
func (s *Service) GetBook(id entity.ID) (*entity.Book, error) {
	return s.repo.Get(id)
}

// Get ListBooks
func (s *Service) ListBooks() ([]*entity.Book, error) {
	return s.repo.List()
}

// Search Book by Title
func (s *Service) SearchBooks(query string) ([]*entity.Book, error) {
	return s.repo.Search(query)
}

// Update Book
func (s *Service) UpdateBook(e *entity.Book) error {
	err := e.ValidateInput()
	if err != nil {
		return err
	}

	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// Delete Book
func (s *Service) DeleteBook(id entity.ID) error {
	b, err := s.GetBook(id)
	if b == nil {
		return entity.ErrNotFound
	}

	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

// Restore deleted Book
func (s *Service) RestoreBook(id entity.ID) error {
	b, err := s.GetBook(id)
	if b == nil {
		return entity.ErrNotFound
	}

	if err != nil {
		return err
	}

	return s.repo.Restore(id)
}
