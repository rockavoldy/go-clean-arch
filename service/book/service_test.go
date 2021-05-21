package book

import (
	"github.com/stretchr/testify/assert"
	"go-clean-arch/entity"
	"testing"
	"time"
)

func newTestBook() *entity.Book {
	return &entity.Book{
		ID: entity.NewID(),
		Title: "Ludonarasi",
		Author: "Rimawarna",
		ISBN: "123444",
		Pages: 170,
		Quantity: 3,
		CreatedAt: time.Now(),
	}
}

func TestService_CreateBook(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	b := newTestBook()

	_, err := m.CreateBook(b.Title, b.Author, b.ISBN, b.Pages, b.Quantity)

	assert.Nil(t, err)
	assert.False(t, b.CreatedAt.IsZero())
	assert.True(t, b.UpdatedAt.IsZero())
	assert.True(t, b.DeletedAt.IsZero())
}

func TestService_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	b1 := newTestBook()
	b2 := newTestBook()
	b2.Title = "Lost Symbol"

	bookId, _ := m.CreateBook(b1.Title, b1.Author, b1.ISBN, b1.Pages, b1.Quantity)
	_, _ = m.CreateBook(b2.Title, b2.Author, b2.ISBN, b2.Pages, b2.Quantity)

	t.Run("Search", func(t *testing.T) {
		c, err := m.SearchBooks("ludo")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "Rimawarna", c[0].Author)

		c, err = m.SearchBooks("ludotarasi")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("List all books", func(t *testing.T) {
		all, err := m.ListBooks()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("Get book", func(t *testing.T) {
		saved, err := m.GetBook(bookId)
		assert.Nil(t, err)
		assert.Equal(t, b1.Title, saved.Title)
	})
}

func TestService_DeleteBook(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	b1 := newTestBook()
	b2 := newTestBook()

	book2Id, _ := m.CreateBook(b2.Title, b2.Author, b2.ISBN, b2.Pages, b2.Quantity)

	t.Run("Delete non-existed book", func(t *testing.T) {
		err := m.DeleteBook(b1.ID)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("Check if book successfully deleted", func(t *testing.T) {
		err := m.DeleteBook(book2Id)
		assert.Nil(t, err)
		_, err = m.GetBook(book2Id)
		assert.Equal(t, entity.ErrNotFound, err)
	})
}