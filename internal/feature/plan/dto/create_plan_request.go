package dto

import (
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/plan/domain"
	"github.com/umardev500/laundry/pkg/types"
	"github.com/umardev500/laundry/pkg/utils"
)

type PlanFeaturesDTO struct {
	MaxUsers *int `json:"max_users,omitempty"`
}

// ToDomain converts PlanFeaturesDTO to domain.PlanFeatures.
func (f *PlanFeaturesDTO) ToDomain() *domain.PlanFeatures {
	if f == nil {
		return nil
	}
	return &domain.PlanFeatures{
		MaxUsers: f.MaxUsers,
	}
}

// CreatePlanRequest represents the input for creating a new Plan.
type CreatePlanRequest struct {
	Name            string           `json:"name" validate:"required,min=2,max=200"`
	Description     string           `json:"description,omitempty" validate:"omitempty,max=1024"`
	Price           float64          `json:"price,omitempty"`
	BillingInterval string           `json:"billing_interval,omitempty"`
	Features        *PlanFeaturesDTO `json:"features,omitempty"` // structured JSON
}

// ToDomain converts the CreatePlanRequest to a domain.Plan.
func (r *CreatePlanRequest) ToDomain(ctx *appctx.Context) *domain.Plan {
	return &domain.Plan{
		Name:            r.Name,
		Description:     utils.NilIfZero(r.Description, ""),
		Price:           r.Price,
		BillingInterval: types.BillingInterval(r.BillingInterval),
		Features:        r.Features.ToDomain(),
	}
}
