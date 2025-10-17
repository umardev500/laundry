package dto

import (
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/subscription/domain"
)

// CreateSubscriptionRequest represents the input payload for creating a subscription.
type CreateSubscriptionRequest struct {
	TenantID  uuid.UUID  `json:"tenant_id" validate:"required"`
	PlanID    uuid.UUID  `json:"plan_id" validate:"required"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}

// ToDomain converts the CreateSubscriptionRequest into a domain.Subscription.
func (r *CreateSubscriptionRequest) ToDomain(ctx *appctx.Context) *domain.Subscription {

	return domain.NewSubscription(
		r.TenantID,
		r.PlanID,
		r.StartDate,
		r.EndDate,
	)
}
