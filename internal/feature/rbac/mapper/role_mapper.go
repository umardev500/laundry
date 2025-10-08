package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

// FromEntModel converts an ent.Role to a domain.Role
func FromEntModel(e *ent.Role) *domain.Role {
	if e == nil {
		return nil
	}

	return &domain.Role{
		ID:          e.ID,
		TenantID:    e.TenantID,
		Name:        e.Name,
		Description: e.Description,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   e.DeletedAt,
	}
}

// FromEntModels converts []*ent.Role to []*domain.Role
func FromEntModels(entities []*ent.Role) []*domain.Role {
	result := make([]*domain.Role, len(entities))
	for i, e := range entities {
		result[i] = FromEntModel(e)
	}
	return result
}

func ToRoleResponse(r *domain.Role) *dto.RoleResponse {
	return &dto.RoleResponse{
		ID:          r.ID.String(),
		TenantID:    r.TenantID,
		Name:        r.Name,
		Description: r.Description,
	}
}

// ToRoleResponsePage converts pagination.PageData[domain.Role] to pagination.PageData[dto.RoleResponse]
func ToRoleResponsePage(data *pagination.PageData[domain.Role]) *pagination.PageData[dto.RoleResponse] {
	res := make([]*dto.RoleResponse, len(data.Data))
	for i, r := range data.Data {
		res[i] = ToRoleResponse(r)
	}

	return &pagination.PageData[dto.RoleResponse]{
		Data:  res,
		Total: data.Total,
	}
}
