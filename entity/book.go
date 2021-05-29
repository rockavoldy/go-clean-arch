package entity

import "time"

//
type Book struct {
	ID        ID
	Title     string
	Author    string
	ISBN      string
	Pages     int
	Quantity  int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

// Add new Book
func NewBook(title string, author string, isbn string, pages int, qty int) (*Book, error) {
	book := &Book{
		ID:        NewID(),
		Title:     title,
		Author:    author,
		ISBN:      isbn,
		Pages:     pages,
		Quantity:  qty,
		CreatedAt: time.Now(),
	}

	err := book.ValidateInput()
	if err != nil {
		return nil, err
	}

	return book, err
}

// Validate input Book
func (book *Book) ValidateInput() error {
	if book.Title == "" || book.Author == "" || book.ISBN == "" || book.Pages <= 0 || book.Quantity <= 0 {
		return ErrInvalidEntity
	}

	return nil
}
