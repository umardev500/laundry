package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/region/domain"
	"github.com/umardev500/laundry/internal/feature/region/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

func FromEntDistrict(e *ent.District) *domain.District {
	return &domain.District{
		ID:        e.ID,
		RegencyID: e.RegencyID,
		Name:      e.Name,
	}
}

func FromEntDistrictList(list []*ent.District) []*domain.District {
	res := make([]*domain.District, len(list))
	for i, e := range list {
		res[i] = FromEntDistrict(e)
	}
	return res
}

func ToResponseDistrict(d *domain.District) *dto.DistrictResponse {
	return &dto.DistrictResponse{
		ID:        d.ID,
		RegencyID: d.RegencyID,
		Name:      d.Name,
	}
}

func ToResponseDistrictList(list []*domain.District) []*dto.DistrictResponse {
	res := make([]*dto.DistrictResponse, len(list))
	for i, d := range list {
		res[i] = ToResponseDistrict(d)
	}
	return res
}

func ToResponsePageDistrict(data *pagination.PageData[domain.District]) *pagination.PageData[dto.DistrictResponse] {
	res := ToResponseDistrictList(data.Data)
	return &pagination.PageData[dto.DistrictResponse]{
		Data:  res,
		Total: data.Total,
	}
}
