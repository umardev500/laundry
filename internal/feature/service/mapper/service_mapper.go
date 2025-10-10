package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/service/domain"
	"github.com/umardev500/laundry/internal/feature/service/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

func FromEnt(e *ent.Service) *domain.Service {
	if e == nil {
		return nil
	}
	return &domain.Service{
		ID:                e.ID,
		TenantID:          e.TenantID,
		ServiceUnitID:     e.ServiceUnitID,
		ServiceCategoryID: e.ServiceCategoryID,
		Name:              e.Name,
		BasePrice:         e.BasePrice,
		Description:       e.Description,
		CreatedAt:         e.CreatedAt,
		UpdatedAt:         e.UpdatedAt,
		DeletedAt:         e.DeletedAt,
	}
}

func FromEntList(list []*ent.Service) []*domain.Service {
	res := make([]*domain.Service, len(list))
	for i, e := range list {
		res[i] = FromEnt(e)
	}
	return res
}

func ToResponse(d *domain.Service) *dto.ServiceResponse {
	if d == nil {
		return nil
	}
	return &dto.ServiceResponse{
		ID:                d.ID,
		TenantID:          d.TenantID,
		ServiceUnitID:     d.ServiceUnitID,
		ServiceCategoryID: d.ServiceCategoryID,
		Name:              d.Name,
		Price:             d.BasePrice,
		Description:       d.Description,
		CreatedAt:         d.CreatedAt,
		UpdatedAt:         d.UpdatedAt,
	}
}

func ToResponsePage(data *pagination.PageData[domain.Service]) *pagination.PageData[dto.ServiceResponse] {
	res := make([]*dto.ServiceResponse, len(data.Data))
	for i, m := range data.Data {
		res[i] = ToResponse(m)
	}
	return &pagination.PageData[dto.ServiceResponse]{
		Data:  res,
		Total: data.Total,
	}
}
