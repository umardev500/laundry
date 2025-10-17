package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/region/domain"
	"github.com/umardev500/laundry/internal/feature/region/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

func FromEntProvice(e *ent.Province) *domain.Province {
	return &domain.Province{
		ID:   e.ID,
		Name: e.Name,
	}
}

func FromEntProvinceList(list []*ent.Province) []*domain.Province {
	res := make([]*domain.Province, len(list))
	for i, e := range list {
		res[i] = FromEntProvice(e)
	}
	return res
}

func ToResponseProvince(d *domain.Province) *dto.ProvinceResponse {
	return &dto.ProvinceResponse{
		ID:   d.ID,
		Name: d.Name,
	}
}

func ToResponseProvinceList(list []*domain.Province) []*dto.ProvinceResponse {
	res := make([]*dto.ProvinceResponse, len(list))
	for i, d := range list {
		res[i] = ToResponseProvince(d)
	}
	return res
}

func ToResponsePageProvince(data *pagination.PageData[domain.Province]) *pagination.PageData[dto.ProvinceResponse] {
	res := ToResponseProvinceList(data.Data)
	return &pagination.PageData[dto.ProvinceResponse]{
		Data:  res,
		Total: data.Total,
	}
}
