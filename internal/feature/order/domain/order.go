package domain

import (
	"fmt"
	"slices"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/errors"
	"github.com/umardev500/laundry/pkg/types"

	"github.com/umardev500/laundry/internal/app/appctx"
	orderItemDomain "github.com/umardev500/laundry/internal/feature/orderitem/domain"
	orderStatusHistoryDomain "github.com/umardev500/laundry/internal/feature/orderstatushistory/domain"
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

	Payment  *paymentDomain.Payment
	Statuses []*orderStatusHistoryDomain.OrderStatusHistory
}

func (o *Order) UpdateStatus(newStatus types.OrderStatus) error {
	if o == nil {
		return fmt.Errorf("order cannot be nil")
	}

	// Prevent updateing to the same status
	if o.Status == newStatus {
		return fmt.Errorf("order is already in status %s", newStatus)
	}

	// Terminal states cannot be changed
	if o.Status == types.OrderStatusCancelled || o.Status == types.OrderStatusCompleted {
		return errors.NewErrInvalidStatusTransition(string(o.Status), string(newStatus))
	}

	// Check allowed transitions
	allowedNext, ok := types.AllowedOrderTransitions[o.Status]
	if !ok {
		return errors.NewErrInvalidStatusTransition(string(o.Status), string(newStatus))
	}

	valid := slices.Contains(allowedNext, newStatus)

	if !valid {
		return errors.NewErrInvalidStatusTransition(string(o.Status), string(newStatus))
	}

	// Apply new status
	o.Status = newStatus

	return nil
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

func (o *Order) IsGuestOrder() bool {
	return o.UserID == nil
}

func (o *Order) IsDeleted() bool {
	return o.DeletedAt != nil
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

// BelongsToTenant checks whether the service belongs to the tenant in context.
func (s *Order) BelongsToTenant(ctx *appctx.Context) bool {
	if ctx.Scope() == appctx.ScopeTenant {
		return ctx.TenantID() != nil && s.TenantID == *ctx.TenantID()
	}
	return true
}
