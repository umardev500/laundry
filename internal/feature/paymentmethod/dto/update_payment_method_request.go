package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/domain"
)

// UpdatePaymentMethodRequest is the payload for updating a payment method
type UpdatePaymentMethodRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
}

func (r UpdatePaymentMethodRequest) ToDomain(id uuid.UUID) *domain.PaymentMethod {
	pm := &domain.PaymentMethod{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
	}

	return pm
}
