package query

import (
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
