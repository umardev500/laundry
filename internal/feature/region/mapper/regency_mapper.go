package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/region/domain"
	"github.com/umardev500/laundry/internal/feature/region/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

func FromEntRegency(e *ent.Regency) *domain.Regency {
	return &domain.Regency{
		ID:         e.ID,
		ProvinceID: e.ProvinceID,
		Name:       e.Name,
	}
}

func FromEntRegencyList(list []*ent.Regency) []*domain.Regency {
	res := make([]*domain.Regency, len(list))
	for i, e := range list {
		res[i] = FromEntRegency(e)
	}
	return res
}

func ToResponseRegency(d *domain.Regency) *dto.RegencyResponse {
	if d == nil {
		return nil
	}

	return &dto.RegencyResponse{
		ID:         d.ID,
		ProvinceID: d.ProvinceID,
		Name:       d.Name,
	}
}

func ToResponseRegencyList(list []*domain.Regency) []*dto.RegencyResponse {
	res := make([]*dto.RegencyResponse, len(list))
	for i, d := range list {
		res[i] = ToResponseRegency(d)
	}
	return res
}

func ToResponsePageRegency(data *pagination.PageData[domain.Regency]) *pagination.PageData[dto.RegencyResponse] {
	res := ToResponseRegencyList(data.Data)
	return &pagination.PageData[dto.RegencyResponse]{
		Data:  res,
		Total: data.Total,
	}
}
