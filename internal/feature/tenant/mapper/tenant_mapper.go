package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/tenant/domain"
	"github.com/umardev500/laundry/internal/feature/tenant/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

func FromEntModel(e *ent.Tenant) *domain.Tenant {
	if e == nil {
		return nil
	}
	return &domain.Tenant{
		ID:        e.ID,
		Name:      e.Name,
		Email:     e.Email,
		Phone:     e.Phone,
		Status:    types.TenantStatus(*e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}
}

func FromEntModels(entities []*ent.Tenant) []*domain.Tenant {
	result := make([]*domain.Tenant, len(entities))
	for i, e := range entities {
		result[i] = FromEntModel(e)
	}
	return result
}

// 🧱 ToTenantResponse maps a domain.Tenant to dto.TenantResponse
func ToTenantResponse(d *domain.Tenant) *dto.TenantResponse {
	if d == nil {
		return nil
	}

	return &dto.TenantResponse{
		ID:     d.ID,
		Name:   d.Name,
		Email:  d.Email,
		Phone:  d.Phone,
		Status: string(d.Status),
	}
}

// 📄 ToTenantResponsePage converts pagination.PageData[domain.Tenant] to pagination.PageData[dto.TenantResponse]
func ToTenantResponsePage(data *pagination.PageData[domain.Tenant]) *pagination.PageData[dto.TenantResponse] {
	res := make([]*dto.TenantResponse, len(data.Data))
	for i, t := range data.Data {
		res[i] = ToTenantResponse(t)
	}

	return &pagination.PageData[dto.TenantResponse]{
		Data:  res,
		Total: data.Total,
	}
}
