package book

import "go-clean-arch/entity"

// Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.Book, error)
	List() ([]*entity.Book, error)
	Search(query string) ([]*entity.Book, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.Book) (entity.ID, error)
	Update(e *entity.Book) error
	Delete(id entity.ID) error
	Restore (id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetBook(id entity.ID) (*entity.Book, error)
	GetDeletedBook(id entity.ID) (*entity.Book, error)
	ListBooks() ([]*entity.Book, error)
	SearchBooks(query string) ([]*entity.Book, error)
	CreateBook(title, author, isbn string, pages, qty int) (entity.ID, error)
	UpdateBook(e *entity.Book) error
	DeleteBook(id entity.ID) error
	RestoreBook(id entity.ID) error
}