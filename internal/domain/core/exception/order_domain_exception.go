package exception

import (
	"github.com/leninner/shared/domain/exception"
)

func NewOrderDomainException(message string) error {
	return exception.NewDomainException(message)
}
