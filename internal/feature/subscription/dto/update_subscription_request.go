package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/subscription/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// UpdateSubscriptionRequest represents fields that can be updated in a subscription.
type UpdateSubscriptionRequest struct {
	PlanID    *uuid.UUID                `json:"plan_id,omitempty"`
	Status    *types.SubscriptionStatus `json:"status,omitempty"`
	StartDate *time.Time                `json:"start_date,omitempty"`
	EndDate   *time.Time                `json:"end_date,omitempty"`
}

// ToDomain converts the UpdateSubscriptionRequest into a domain.Subscription for update operations.
func (r *UpdateSubscriptionRequest) ToDomain(ctx *appctx.Context, id uuid.UUID) *domain.Subscription {
	sub := &domain.Subscription{
		ID: id,
	}

	if r.PlanID != nil {
		sub.PlanID = *r.PlanID
	}

	if r.Status != nil {
		sub.Status = *r.Status
	}

	if r.StartDate != nil {
		sub.StartDate = r.StartDate
	}

	if r.EndDate != nil {
		sub.EndDate = r.EndDate
	}

	return sub
}
