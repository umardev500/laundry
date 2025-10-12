package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"

	orderItemDomain "github.com/umardev500/laundry/internal/feature/orderitem/domain"
	paymentDomain "github.com/umardev500/laundry/internal/feature/payment/domain"
	serviceDomain "github.com/umardev500/laundry/internal/feature/service/domain"
)

type Order struct {
	ID           uuid.UUID
	TenantID     uuid.UUID
	UserID       *uuid.UUID
	Status       types.OrderStatus
	TotalAmount  float64
	Notes        *string
	GuestName    *string
	GuestEmail   *string
	GuestPhone   *string
	GuestAddress *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    *time.Time
	Items        []*orderItemDomain.OrderItem

	Payment *paymentDomain.Payment
}

// AttachOrderID assigns the given order ID to all order items.
func (o *Order) AttachOrderID(orderID uuid.UUID) {
	if o == nil || len(o.Items) == 0 {
		return
	}

	for _, item := range o.Items {
		if item != nil {
			item.OrderID = orderID
		}
	}
}

func (o *Order) Place(availableServices []*serviceDomain.Service) error {
	if o == nil {
		return fmt.Errorf("order cannot be nil")
	}

	// Build look up for service pricing
	serviceMap := make(map[uuid.UUID]*serviceDomain.Service, len(availableServices))
	for _, s := range availableServices {
		serviceMap[s.ID] = s
	}

	// Calculate totals
	var total float64
	for _, item := range o.Items {
		svc, exists := serviceMap[item.ServiceID]
		if !exists {
			return fmt.Errorf("service %s not found", item.ServiceID)
		}

		item.Price = svc.BasePrice
		item.CalculateTotals()
		total += item.Subtotal
	}

	o.TotalAmount = total

	// Init defaults
	o.InitDefaults()

	return nil
}

func (o *Order) Validate() error {
	if o.TenantID == uuid.Nil {
		return types.ErrTenantIDRequired
	}

	return nil
}

func (o *Order) InitDefaults() {
	if o.Status == "" {
		o.Status = types.OrderStatusPending
	}
}

// GetServiceIDs returns a list of service IDs from the order items.
func (o *Order) GetServiceIDs() []uuid.UUID {
	if o == nil || len(o.Items) == 0 {
		return []uuid.UUID{}
	}

	ids := make([]uuid.UUID, 0, len(o.Items))
	for _, item := range o.Items {
		if item != nil {
			ids = append(ids, item.ServiceID)
		}
	}
	return ids
}
