package loan

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go-clean-arch/entity"
	"testing"
	umock "go-clean-arch/service/user/mock"
	bmock "go-clean-arch/service/book/mock"
)

func TestService_Borrow(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	userMock := umock.NewMockUseCase(controller)
	bookMock := bmock.NewMockUseCase(controller)

	uc := NewService(userMock, bookMock)

	t.Run("User not found", func(t *testing.T) {
		u := &entity.User{
			ID: entity.NewID(),
		}
		b := &entity.Book{
			ID: entity.NewID(),
		}

		userMock.EXPECT().GetUser(u.ID).Return(nil, entity.ErrNotFound)
		err := uc.Borrow(u, b)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("Book not found", func(t *testing.T) {
		u := &entity.User{
			ID: entity.NewID(),
		}
		b := &entity.Book{
			ID: entity.NewID(),
		}

		userMock.EXPECT().GetUser(u.ID).Return(u, nil)
		bookMock.EXPECT().GetBook(b.ID).Return(nil, entity.ErrNotFound)
		err := uc.Borrow(u, b)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("Not enough book to borrow", func(t *testing.T) {
		u := &entity.User{
			ID: entity.NewID(),
		}
		b := &entity.Book{
			ID: entity.NewID(),
		}
		b.Quantity = 0

		userMock.EXPECT().GetUser(u.ID).Return(u, nil)
		bookMock.EXPECT().GetBook(b.ID).Return(b, nil)
		err := uc.Borrow(u, b)
		assert.Equal(t, entity.ErrNotEnoughBook, err)
	})

	t.Run("Book already borrowed", func(t *testing.T) {
		u := &entity.User{
			ID: entity.NewID(),
		}
		b := &entity.Book{
			ID: entity.NewID(),
		}
		u.AddBook(b.ID)
		b.Quantity = 1

		userMock.EXPECT().GetUser(u.ID).Return(u, nil)
		bookMock.EXPECT().GetBook(b.ID).Return(b, nil)

		err := uc.Borrow(u, b)
		assert.Equal(t, entity.ErrBookAlreadyBorrowed, err)
	})

	t.Run("Success", func(t *testing.T) {
		u := &entity.User{
			ID: entity.NewID(),
		}
		b := &entity.Book{
			ID: entity.NewID(),
			Quantity: 10,
		}

		userMock.EXPECT().GetUser(u.ID).Return(u, nil)
		bookMock.EXPECT().GetBook(b.ID).Return(b, nil)
		userMock.EXPECT().UpdateUser(u).Return(nil)
		bookMock.EXPECT().UpdateBook(b).Return(nil)
		err := uc.Borrow(u, b)
		assert.Nil(t, err)
	})
}

func TestService_Return(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	userMock := umock.NewMockUseCase(controller)
	bookMock := bmock.NewMockUseCase(controller)
	uc := NewService(userMock, bookMock)

	t.Run("book not found", func(t *testing.T) {
		u := &entity.User{
			ID: entity.NewID(),
		}
		b := &entity.Book{
			ID: entity.NewID(),
		}

		bookMock.EXPECT().GetBook(b.ID).Return(nil, entity.ErrNotFound)
		err := uc.Return(u, b)
		assert.Equal(t, entity.ErrNotFound, err)
	})

	t.Run("book not borrowed by this user", func(t *testing.T) {
		u := &entity.User{
			ID: entity.NewID(),
		}
		b := &entity.Book{
			ID: entity.NewID(),
		}

		userMock.EXPECT().GetUser(u.ID).Return(u, nil)
		bookMock.EXPECT().GetBook(b.ID).Return(b, nil)
		err := uc.Return(u, b)
		assert.Equal(t, entity.ErrBookNotBorrowed, err)
	})

	t.Run("Success", func(t *testing.T) {
		u := &entity.User{
			ID: entity.NewID(),
		}
		b := &entity.Book{
			ID: entity.NewID(),
		}

		u.AddBook(b.ID)

		userMock.EXPECT().GetUser(u.ID).Return(u, nil)
		bookMock.EXPECT().GetBook(b.ID).Return(b, nil)

		userMock.EXPECT().UpdateUser(u).Return(nil)
		bookMock.EXPECT().UpdateBook(b).Return(nil)
		err := uc.Return(u, b)
		assert.Nil(t, err)
	})
}