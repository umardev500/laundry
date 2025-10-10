package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/servicecategory/contract"
	"github.com/umardev500/laundry/internal/feature/servicecategory/query"
	"github.com/umardev500/laundry/internal/feature/servicecategory/repository"
	"github.com/umardev500/laundry/pkg/pagination"

	domain "github.com/umardev500/laundry/internal/feature/servicecategory/domain"
)

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) contract.Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) Create(ctx *appctx.Context, c *domain.ServiceCategory) (*domain.ServiceCategory, error) {
	existing, err := s.repo.FindByName(ctx, c.Name)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if existing != nil {
		return nil, domain.ErrServiceCategoryAlreadyExists
	}

	return s.repo.Create(ctx, c)
}

func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListServiceCategoryQuery) (*pagination.PageData[domain.ServiceCategory], error) {
	return s.repo.List(ctx, q)
}

func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.ServiceCategory, error) {
	return s.findExisting(ctx, id)
}

func (s *serviceImpl) Update(ctx *appctx.Context, c *domain.ServiceCategory) (*domain.ServiceCategory, error) {
	category, err := s.findExisting(ctx, c.ID)
	if err != nil {
		return nil, err
	}

	category.Update(c.Name, c.Description)
	return s.repo.Update(ctx, category)
}

func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	category, err := s.findExisting(ctx, id)
	if err != nil {
		return err
	}

	category.SoftDelete()
	_, err = s.repo.Update(ctx, category)
	return err
}

func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	c, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, c.ID)
}

// --- helpers ---

func (s *serviceImpl) findExisting(ctx *appctx.Context, id uuid.UUID) (*domain.ServiceCategory, error) {
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrServiceCategoryNotFound
		}
		return nil, err
	}

	if c.IsDeleted() {
		return nil, domain.ErrServiceCategoryDeleted
	}

	if !c.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedAccess
	}

	return c, nil
}

func (s *serviceImpl) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.ServiceCategory, error) {
	c, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrServiceCategoryNotFound
		}
		return nil, err
	}

	if !c.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedAccess
	}

	return c, nil
}
