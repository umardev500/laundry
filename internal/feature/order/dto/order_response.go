package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

type OrderResponse struct {
	ID           uuid.UUID            `json:"id"`
	TenantID     uuid.UUID            `json:"tenant_id"`
	UserID       *uuid.UUID           `json:"user_id,omitempty"`
	Status       types.OrderStatus    `json:"status"`
	TotalAmount  float64              `json:"total_amount"`
	GuestName    *string              `json:"guest_name,omitempty"`
	GuestEmail   *string              `json:"guest_email,omitempty"`
	GuestPhone   *string              `json:"guest_phone,omitempty"`
	GuestAddress *string              `json:"guest_address,omitempty"`
	CreatedAt    time.Time            `json:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at"`
	DeletedAt    *time.Time           `json:"deleted_at,omitempty"`
	Items        []*OrderItemResponse `json:"items"`
}
