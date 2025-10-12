package dto

import "github.com/umardev500/laundry/internal/feature/paymentmethod/domain"

// CreatePaymentMethodRequest is the payload for creating a new payment method
type CreatePaymentMethodRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description *string `json:"description,omitempty"`
}

func (r CreatePaymentMethodRequest) ToDomain() *domain.PaymentMethod {
	pm := &domain.PaymentMethod{
		Name:        r.Name,
		Description: r.Description,
	}

	pm.InitDefaults()
	return pm
}
