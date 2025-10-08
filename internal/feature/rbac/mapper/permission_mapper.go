package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

func FromEntPermission(e *ent.Permission) *domain.Permission {
	return &domain.Permission{
		ID:          e.ID,
		Name:        e.Name,
		DisplayName: e.DisplayName,
		Description: e.Description,
		Status:      types.Status(e.Status),
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   e.DeletedAt,
	}
}

func FromEntPermissionList(list []*ent.Permission) []*domain.Permission {
	res := make([]*domain.Permission, len(list))
	for i, e := range list {
		res[i] = FromEntPermission(e)
	}
	return res
}

func ToPermissionResponse(p *domain.Permission) *dto.PermissionResponse {
	if p == nil {
		return nil
	}
	return &dto.PermissionResponse{
		ID:          p.ID,
		Name:        p.Name,
		DisplayName: p.DisplayName,
		Description: p.Description,
		Status:      p.Status,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   p.DeletedAt,
	}
}

func ToPermissionResponsePage(page *pagination.PageData[domain.Permission]) *pagination.PageData[dto.PermissionResponse] {
	data := make([]*dto.PermissionResponse, len(page.Data))
	for i, p := range page.Data {
		data[i] = ToPermissionResponse(p)
	}
	return &pagination.PageData[dto.PermissionResponse]{Data: data, Total: page.Total}
}
