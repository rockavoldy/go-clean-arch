package entity

import "errors"

// Not Found
var ErrNotFound = errors.New("Not found")

// Invalid Entity
var ErrInvalidEntity = errors.New("Invalid Entity")

// Cannot be Deleted
var ErrCannotBeDeleted = errors.New("Cannot be Deleted")

// Cannot restore
var ErrCannotBeRestored = errors.New("Cannot be Restored")

// Already exist
var ErrAlreadyExist = errors.New("Already Exist")

// Has been deleted
var ErrBeenDeleted = errors.New("Has Been Deleted")