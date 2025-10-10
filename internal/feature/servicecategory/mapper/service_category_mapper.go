package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/servicecategory/domain"
	"github.com/umardev500/laundry/internal/feature/servicecategory/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

// FromEnt converts an Ent model to a domain model.
func FromEnt(e *ent.ServiceCategory) *domain.ServiceCategory {
	if e == nil {
		return nil
	}

	return &domain.ServiceCategory{
		ID:          e.ID,
		TenantID:    e.TenantID,
		Name:        e.Name,
		Description: e.Description,
		DeletedAt:   e.DeletedAt,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
	}
}

// FromEntList converts a list of Ent models to domain models.
func FromEntList(ents []*ent.ServiceCategory) []*domain.ServiceCategory {
	items := make([]*domain.ServiceCategory, len(ents))
	for i, e := range ents {
		items[i] = FromEnt(e)
	}
	return items
}

// ToResponse converts a domain model to a response DTO.
func ToResponse(d *domain.ServiceCategory) *dto.ServiceCategoryResponse {
	if d == nil {
		return nil
	}

	return &dto.ServiceCategoryResponse{
		ID:          d.ID,
		TenantID:    d.TenantID,
		Name:        d.Name,
		Description: d.Description,
		DeletedAt:   d.DeletedAt,
		CreatedAt:   d.CreatedAt,
		UpdatedAt:   d.UpdatedAt,
	}
}

// ToResponseList converts a list of domain models to response DTOs.
func ToResponseList(items []*domain.ServiceCategory) []*dto.ServiceCategoryResponse {
	res := make([]*dto.ServiceCategoryResponse, len(items))
	for i, d := range items {
		res[i] = ToResponse(d)
	}
	return res
}

// ToResponsePage converts a paginated domain result into a paginated response DTO.
func ToResponsePage(data *pagination.PageData[domain.ServiceCategory]) *pagination.PageData[dto.ServiceCategoryResponse] {
	if data == nil {
		return nil
	}

	return &pagination.PageData[dto.ServiceCategoryResponse]{
		Data:  ToResponseList(data.Data),
		Total: data.Total,
	}
}
