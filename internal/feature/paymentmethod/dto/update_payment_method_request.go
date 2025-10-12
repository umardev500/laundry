package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// UpdatePaymentMethodRequest is the payload for updating a payment method
type UpdatePaymentMethodRequest struct {
	Name        string              `json:"name"`
	Description *string             `json:"description,omitempty"`
	Type        types.PaymentMethod `json:"type"`
}

func (r UpdatePaymentMethodRequest) ToDomain(id uuid.UUID) *domain.PaymentMethod {
	pm := &domain.PaymentMethod{
		ID:          id,
		Name:        r.Name,
		Description: r.Description,
		Type:        r.Type,
	}

	return pm
}
