package dto

import (
	"time"

	"github.com/google/uuid"
)

// PaymentMethodResponse is the response DTO for returning payment method details
type PaymentMethodResponse struct {
	ID          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}
