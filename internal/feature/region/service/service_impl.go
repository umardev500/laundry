package service

import (
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/region/contract"
	"github.com/umardev500/laundry/internal/feature/region/domain"
	"github.com/umardev500/laundry/internal/feature/region/query"
	"github.com/umardev500/laundry/internal/feature/region/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) contract.Service {
	return &serviceImpl{
		repo: repo,
	}
}

// FindVillages implements contract.Service.
func (s *serviceImpl) FindVillages(ctx *appctx.Context, districtID string, q *query.ListVillageQuery) (*pagination.PageData[domain.Village], error) {
	return s.repo.FindVillages(ctx, districtID, q)
}

// FindDistricts implements contract.Service.
func (s *serviceImpl) FindDistricts(ctx *appctx.Context, regencyID string, q *query.ListDistrictQuery) (*pagination.PageData[domain.District], error) {
	return s.repo.FindDistricts(ctx, regencyID, q)
}

// FindRegencies implements contract.Service.
func (s *serviceImpl) FindRegencies(ctx *appctx.Context, provinceID string, q *query.ListRegencyQuery) (*pagination.PageData[domain.Regency], error) {
	return s.repo.FindRegencies(ctx, provinceID, q)
}

// FindProvinces implements contract.Service.
func (s *serviceImpl) FindProvinces(ctx *appctx.Context, q *query.ListProvinceQuery) (*pagination.PageData[domain.Province], error) {
	return s.repo.FindProvinces(ctx, q)
}
