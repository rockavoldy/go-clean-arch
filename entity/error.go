package entity

import "errors"

// Not Found
var ErrNotFound = errors.New("not found")

// Invalid Entity
var ErrInvalidEntity = errors.New("invalid entity")

// Cannot be Deleted
var ErrCannotBeDeleted = errors.New("cannot be deleted")

// Cannot restore
var ErrCannotBeRestored = errors.New("cannot be restored")

// Already exist
var ErrAlreadyExist = errors.New("already exists")

// Has been deleted
var ErrAlreadyDeleted = errors.New("already deleted")

// Not enough Book
var ErrNotEnoughBook = errors.New("not enough book")

// Book not borrowed by *this* User
var ErrBookNotBorrowed = errors.New("book not borrowed by this user")

// Book already borrowed by this User
var ErrBookAlreadyBorrowed = errors.New("book already borrowed")