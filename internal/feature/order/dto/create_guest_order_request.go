package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/order/domain"
)

type CreateGuestOrderRequest struct {
	Name    string  `json:"name" validate:"required"`
	Email   *string `json:"email,omitempty" validate:"omitempty,email"`
	Phone   *string `json:"phone,omitempty" validate:"omitempty,e164"`
	Address string  `json:"address" validate:"required,min=5,max=200"`

	Items []CreateOrderItemRequest `json:"items" validate:"required,min=1"`
}

func (r *CreateGuestOrderRequest) Validate() error {
	if (r.Email == nil || *r.Email == "") && (r.Phone == nil || *r.Phone == "") {
		return domain.ErrGuestEmailOrPhoneRequired
	}

	return nil
}

func (r *CreateGuestOrderRequest) ToDomain(tenantID uuid.UUID) (*domain.Order, error) {
	var items []*domain.OrderItem
	for _, i := range r.Items {
		item, err := i.ToDomain()
		if err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return &domain.Order{
		TenantID:     tenantID,
		GuestName:    &r.Name,
		GuestEmail:   r.Email,
		GuestPhone:   r.Phone,
		GuestAddress: &r.Address,
		Items:        items,
	}, nil
}
