package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

func FromFeatureEntModel(e *ent.Feature) *domain.Feature {
	if e == nil {
		return nil
	}
	return &domain.Feature{
		ID:          e.ID,
		Name:        e.Name,
		Description: e.Description,
		Status:      types.Status(e.Status),
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   e.DeletedAt,
	}
}

func FromFeatureEntModels(entities []*ent.Feature) []*domain.Feature {
	result := make([]*domain.Feature, len(entities))
	for i, e := range entities {
		result[i] = FromFeatureEntModel(e)
	}
	return result
}

func ToFeatureResponse(d *domain.Feature) *dto.FeatureResponse {
	if d == nil {
		return nil
	}
	return &dto.FeatureResponse{
		ID:          d.ID,
		Name:        d.Name,
		Description: d.Description,
		Status:      d.Status,
	}
}

func ToFeatureResponsePage(data *pagination.PageData[domain.Feature]) *pagination.PageData[dto.FeatureResponse] {
	res := make([]*dto.FeatureResponse, len(data.Data))
	for i, f := range data.Data {
		res[i] = ToFeatureResponse(f)
	}
	return &pagination.PageData[dto.FeatureResponse]{Data: res, Total: data.Total}
}
