package entity_test

import (
	"github.com/stretchr/testify/assert"
	"go-clean-arch/entity"
	"testing"
)

var testUserData = map[string]string{
	"email": "fuu@fuu.com",
	"password": "new_pass",
	"name": "Fuu",
}

// Test NewUser function in User entity
func TestNewUser(t *testing.T) {
	user, err := entity.NewUser(testUserData["email"],testUserData["password"], testUserData["name"])

	assert.Nil(t, err) // check if no error thrown
	assert.NotNil(t, user.ID) // check if ID generated
	assert.Equal(t, user.Name, testUserData["name"]) // check if name equal to input
	assert.Equal(t, user.Email, testUserData["email"])  // check if email equal to input
	assert.NotEqual(t, user.Password, testUserData["password"]) // check if password successfully hashed
}

// Test Validate function in User entity
func TestUser_ValidatePassword(t *testing.T) {
	user, _ := entity.NewUser(testUserData["email"],testUserData["password"], testUserData["name"])
	err := user.ValidatePassword(testUserData["password"])
	assert.Nil(t, err) // ValidatePassword should NOT thrown error

	err = user.ValidatePassword("wrong_password")
	assert.NotNil(t, err) // ValidatePassword SHOULD thrown error
}