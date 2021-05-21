package entity

import "errors"

// Not Found
var ErrNotFound = errors.New("Not found")

// Invalid Entity
var ErrInvalidEntity = errors.New("Invalid Entity")

// Cannot be Deleted
var ErrCannotBeDeleted = errors.New("Cannot be Deleted")

// Already exist
var ErrAlreadyExist = errors.New("Already Exist")