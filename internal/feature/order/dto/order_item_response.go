package dto

import "github.com/google/uuid"

type OrderItemResponse struct {
	ID        uuid.UUID `json:"id"`
	OrderID   uuid.UUID `json:"order_id"`
	ServiceID uuid.UUID `json:"service_id"`
	Quantity  float64   `json:"quantity"`
	Price     float64   `json:"price"`
	Total     float64   `json:"total"`
}
