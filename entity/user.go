package entity

import (
	"golang.org/x/crypto/bcrypt"
	"time"
)

// User structure
type User struct {
	ID        ID
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Books     []ID
}

// Create a new User
func NewUser(email string, password string, name string) (*User, error) {
	user := &User{
		ID:        NewID(),
		Email:     email,
		Password:  password,
		Name:      name,
		CreatedAt: time.Now(),
	}

	pwd, err := hashPasswordBcrypt(password)
	if err != nil {
		return nil, err
	}

	user.Password = pwd
	err = user.ValidateInput()

	if err != nil {
		return nil, err
	}

	return user, nil
}

// Add book to User
func (user *User) AddBook(id ID) error {
	_, err := user.GetBook(id)
	if err == nil {
		return ErrAlreadyExist
	}

	user.Books = append(user.Books, id)

	return nil
}

// Remove book from User
func (user *User) RemoveBook(id ID) error {
	for i, val := range user.Books {
		if val == id {
			user.Books = append(user.Books[:i], user.Books[:i+1]...)
			return nil
		}
	}

	return ErrNotFound
}

// Get book from User
func (user *User) GetBook(id ID) (ID, error) {
	for _, val := range user.Books {
		if val == id {
			return id, nil
		}
	}

	return id, ErrNotFound
}

// Validate input entity User
func (user *User) ValidateInput() error {
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return ErrInvalidEntity
	}

	return nil
}

// Validate Hashed password
func (user *User) ValidatePassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	}

	return nil
}

// Generate hash from password input
func hashPasswordBcrypt(raw string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
