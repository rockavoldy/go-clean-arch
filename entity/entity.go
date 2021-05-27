package entity

import "github.com/google/uuid"

//ID entity ID
type ID = uuid.UUID

// FormatDateTimeSQL
// TODO: move to other place next time
const FormatDateTimeSQL = "2006-01-02 03:04:05"

//NewID Create a new entity ID
func NewID() ID {
	return uuid.New()
}

//StringToID convert string to entity ID
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)

	return id, err
}