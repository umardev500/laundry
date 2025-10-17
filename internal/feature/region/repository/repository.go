package repository

import (
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/region/domain"
	"github.com/umardev500/laundry/internal/feature/region/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	FindProvinces(ctx *appctx.Context, q *query.ListProvinceQuery) (*pagination.PageData[domain.Province], error)
	FindRegencies(ctx *appctx.Context, provinceID string, q *query.ListRegencyQuery) (*pagination.PageData[domain.Regency], error)
	FindDistricts(ctx *appctx.Context, regencyID string, q *query.ListDistrictQuery) (*pagination.PageData[domain.District], error)
	FindVillages(ctx *appctx.Context, districtID string, q *query.ListVillageQuery) (*pagination.PageData[domain.Village], error)
}
