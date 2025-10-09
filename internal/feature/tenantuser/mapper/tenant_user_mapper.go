package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/tenantuser/domain"
	"github.com/umardev500/laundry/internal/feature/tenantuser/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

// FromEntModel converts an ent.TenantUser to domain.TenantUser
func FromEntModel(e *ent.TenantUser) *domain.TenantUser {
	if e == nil {
		return nil
	}

	return &domain.TenantUser{
		ID:        e.ID,
		UserID:    e.UserID,
		TenantID:  e.TenantID,
		Status:    types.Status(*e.Status),
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		DeletedAt: e.DeletedAt,
	}
}

// FromEntModels converts a slice of ent.TenantUser to []*domain.TenantUser
func FromEntModels(entities []*ent.TenantUser) []*domain.TenantUser {
	if entities == nil {
		return nil
	}
	result := make([]*domain.TenantUser, len(entities))
	for i, e := range entities {
		result[i] = FromEntModel(e)
	}
	return result
}

// ToTenantUserResponse maps a domain.TenantUser to dto.TenantUserResponse
func ToTenantUserResponse(d *domain.TenantUser) *dto.TenantUserResponse {
	if d == nil {
		return nil
	}
	return &dto.TenantUserResponse{
		ID:        d.ID,
		UserID:    d.UserID,
		TenantID:  d.TenantID,
		Status:    d.Status,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
		DeletedAt: d.DeletedAt,
	}
}

// ToTenantUserResponsePage converts pagination.PageData[domain.TenantUser] to pagination.PageData[dto.TenantUserResponse]
func ToTenantUserResponsePage(data *pagination.PageData[domain.TenantUser]) *pagination.PageData[dto.TenantUserResponse] {
	res := make([]*dto.TenantUserResponse, len(data.Data))
	for i, t := range data.Data {
		res[i] = ToTenantUserResponse(t)
	}

	return &pagination.PageData[dto.TenantUserResponse]{
		Data:  res,
		Total: data.Total,
	}
}
