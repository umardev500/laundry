package dto

import (
	"github.com/umardev500/laundry/internal/feature/paymentmethod/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// CreatePaymentMethodRequest is the payload for creating a new payment method
type CreatePaymentMethodRequest struct {
	Name        string              `json:"name" validate:"required"`
	Description *string             `json:"description,omitempty"`
	Type        types.PaymentMethod `json:"type" validate:"required,oneof=cash card transfer"`
}

func (r CreatePaymentMethodRequest) ToDomain() *domain.PaymentMethod {
	pm := &domain.PaymentMethod{
		Name:        r.Name,
		Description: r.Description,
		Type:        r.Type,
	}

	pm.InitDefaults()
	return pm
}
