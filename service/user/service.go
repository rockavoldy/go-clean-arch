package user

import (
	"go-clean-arch/entity"
	"time"
)

type Service struct {
	repo Repository
}

// create NewService
func NewService(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

// create User
func (s *Service) CreateUser(email, password, name string) (entity.ID, error) {
	e, err := entity.NewUser(email, password, name)
	if err != nil {
		return e.ID, err
	}

	return s.repo.Create(e)
}

// Get User
func (s *Service) GetUser(id entity.ID) (*entity.User, error) {
	return s.repo.Get(id)
}

// Get User lists
func (s *Service) ListUsers() ([]*entity.User, error) {
	return s.repo.List()
}

// Search User
func (s *Service) SearchUsers(query string) ([]*entity.User, error) {
	return s.repo.Search(query)
}

// Update User
func (s *Service) UpdateUser(e *entity.User) error {
	err := e.ValidateInput()
	if err != nil {
		return err
	}

	e.UpdatedAt = time.Now()
	return s.repo.Update(e)
}

// Delete User
func (s *Service) DeleteUser(id entity.ID) error {
	u, err := s.GetUser(id)
	if u == nil {
		return entity.ErrNotFound
	}

	if err != nil {
		return err
	}

	if len(u.Books) > 0 {
		return entity.ErrCannotBeDeleted
	}

	return s.repo.Delete(id)
}

// Restore deleted User
func (s *Service) RestoreUser(id entity.ID) error {
	u, err := s.GetUser(id)
	if u == nil {
		return entity.ErrNotFound
	}

	if err != nil {
		return err
	}

	return s.repo.Restore(id)
}