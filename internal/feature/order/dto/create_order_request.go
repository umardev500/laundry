package dto

import "time"

type CreateOrderRequest struct {
	// Guest info (required only if user is guest)
	GuestName    *string `json:"guest_name,omitempty" binding:"omitempty,min=2,max=100"`
	GuestEmail   *string `json:"guest_email,omitempty" binding:"omitempty,email"`
	GuestPhone   *string `json:"guest_phone,omitempty" binding:"omitempty,e164"` // E.164 phone format
	GuestAddress *string `json:"guest_address,omitempty" binding:"omitempty,min=5,max=200"`

	PickupDate   *time.Time `json:"pickup_date,omitempty" binding:"omitempty,gt"` // gt ensures date is in the future
	DeliveryDate *time.Time `json:"delivery_date,omitempty" binding:"omitempty,gtfield=PickupDate"`

	Items []CreateOrderItemRequest `json:"items" binding:"required,min=1"`
}
