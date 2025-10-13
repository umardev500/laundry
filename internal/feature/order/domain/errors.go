package domain

import (
	"fmt"

	"github.com/google/uuid"
)

var (
	ErrGuestEmailOrPhoneRequired = fmt.Errorf("guest email or phone is required")
	ErrOrderItemsRequired        = fmt.Errorf("order items are required")
	ErrOrderNotFound             = fmt.Errorf("order not found")
	ErrOrderDeleted              = fmt.Errorf("order has been deleted")
	ErrUnauthorizedOrderAccess   = fmt.Errorf("unauthorized access to order")
)

// ServiceUnavailableError is an error that occurs when one or more services are unavailable.
type ServiceUnavailableError struct {
	UnavailableIDs []uuid.UUID
}

func (e *ServiceUnavailableError) Error() string {
	return fmt.Sprintf("services not available: %v", e.UnavailableIDs)
}

func NewServiceUnavailableError(ids []uuid.UUID) *ServiceUnavailableError {
	return &ServiceUnavailableError{UnavailableIDs: ids}
}
