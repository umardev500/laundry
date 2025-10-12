package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/order/domain"

	orderItemDomain "github.com/umardev500/laundry/internal/feature/orderitem/domain"
	paymentDomain "github.com/umardev500/laundry/internal/feature/payment/domain"
	paymentDto "github.com/umardev500/laundry/internal/feature/payment/dto"
)

type CreateGuestOrderRequest struct {
	Name    string  `json:"name" validate:"required"`
	Email   *string `json:"email,omitempty" validate:"omitempty,email"`
	Phone   *string `json:"phone,omitempty" validate:"omitempty,e164"`
	Address string  `json:"address" validate:"required,min=5,max=200"`
	Notes   *string `json:"notes,omitempty" validate:"omitempty,max=255"`

	Items []CreateOrderItemRequest `json:"items" validate:"required,min=1"`

	Payment *paymentDto.CreatePaymentRequest `json:"payment" validate:"required"`
}

func (r *CreateGuestOrderRequest) Validate() error {
	if (r.Email == nil || *r.Email == "") && (r.Phone == nil || *r.Phone == "") {
		return domain.ErrGuestEmailOrPhoneRequired
	}

	return nil
}

func (r *CreateGuestOrderRequest) ToDomain(tenantID uuid.UUID) (*domain.Order, error) {
	var items []*orderItemDomain.OrderItem
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
		Payment: &paymentDomain.Payment{
			PaymentMethodID: r.Payment.PaymentMethodID,
			ReceivedAmount:  r.Payment.ReceivedAmount,
			Notes:           *r.Payment.Notes,
		},
	}, nil
}
