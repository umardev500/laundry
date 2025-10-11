package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

type Order struct {
	ID           uuid.UUID
	TenantID     uuid.UUID
	UserID       *uuid.UUID
	Status       types.OrderStatus
	TotalAmount  float64
	GuestName    *string
	GuestEmail   *string
	GuestPhone   *string
	GuestAddress *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Items        []*OrderItem
}
