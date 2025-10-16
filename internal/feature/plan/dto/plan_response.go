package dto

import (
	"time"

	"github.com/google/uuid"
)

// PlanResponse represents the Plan data returned to clients.
type PlanResponse struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	Description     *string   `json:"description,omitempty"`
	Price           float64   `json:"price"`
	BillingInterval string    `json:"billing_interval"`
	Features        any       `json:"features,omitempty"` // structured features
	Active          bool      `json:"active"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
