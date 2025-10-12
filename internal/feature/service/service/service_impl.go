package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/service/contract"
	"github.com/umardev500/laundry/internal/feature/service/domain"
	"github.com/umardev500/laundry/internal/feature/service/query"
	"github.com/umardev500/laundry/internal/feature/service/repository"
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

// AreItemsAvailable returns a map of service ids to availability.
func (s *serviceImpl) AreItemsAvailable(ctx *appctx.Context, ids []uuid.UUID) (*domain.AvailabilityResult, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	return s.repo.AreItemsAvailable(ctx, ids)
}

// Create adds a new service.
func (s *serviceImpl) Create(ctx *appctx.Context, svc *domain.Service) (*domain.Service, error) {
	existing, err := s.repo.FindByName(ctx, svc.Name)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrServiceAlreadyExists
	}
	return s.repo.Create(ctx, svc)
}

// List returns paginated services.
func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListServiceQuery) (*pagination.PageData[domain.Service], error) {
	return s.repo.List(ctx, q)
}

// GetByID returns service by id (must exist and not deleted).
func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Service, error) {
	return s.findExisting(ctx, id)
}

// GetByName returns service by name.
func (s *serviceImpl) GetByName(ctx *appctx.Context, name string) (*domain.Service, error) {
	return s.repo.FindByName(ctx, name)
}

// Update modifies a service.
func (s *serviceImpl) Update(ctx *appctx.Context, svc *domain.Service) (*domain.Service, error) {
	existing, err := s.findExisting(ctx, svc.ID)
	if err != nil {
		return nil, err
	}

	// Update fields; price only if non-negative sentinel.
	existing.Update(svc.Name, svc.BasePrice, svc.Description, svc.ServiceUnitID, svc.ServiceCategoryID)
	return s.repo.Update(ctx, existing)
}

// Delete performs soft-delete (marks deleted_at and updates).
func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	svc, err := s.findExisting(ctx, id)
	if err != nil {
		return err
	}

	svc.SoftDelete()
	_, err = s.repo.Update(ctx, svc)
	return err
}

// Purge permanently deletes a service (allowed even if soft-deleted).
func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	svc, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, svc.ID)
}

// findExisting ensures service exists, is not soft-deleted and belongs to tenant.
func (s *serviceImpl) findExisting(ctx *appctx.Context, id uuid.UUID) (*domain.Service, error) {
	svc, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrServiceNotFound
		}
		return nil, err
	}

	if svc.IsDeleted() {
		return nil, domain.ErrServiceDeleted
	}

	if !svc.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedServiceAccess
	}

	return svc, nil
}

// findAllowDeleted fetches service regardless of deleted status but checks tenant ownership.
func (s *serviceImpl) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.Service, error) {
	svc, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrServiceNotFound
		}
		return nil, err
	}

	if !svc.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedServiceAccess
	}

	return svc, nil
}
