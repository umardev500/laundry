package dto

import "github.com/google/uuid"

// OrderItemResponse represents the order item returned in API responses.
type OrderItemResponse struct {
	ID          uuid.UUID `json:"id"`
	OrderID     uuid.UUID `json:"order_id"`
	ServiceID   uuid.UUID `json:"service_id"`
	Quantity    float64   `json:"quantity"`
	Price       float64   `json:"price"`
	Subtotal    float64   `json:"subtotal"`
	TotalAmount float64   `json:"total_amount"`
}
