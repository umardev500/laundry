package dto

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/plan/domain"
	"github.com/umardev500/laundry/pkg/types"
)

// UpdatePlanRequest represents the payload for updating a Plan.
type UpdatePlanRequest struct {
	Name            string                `json:"name,omitempty" validate:"omitempty,min=2,max=200"`
	Description     string                `json:"description,omitempty" validate:"omitempty,max=1024"`
	Price           float64               `json:"price,omitempty"`
	BillingInterval types.BillingInterval `json:"billing_interval,omitempty"`
	Features        *PlanFeaturesDTO      `json:"features,omitempty"`
}

// ToDomain converts the UpdatePlanRequest into a domain.Plan with the given ID.
// Only non-nil fields will be updated.
func (r *UpdatePlanRequest) ToDomain(ctx *appctx.Context, id uuid.UUID) *domain.Plan {
	return &domain.Plan{
		ID:              id,
		Name:            r.Name,
		Description:     &r.Description,
		Price:           r.Price,
		BillingInterval: r.BillingInterval,
		Features:        r.Features.ToDomain(),
	}
}
