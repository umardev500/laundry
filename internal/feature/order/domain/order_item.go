package domain

import (
	"github.com/google/uuid"
)

type OrderItem struct {
	ID          uuid.UUID
	OrderID     uuid.UUID
	ServiceID   uuid.UUID
	Quantity    float64
	Price       float64
	Subtotal    float64
	TotalAmount float64
}
