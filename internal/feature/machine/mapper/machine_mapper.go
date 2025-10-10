package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/machine/domain"
	"github.com/umardev500/laundry/internal/feature/machine/dto"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

func FromEntModel(e *ent.Machine) *domain.Machine {
	if e == nil {
		return nil
	}
	return &domain.Machine{
		ID:            e.ID,
		TenantID:      e.TenantID,
		MachineTypeID: e.MachineTypeID,
		Name:          e.Name,
		Description:   e.Description,
		Status:        types.MachineStatus(e.Status),
		CreatedAt:     e.CreatedAt,
		UpdatedAt:     e.UpdatedAt,
		DeletedAt:     e.DeletedAt,
	}
}

func FromEntModels(entities []*ent.Machine) []*domain.Machine {
	res := make([]*domain.Machine, len(entities))
	for i, e := range entities {
		res[i] = FromEntModel(e)
	}
	return res
}

func ToMachineResponse(d *domain.Machine) *dto.MachineResponse {
	if d == nil {
		return nil
	}
	return &dto.MachineResponse{
		ID:            d.ID,
		TenantID:      d.TenantID,
		MachineTypeID: d.MachineTypeID,
		Name:          d.Name,
		Description:   d.Description,
		Status:        d.Status,
	}
}

func ToMachineResponsePage(data *pagination.PageData[domain.Machine]) *pagination.PageData[dto.MachineResponse] {
	res := make([]*dto.MachineResponse, len(data.Data))
	for i, m := range data.Data {
		res[i] = ToMachineResponse(m)
	}
	return &pagination.PageData[dto.MachineResponse]{
		Data:  res,
		Total: data.Total,
	}
}
