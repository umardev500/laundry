package domain

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// OrderStatusHistory represents a status milestone for an order.
type OrderStatusHistory struct {
	ID        uuid.UUID
	OrderID   uuid.UUID
	Status    types.OrderStatus
	Notes     *string
	CreatedAt time.Time
	Order     any
}
