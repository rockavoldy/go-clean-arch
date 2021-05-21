package entity_test

import (
	"github.com/stretchr/testify/assert"
	"go-clean-arch/entity"
	"testing"
)

// Test NewBook
func TestNewBook(t *testing.T) {
	book, err := entity.NewBook("Ludonarasi", "Rimawarna", "102344", 170, 1)

	assert.Nil(t, err) // check if no error thrown
	assert.NotNil(t, book.ID) // check if ID generated
	assert.Equal(t, book.Title, "Ludonarasi") // check if Title equal
}