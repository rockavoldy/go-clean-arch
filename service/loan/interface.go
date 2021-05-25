package loan

import "go-clean-arch/entity"

// UseCase interface
type UseCase interface {
	Borrow(u *entity.User, b *entity.Book) error
	Return(u *entity.User, b *entity.Book) error
}