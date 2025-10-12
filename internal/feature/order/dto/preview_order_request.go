package dto

import (
	"github.com/umardev500/laundry/internal/feature/order/domain"
	orderItemDomain "github.com/umardev500/laundry/internal/feature/orderitem/domain"
)

type PreviewOrderRequest struct {
	Items []CreateOrderItemRequest `json:"items" validate:"required,min=1"`
}

// Validate basic structure
func (r *PreviewOrderRequest) Validate() error {
	if len(r.Items) == 0 {
		return domain.ErrOrderItemsRequired
	}
	return nil
}

// Convert to domain Order (without guest or payment info)
func (r *PreviewOrderRequest) ToDomain() (*domain.Order, error) {
	var items []*orderItemDomain.OrderItem
	for _, i := range r.Items {
		item, err := i.ToDomain()
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &domain.Order{
		Items: items,
	}, nil
}
