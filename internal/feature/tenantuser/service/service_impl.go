package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/tenantuser/contract"
	"github.com/umardev500/laundry/internal/feature/tenantuser/domain"
	"github.com/umardev500/laundry/internal/feature/tenantuser/query"
	"github.com/umardev500/laundry/internal/feature/tenantuser/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) contract.Service {
	return &serviceImpl{repo: repo}
}

// Create creates a tenant-user mapping if it doesn't already exist.
func (s *serviceImpl) Create(ctx *appctx.Context, tu *domain.TenantUser) (*domain.TenantUser, error) {
	// Check existing mapping by user+tenant
	existing, err := s.repo.FindByUserAndTenant(ctx, tu.UserID, tu.TenantID)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrTenantUserAlreadyExists
	}

	return s.repo.Create(ctx, tu)
}

// GetByUser implements contract.Service.
func (s *serviceImpl) GetByUser(ctx *appctx.Context, userID uuid.UUID) ([]*domain.TenantUser, error) {
	return s.repo.FindByUser(ctx, userID)
}

func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.TenantUser, error) {
	return s.findExistingTenantUser(ctx, id)
}

func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListTenantUserQuery) (*pagination.PageData[domain.TenantUser], error) {
	return s.repo.List(ctx, q)
}

func (s *serviceImpl) UpdateStatus(ctx *appctx.Context, tu *domain.TenantUser) (*domain.TenantUser, error) {
	existing, err := s.findExistingTenantUser(ctx, tu.ID)
	if err != nil {
		return nil, err
	}

	if err := existing.SetStatus(tu.Status); err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, existing)
}

func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	tu, err := s.findExistingTenantUser(ctx, id)
	if err != nil {
		return err
	}

	tu.SoftDelete()
	_, err = s.repo.Update(ctx, tu)
	return err
}

func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	tu, err := s.findTenantUserAllowDeleted(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, tu.ID)
}

func (s *serviceImpl) findExistingTenantUser(ctx *appctx.Context, id uuid.UUID) (*domain.TenantUser, error) {
	tu, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrTenantUserNotFound
		}
		return nil, err
	}

	if tu.IsDeleted() {
		return nil, domain.ErrTenantUserDeleted
	}

	return tu, nil
}

func (s *serviceImpl) findTenantUserAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.TenantUser, error) {
	tu, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrTenantUserNotFound
		}
		return nil, err
	}
	return tu, nil
}
