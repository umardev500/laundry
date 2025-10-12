package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// PaymentResponse represents a Payment for API responses.
type PaymentResponse struct {
	ID              uuid.UUID           `json:"id"`
	RefID           uuid.UUID           `json:"ref_id"`
	RefType         types.PaymentType   `json:"ref_type"`
	PaymentMethodID uuid.UUID           `json:"payment_method_id,omitempty"`
	Amount          float64             `json:"amount"`
	ReceivedAmount  *float64            `json:"received_amount,omitempty"`
	ChangeAmount    *float64            `json:"change_amount,omitempty"`
	Notes           string              `json:"notes,omitempty"`
	Status          types.PaymentStatus `json:"status"`
	PaidAt          *time.Time          `json:"paid_at,omitempty"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
	DeletedAt       *time.Time          `json:"deleted_at,omitempty"`
}
