package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/machinetype/domain"
	"github.com/umardev500/laundry/internal/feature/machinetype/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

func FromEntModel(e *ent.MachineType) *domain.MachineType {
	if e == nil {
		return nil
	}
	return &domain.MachineType{
		ID:          e.ID,
		TenantID:    e.TenantID,
		Name:        e.Name,
		Description: e.Description,
		Capacity:    e.Capacity,
		CreatedAt:   e.CreatedAt,
		UpdatedAt:   e.UpdatedAt,
		DeletedAt:   e.DeletedAt,
	}
}

func FromEntModels(list []*ent.MachineType) []*domain.MachineType {
	res := make([]*domain.MachineType, len(list))
	for i, e := range list {
		res[i] = FromEntModel(e)
	}
	return res
}

func ToResponse(d *domain.MachineType) *dto.MachineTypeResponse {
	if d == nil {
		return nil
	}
	return &dto.MachineTypeResponse{
		ID:          d.ID,
		TenantID:    d.TenantID,
		Name:        d.Name,
		Description: d.Description,
		Capacity:    d.Capacity,
	}
}

func ToResponsePage(data *pagination.PageData[domain.MachineType]) *pagination.PageData[dto.MachineTypeResponse] {
	res := make([]*dto.MachineTypeResponse, len(data.Data))
	for i, m := range data.Data {
		res[i] = ToResponse(m)
	}
	return &pagination.PageData[dto.MachineTypeResponse]{Data: res, Total: data.Total}
}
