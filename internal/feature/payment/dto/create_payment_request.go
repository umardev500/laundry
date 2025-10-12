package dto

import (
	"github.com/google/uuid"
)

type CreatePaymentRequest struct {
	PaymentMethodID uuid.UUID `json:"payment_method_id" validate:"required"`
	ReceivedAmount  *float64  `json:"received_amount,omitempty" validate:"omitempty,gte=0"`
	Notes           *string   `json:"notes,omitempty" validate:"omitempty,max=255"`
}
