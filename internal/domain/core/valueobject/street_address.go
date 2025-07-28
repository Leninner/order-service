package valueobject

import "github.com/google/uuid"

type StreetAddress struct {
	ID         uuid.UUID
	Street     string
	PostalCode string
	City       string
}
