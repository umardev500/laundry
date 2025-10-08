package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/platformuser/contract"
	"github.com/umardev500/laundry/internal/feature/platformuser/domain"
	"github.com/umardev500/laundry/internal/feature/platformuser/query"
	"github.com/umardev500/laundry/internal/feature/platformuser/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type serviceImpl struct {
	repo repository.Repository
}

// NewService creates a new PlatformUser service
func NewService(repo repository.Repository) contract.Service {
	return &serviceImpl{
		repo: repo,
	}
}

// Create creates a new PlatformUser for an existing user
func (s *serviceImpl) Create(ctx *appctx.Context, pu *domain.PlatformUser) (*domain.PlatformUser, error) {
	// Ensure PlatformUser does not already exist for this user
	existing, err := s.repo.FindByIdUserID(ctx, pu.UserID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrPlatformUserAlreadyExists
	}

	return s.repo.Create(ctx, pu)
}

// GetByID retrieves a PlatformUser by its ID
func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.PlatformUser, error) {
	return s.findExistingPlatformUser(ctx, id)
}

// GetByUserID retrieves a PlatformUser by the linked UserID
func (s *serviceImpl) GetByUserID(ctx *appctx.Context, userID uuid.UUID) (*domain.PlatformUser, error) {
	pu, err := s.repo.FindByIdUserID(ctx, userID)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrPlatformUserNotFound
		}
		return nil, err
	}

	if pu.IsDeleted() {
		return nil, domain.ErrPlatformUserDeleted
	}

	return pu, nil
}

// List returns paginated PlatformUsers
func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListPlatformUserQuery) (*pagination.PageData[domain.PlatformUser], error) {
	return s.repo.List(ctx, q)
}

// UpdateStatus updates only the status of a PlatformUser
func (s *serviceImpl) UpdateStatus(ctx *appctx.Context, pu *domain.PlatformUser) (*domain.PlatformUser, error) {
	existing, err := s.findExistingPlatformUser(ctx, pu.ID)
	if err != nil {
		return nil, err
	}

	err = existing.SetStatus(pu.Status)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, existing)
}

// Delete soft deletes a PlatformUser
func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	pu, err := s.findExistingPlatformUser(ctx, id)
	if err != nil {
		return err
	}

	pu.SoftDelete()
	_, err = s.repo.Update(ctx, pu)
	return err
}

// Purge permanently deletes a PlatformUser
func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	_, err := s.findExistingPlatformUser(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}

// findExistingPlatformUser ensures the PlatformUser exists and is not deleted
func (s *serviceImpl) findExistingPlatformUser(ctx *appctx.Context, id uuid.UUID) (*domain.PlatformUser, error) {
	pu, err := s.repo.FindById(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrPlatformUserNotFound
		}
		return nil, err
	}

	if pu.IsDeleted() {
		return nil, domain.ErrPlatformUserDeleted
	}

	return pu, nil
}
