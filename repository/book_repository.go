package repository

import (
	"database/sql"
	"go-clean-arch/entity"
	"time"
)

type BookMySQL struct {
	db *sql.DB
}

// Create new MySQL Repository
func NewBookMySQL(db *sql.DB) *BookMySQL {
	return &BookMySQL{db: db}
}

// create Book
func (r *BookMySQL) Create(e *entity.Book) (entity.ID, error) {
	stmt, err := r.db.Prepare(`INSERT INTO book (id, title, author, isbn, pages, quantity, created_at) VALUES (?,?,?,?,?,?,?)`)
	defer stmt.Close()
	if err != nil {
		return e.ID, err
	}

	_, err = stmt.Exec(
		e.ID,
		e.Title,
		e.Author,
		e.ISBN,
		e.Pages,
		e.Quantity,
		time.Now().Format(entity.FormatDateTimeSQL),
	)
	if err != nil {
		return e.ID, err
	}

	return e.ID, nil
}

// Get Book
func (r *BookMySQL) Get(id entity.ID) (*entity.Book, error) {
	rows, err := r.db.Query(`SELECT id, title, author, isbn, pages, quantity, created_at FROM book WHERE id = ? AND deleted_at IS NULL`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var book entity.Book
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Pages, &book.Quantity, &book.CreatedAt)
	}

	return &book, nil
}

// List Book
func (r *BookMySQL) List() ([]*entity.Book, error) {
	rows, err := r.db.Query(`SELECT id, title, author, isbn, pages, quantity, created_at FROM book WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Pages, &book.Quantity, &book.CreatedAt)
		books = append(books, &book)
	}

	return books, nil
}

// Search Book by Book.Title
func (r *BookMySQL) Search(query string) ([]*entity.Book, error) {
	rows, err := r.db.Query(`SELECT id, title, author, isbn, pages, quantity, created_at FROM book WHERE title LIKE ? AND deleted_at IS NULL`, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []*entity.Book
	for rows.Next() {
		var book entity.Book
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Pages, &book.Quantity, &book.CreatedAt)
		books = append(books, &book)
	}

	return books, nil
}

// Update Book
func (r *BookMySQL) Update(e *entity.Book) error {
	e.UpdatedAt = time.Now()
	_, err := r.db.Exec(`UPDATE book SET title = ?, author = ?, isbn = ?, pages = ?, quantity = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL`,
		e.Title, e.Author, e.ISBN, e.Pages, e.Quantity, e.UpdatedAt.Format(entity.FormatDateTimeSQL), e.ID)
	if err != nil {
		return err
	}

	return nil
}

// Delete a Book
func (r *BookMySQL) Delete(id entity.ID) error {
	_, err := r.db.Exec(`UPDATE book SET deleted_at = ? WHERE id = ?`, time.Now().Format(entity.FormatDateTimeSQL), id)
	if err != nil {
		return err
	}

	return nil
}

// GetD deleted Book
func (r *BookMySQL) GetDeletedBook(id entity.ID) (*entity.Book, error) {
	rows, err := r.db.Query(`SELECT id, title, author, isbn, pages, quantity, created_at FROM book WHERE id = ? AND deleted_at IS NOT NULL`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var book entity.Book
	for rows.Next() {
		err = rows.Scan(&book.ID, &book.Title, &book.Author, &book.ISBN, &book.Pages, &book.Quantity, &book.CreatedAt, &book.DeletedAt)
	}

	return &book, nil
}

// Restore a Book
func (r *BookMySQL) Restore(id entity.ID) error {
	_, err := r.db.Exec(`UPDATE book SET deleted_at = null WHERE id = ? AND deleted_at IS NOT NULL`, id)
	if err != nil {
		return err
	}

	return nil
}
