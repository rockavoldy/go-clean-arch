package user

import "go-clean-arch/entity"

// Reader interface
type Reader interface {
	Get(id entity.ID) (*entity.User, error)
	List() ([]*entity.User, error)
	Search(query string) ([]*entity.User, error)
}

// Writer interface
type Writer interface {
	Create(e *entity.User) (entity.ID, error)
	Update(e *entity.User) error
	Delete(id entity.ID) error
	Restore(id entity.ID) error
}

type Repository interface {
	Reader
	Writer
}

type UseCase interface {
	GetUser(id entity.ID) (*entity.User, error)
	ListUsers() ([]*entity.User, error)
	SearchUsers(query string) ([]*entity.User, error)
	CreateUser(email, password, name string) (entity.ID, error)
	UpdateUser(e *entity.User) error
	DeleteUser(id entity.ID) error
	RestoreUser(id entity.ID) error
}