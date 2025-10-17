package mapper

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/feature/region/domain"
	"github.com/umardev500/laundry/internal/feature/region/dto"
	"github.com/umardev500/laundry/pkg/pagination"
)

func FromEntVillage(e *ent.Village) *domain.Village {
	return &domain.Village{
		ID:         e.ID,
		DistrictID: e.DistrictID,
		Name:       e.Name,
	}
}

func FromEntVillageList(list []*ent.Village) []*domain.Village {
	res := make([]*domain.Village, len(list))
	for i, e := range list {
		res[i] = FromEntVillage(e)
	}
	return res
}

func ToResponseVillage(d *domain.Village) *dto.VillageResponse {
	return &dto.VillageResponse{
		ID:         d.ID,
		DistrictID: d.DistrictID,
		Name:       d.Name,
	}
}

func ToResponseVillageList(list []*domain.Village) []*dto.VillageResponse {
	res := make([]*dto.VillageResponse, len(list))
	for i, d := range list {
		res[i] = ToResponseVillage(d)
	}
	return res
}

func ToResponsePageVillage(data *pagination.PageData[domain.Village]) *pagination.PageData[dto.VillageResponse] {
	res := ToResponseVillageList(data.Data)
	return &pagination.PageData[dto.VillageResponse]{
		Data:  res,
		Total: data.Total,
	}
}
