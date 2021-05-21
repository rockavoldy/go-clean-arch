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