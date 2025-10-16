package mapper

import (
	"encoding/json"

	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/plan/domain"
	"github.com/umardev500/laundry/internal/feature/plan/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

// FromEnt converts an Ent Plan model to a domain Plan.
func FromEnt(e *ent.Plan) *domain.Plan {
	if e == nil {
		return nil
	}

	// Map Ent features (map[string]interface{}) to PlanFeatures
	var features *domain.PlanFeatures
	if e.Features != nil {
		// Convert map[string]interface{} to JSON bytes
		data, err := json.Marshal(e.Features)
		if err == nil {
			// Unmarshal into PlanFeatures struct
			features = &domain.PlanFeatures{}
			_ = json.Unmarshal(data, features) // ignore error for simplicity; handle in prod
		}
	}

	return &domain.Plan{
		ID:              e.ID,
		Name:            e.Name,
		Description:     e.Description,
		Price:           e.Price,
		BillingInterval: types.BillingInterval(e.BillingInterval),
		Features:        features,
		Active:          e.Active,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
		DeletedAt:       e.DeletedAt,
	}
}

// FromEntList converts a slice of Ent Plan models to a slice of domain Plans.
func FromEntList(list []*ent.Plan) []*domain.Plan {
	res := make([]*domain.Plan, len(list))
	for i, e := range list {
		res[i] = FromEnt(e)
	}
	return res
}

// ToResponse converts a domain Plan to a PlanResponse DTO.
func ToResponse(d *domain.Plan) *dto.PlanResponse {
	if d == nil {
		return nil
	}
	return &dto.PlanResponse{
		ID:              d.ID,
		Name:            d.Name,
		Description:     d.Description,
		Price:           d.Price,
		BillingInterval: string(d.BillingInterval),
		Features:        d.Features,
		Active:          d.Active,
		CreatedAt:       d.CreatedAt,
		UpdatedAt:       d.UpdatedAt,
	}
}

func ToResponseList(list []*domain.Plan) []*dto.PlanResponse {
	res := make([]*dto.PlanResponse, len(list))
	for i, d := range list {
		res[i] = ToResponse(d)
	}
	return res
}

// ToResponsePage converts a paginated domain Plan slice to a paginated DTO slice.
func ToResponsePage(data *pagination.PageData[domain.Plan]) *pagination.PageData[dto.PlanResponse] {
	res := make([]*dto.PlanResponse, len(data.Data))
	for i, m := range data.Data {
		res[i] = ToResponse(m)
	}
	return &pagination.PageData[dto.PlanResponse]{
		Data:  res,
		Total: data.Total,
	}
}
