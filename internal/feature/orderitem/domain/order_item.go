package domain

import (
	"errors"

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

// Validate ensures the order item has valid values before saving.
func (i *OrderItem) Validate() error {
	if i == nil {
		return errors.New("order item cannot be nil")
	}

	if i.OrderID == uuid.Nil {
		return errors.New("order ID is required")
	}

	if i.ServiceID == uuid.Nil {
		return errors.New("service ID is required")
	}

	if i.Quantity <= 0 {
		return errors.New("quantity must be greater than zero")
	}

	if i.Price < 0 {
		return errors.New("price cannot be negative")
	}

	return nil
}

// CalculateTotals recalculates the Subtotal and TotalAmount based on Quantity and Price.
func (i *OrderItem) CalculateTotals() {
	if i == nil {
		return
	}

	i.Subtotal = i.Price * i.Quantity
	i.TotalAmount = i.Subtotal // 💡 If you later add taxes or discounts, adjust here
}
