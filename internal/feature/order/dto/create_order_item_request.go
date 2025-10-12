package dto

import (
	"github.com/google/uuid"
	orderItemDomain "github.com/umardev500/laundry/internal/feature/orderitem/domain"
)

type CreateOrderItemRequest struct {
	ServiceID uuid.UUID `json:"service_id" binding:"required"`
	Quantity  float64   `json:"quantity" binding:"required,min=1"` // At least 1
}

func (req CreateOrderItemRequest) ToDomain() (*orderItemDomain.OrderItem, error) {
	return &orderItemDomain.OrderItem{
		ServiceID: req.ServiceID,
		Quantity:  req.Quantity,
	}, nil
}
