package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/serviceunit/domain"
	"github.com/umardev500/laundry/internal/feature/serviceunit/dto"
)

func FromEnt(e *ent.ServiceUnit) *domain.ServiceUnit {
	if e == nil {
		return nil
	}
	return &domain.ServiceUnit{
		ID:        e.ID,
		TenantID:  e.TenantID,
		Name:      e.Name,
		Symbol:    e.Symbol,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func FromEntList(list []*ent.ServiceUnit) []*domain.ServiceUnit {
	res := make([]*domain.ServiceUnit, len(list))
	for i, e := range list {
		res[i] = FromEnt(e)
	}
	return res
}

func ToResponse(d *domain.ServiceUnit) *dto.ServiceUnitResponse {
	if d == nil {
		return nil
	}
	return &dto.ServiceUnitResponse{
		ID:        d.ID,
		TenantID:  d.TenantID,
		Name:      d.Name,
		Symbol:    d.Symbol,
		CreatedAt: d.CreatedAt,
		UpdatedAt: d.UpdatedAt,
	}
}

func ToResponseList(list []*domain.ServiceUnit) []*dto.ServiceUnitResponse {
	res := make([]*dto.ServiceUnitResponse, len(list))
	for i, d := range list {
		res[i] = ToResponse(d)
	}
	return res
}
