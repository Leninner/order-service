package customer

import (
	"github.com/google/uuid"
)

type CustomerModel struct {
	ID uuid.UUID `json:"id"`
}