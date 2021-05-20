package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

//User data
type User struct {
	ID        ID
	Email     string
	Password  string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Create a new User
func NewUser(email string, password string, name string) (*User, error) {
	user := &User{
		ID: NewID(),
		Email: email,
		Password: password,
		Name: name,
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

// Validate input entity User
func (user *User) ValidateInput() error {
	if user.Email == "" || user.Password == "" || user.Name == "" {
		return errors.New("'Name' or 'Email' or 'Password' is empty")
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