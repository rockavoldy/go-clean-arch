package user

import (
	"github.com/stretchr/testify/assert"
	"go-clean-arch/entity"
	"testing"
	"time"
)

func newTestUser() *entity.User {
	return &entity.User{
		ID: entity.NewID(),
		Email: "fuu@fuu.com",
		Password: "new_pass",
		Name: "Fuu",
		CreatedAt: time.Now(),
	}
}

func TestService_CreateUser(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newTestUser()

	_, err := m.CreateUser(u.Email, u.Password, u.Name)

	assert.Nil(t, err)
	assert.False(t, u.CreatedAt.IsZero())
	assert.True(t, u.UpdatedAt.IsZero())
	assert.True(t, u.DeletedAt.IsZero())
}

func TestService_SearchAndFind(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newTestUser()
	u2 := newTestUser()
	u2.Name = "Fii"

	userId, _ := m.CreateUser(u1.Email, u1. Password, u1.Name)
	_, _ = m.CreateUser(u2.Email, u2.Password, u2.Name)

	t.Run("Search", func(t *testing.T) {
		c, err := m.SearchUsers("fuu")
		assert.Nil(t, err)
		assert.Equal(t, 1, len(c))
		assert.Equal(t, "fuu@fuu.com", c[0].Email)

		c, err = m.SearchUsers("faa")
		assert.Equal(t, entity.ErrNotFound, err)
		assert.Nil(t, c)
	})

	t.Run("List all users", func(t *testing.T) {
		all, err := m.ListUsers()
		assert.Nil(t, err)
		assert.Equal(t, 2, len(all))
	})

	t.Run("Get user", func(t *testing.T) {
		saved, err := m.GetUser(userId)
		assert.Nil(t, err)
		assert.Equal(t, u1.Name, saved.Name)
	})
}

func TestService_UpdateUser(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u := newTestUser()

	id, err := m.CreateUser(u.Email, u.Password, u.Name)

	assert.Nil(t, err)
	saved, _ := m.GetUser(id)
	saved.Name = "Fee"
	saved.Books = append(saved.Books, entity.NewID())
	assert.Nil(t, m.UpdateUser(saved))

	updated, err := m.GetUser(id)
	assert.Nil(t, err)
	assert.Equal(t, "Fee", updated.Name)
	assert.False(t, updated.UpdatedAt.IsZero())
	assert.Equal(t, 1, len(updated.Books))
}

func TestService_DeleteUser(t *testing.T) {
	repo := newInmem()
	m := NewService(repo)
	u1 := newTestUser()
	u2 := newTestUser()

	u2Id, _ := m.CreateUser(u2.Email, u2.Password, u2.Name)

	err := m.DeleteUser(u1.ID)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteUser(u2Id)
	assert.Nil(t, err)
	_, err = m.GetUser(u2Id)
	assert.Equal(t, entity.ErrNotFound, err)

	err = m.DeleteUser(u2Id)
	assert.Equal(t, entity.ErrNotFound, err)

	u3 := newTestUser()
	id, _ := m.CreateUser(u3.Email, u3.Password, u3.Name)
	saved, _ := m.GetUser(id)
	saved.Books = []entity.ID{entity.NewID()}
	_ = m.UpdateUser(saved)
	err = m.DeleteUser(id)
	assert.Equal(t, entity.ErrCannotBeDeleted, err)
}