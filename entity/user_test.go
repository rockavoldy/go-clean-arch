package entity_test

import (
	"github.com/stretchr/testify/assert"
	"go-clean-arch/entity"
	"testing"
)

type userRaw struct {
	email    string
	password string
	name     string
}

// create User entity
func createUser() (*entity.User, userRaw, error) {

	u := userRaw{
		email:    "fuu@fuu.com",
		password: "new_pass",
		name:     "Fuu",
	}
	user, err := entity.NewUser(u.email, u.password, u.name)
	return user, u, err
}

// Test NewUser function in User entity
func TestNewUser(t *testing.T) {
	user, userRaw, err := createUser()

	assert.Nil(t, err)                                  // check if no error thrown
	assert.NotNil(t, user.ID)                           // check if ID generated
	assert.Equal(t, user.Name, userRaw.name)            // check if name equal to input
	assert.Equal(t, user.Email, userRaw.email)          // check if email equal to input
	assert.NotEqual(t, user.Password, userRaw.password) // check if password successfully hashed
}

// Test Validate function in User entity
func TestUser_ValidatePassword(t *testing.T) {
	user, userRaw, _ := createUser()
	err := user.ValidatePassword(userRaw.password)
	assert.Nil(t, err) // ValidatePassword should NOT thrown error

	err = user.ValidatePassword("wrong_password")
	assert.NotNil(t, err) // ValidatePassword SHOULD thrown error
}

// Test add new Book to User
func TestUser_AddBook(t *testing.T) {
	user, _, _ := createUser()
	bookId := entity.NewID()
	err := user.AddBook(bookId)

	assert.Nil(t, err)                  // check if add book thrown error
	assert.Equal(t, 1, len(user.Books)) // check if user.Books have 1 Book

	// check if function thrown ErrAlreadyExist when same Book added to User
	err = user.AddBook(bookId)
	assert.Equal(t, entity.ErrAlreadyExist, err)
}

func TestUser_GetBook(t *testing.T) {
	user, _, _ := createUser()
	bookId := entity.NewID()
	_ = user.AddBook(bookId)

	got, err := user.GetBook(bookId)
	assert.Nil(t, err)           // check if GetBook thrown error
	assert.Equal(t, bookId, got) // check if function found same ID

	invalidBookId := entity.NewID()
	_, err = user.GetBook(invalidBookId)
	assert.Equal(t, entity.ErrNotFound, err) // check function thrown error
}
