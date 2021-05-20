package entity

import "github.com/google/uuid"

//ID entity ID
type ID = uuid.UUID

//NewID Create a new entity ID
func NewID() ID {
	return ID(uuid.New())
}

//StringToID convert string to entity ID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)

	return ID(id), err
}