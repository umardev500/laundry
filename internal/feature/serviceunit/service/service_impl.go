package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/serviceunit/contract"
	"github.com/umardev500/laundry/internal/feature/serviceunit/domain"
	"github.com/umardev500/laundry/internal/feature/serviceunit/query"
	"github.com/umardev500/laundry/internal/feature/serviceunit/repository"
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

// Create adds a new service unit.
func (s *serviceImpl) Create(ctx *appctx.Context, u *domain.ServiceUnit) (*domain.ServiceUnit, error) {
	existing, err := s.repo.FindByName(ctx, u.Name)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if existing != nil {
		return nil, domain.ErrServiceUnitAlreadyExists
	}

	return s.repo.Create(ctx, u)
}

// List retrieves paginated service units.
func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListServiceUnitQuery) (*pagination.PageData[domain.ServiceUnit], error) {
	return s.repo.List(ctx, q)
}

// GetByID returns a single service unit by ID.
func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.ServiceUnit, error) {
	return s.findExisting(ctx, id)
}

// Update modifies an existing service unit.
func (s *serviceImpl) Update(ctx *appctx.Context, u *domain.ServiceUnit) (*domain.ServiceUnit, error) {
	unit, err := s.findExisting(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	unit.Update(u.Name, u.Symbol)
	return s.repo.Update(ctx, unit)
}

// Delete removes a service unit record permanently.
func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	unit, err := s.findExisting(ctx, id)
	if err != nil {
		return err
	}

	unit.SoftDelete()
	_, err = s.repo.Update(ctx, unit)
	return err
}

func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	u, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, u.ID)
}

// -------------------------
// Helpers
// -------------------------

// findExisting fetches a service unit that must exist and belong to the current tenant.
func (s *serviceImpl) findExisting(ctx *appctx.Context, id uuid.UUID) (*domain.ServiceUnit, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrServiceUnitNotFound
		}
		return nil, err
	}

	if u.IsDeleted() {
		return nil, domain.ErrServiceUnitDeleted
	}

	if !u.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedAccess
	}

	return u, nil
}

// findAllowDeleted fetches a service unit that may already be deleted, but still checks tenant ownership.
func (s *serviceImpl) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.ServiceUnit, error) {
	u, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrServiceUnitNotFound
		}
		return nil, err
	}

	if !u.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedAccess
	}

	return u, nil
}
