package query

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/subscription/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type UpdateStatusQuery struct {
	Status string `json:"status" query:"status" params:"status"`
}

func (q *UpdateStatusQuery) ToDomain(id uuid.UUID) (*domain.Subscription, error) {
	return &domain.Subscription{
		ID:     id,
		Status: types.SubscriptionStatus(q.Status),
	}, nil
}
