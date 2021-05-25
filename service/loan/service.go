package loan

import (
	"go-clean-arch/entity"
	"go-clean-arch/service/book"
	"go-clean-arch/service/user"
)

type Service struct {
	userService user.UseCase
	bookService book.UseCase
}

// create NewService
func NewService(u user.UseCase, b book.UseCase) *Service {
	return &Service{
		userService: u,
		bookService: b,
	}
}

// Borrow a Book
func (s *Service) Borrow(u *entity.User, b *entity.Book) error {
	u, err := s.userService.GetUser(u.ID)
	if err != nil {
		return err
	}

	b, err = s.bookService.GetBook(b.ID)
	if err != nil {
		return err
	}

	if b.Quantity <= 0 {
		return entity.ErrNotEnoughBook
	}

	err = u.AddBook(b.ID)
	if err != nil {
		return err
	}

	err = s.userService.UpdateUser(u)
	if err != nil {
		return err
	}

	b.Quantity--
	err = s.bookService.UpdateBook(b)
	if err != nil {
		return err
	}

	return nil
}

// Return a Book
func (s *Service) Return(u *entity.User, b *entity.Book) error {
	b, err := s.bookService.GetBook(b.ID)
	if err != nil {
		return err
	}

	u, err = s.userService.GetUser(u.ID)
	if err != nil {
		return err
	}

	bookBorrowedId, err := u.GetBook(b.ID)
	if err != nil {
		return entity.ErrBookNotBorrowed
	}

	err = u.RemoveBook(bookBorrowedId)
	if err != nil {
		return err
	}

	err = s.userService.UpdateUser(u)
	if err != nil {
		return err
	}

	b.Quantity++
	err = s.bookService.UpdateBook(b)
	if err != nil {
		return err
	}

	return nil
}