package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// OrderStatusHistoryResponse represents a status milestone in API responses
type OrderStatusHistoryResponse struct {
	ID        uuid.UUID         `json:"id"`              // Milestone ID
	OrderID   uuid.UUID         `json:"order_id"`        // Related Order ID
	Status    types.OrderStatus `json:"status"`          // Status at this milestone
	Notes     *string           `json:"notes,omitempty"` // Optional notes
	CreatedAt time.Time         `json:"created_at"`      // Timestamp when the milestone was recorded
	Order     any               `json:"order,omitempty"` // Optional order details
}
