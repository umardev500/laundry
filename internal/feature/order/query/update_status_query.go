package query

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type UpdateStatusQuery struct {
	Status types.OrderStatus `params:"status"`
}

func (q *UpdateStatusQuery) ToDomain(id uuid.UUID) (*domain.Order, error) {
	return &domain.Order{
		ID:     id,
		Status: q.Status,
	}, nil
}

func (q *UpdateStatusQuery) Validate() error {
	if q.Status == "" {
		return errors.New("status is required")
	}

	// Normalize input to uppercase for case-insensitive comparison
	q.Status = types.OrderStatus(strings.ToUpper(string(q.Status)))

	allowed := []types.OrderStatus{
		types.OrderStatusPending,
		types.OrderStatusConfirmed,
		types.OrderStatusPickedUp,
		types.OrderStatusInWashing,
		types.OrderStatusInDrying,
		types.OrderStatusInIroning,
		types.OrderStatusReadyForDelivery,
		types.OrderStatusOutForDelivery,
		types.OrderStatusDelivered,
		types.OrderStatusCompleted,
		types.OrderStatusCancelled,
		types.OrderStatusFailed,
	}

	if slices.Contains(allowed, q.Status) {
		return nil // ✅ valid
	}

	// Build a readable list of allowed statuses
	available := make([]string, len(allowed))
	for i, s := range allowed {
		available[i] = string(s)
	}

	return fmt.Errorf("invalid status: %s. available statuses are: %s", q.Status, strings.Join(available, ", "))
}
