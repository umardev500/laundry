package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/rbac/contract"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/query"
	"github.com/umardev500/laundry/internal/feature/rbac/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type featureServiceImpl struct {
	repo repository.FeatureRepository
}

func NewFeatureService(repo repository.FeatureRepository) contract.FeatureService {
	return &featureServiceImpl{repo: repo}
}

func (s *featureServiceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Feature, error) {
	feature, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrFeatureNotFound
		}
		return nil, err
	}

	if feature.IsDeleted() {
		return nil, domain.ErrFeatureDeleted
	}

	return feature, nil
}

func (s *featureServiceImpl) List(ctx *appctx.Context, q *query.ListFeatureQuery) (*pagination.PageData[domain.Feature], error) {
	return s.repo.List(ctx, q)
}

func (s *featureServiceImpl) Update(ctx *appctx.Context, f *domain.Feature) (*domain.Feature, error) {
	existing, err := s.repo.FindByID(ctx, f.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrFeatureNotFound
		}
		return nil, err
	}

	existing.Update(f.Name, f.Description)

	return s.repo.Update(ctx, existing)
}

func (s *featureServiceImpl) UpdateStatus(ctx *appctx.Context, f *domain.Feature) (*domain.Feature, error) {
	existing, err := s.repo.FindByID(ctx, f.ID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrFeatureNotFound
		}
		return nil, err
	}

	err = existing.SetStatus(f.Status)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, existing)
}
