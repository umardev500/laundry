package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/tenant/contract"
	"github.com/umardev500/laundry/internal/feature/tenant/domain"
	"github.com/umardev500/laundry/internal/feature/tenant/query"
	"github.com/umardev500/laundry/internal/feature/tenant/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) contract.Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) Create(ctx *appctx.Context, t *domain.Tenant) (*domain.Tenant, error) {
	tenant, err := s.repo.FindByEmail(ctx, t.Email)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if tenant != nil {
		return nil, domain.ErrTenantAlreadyExists
	}

	return s.repo.Create(ctx, t)
}

func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListTenantQuery) (*pagination.PageData[domain.Tenant], error) {
	return s.repo.List(ctx, q)
}

func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Tenant, error) {
	return s.findExistingTenant(ctx, id)
}

func (s *serviceImpl) GetByEmail(ctx *appctx.Context, email string) (*domain.Tenant, error) {
	return s.repo.FindByEmail(ctx, email)
}

func (s *serviceImpl) Update(ctx *appctx.Context, t *domain.Tenant) (*domain.Tenant, error) {
	tenant, err := s.findExistingTenant(ctx, t.ID)
	if err != nil {
		return nil, err
	}

	tenant.Update(t.Name, t.Email)
	return s.repo.Update(ctx, tenant)
}

func (s *serviceImpl) UpdateStatus(ctx *appctx.Context, t *domain.Tenant) (*domain.Tenant, error) {
	tenant, err := s.findExistingTenant(ctx, t.ID)
	if err != nil {
		return nil, err
	}

	err = tenant.SetStatus(t.Status)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, tenant)
}

func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	tenant, err := s.findExistingTenant(ctx, id)
	if err != nil {
		return err
	}

	tenant.SoftDelete()

	_, err = s.repo.Update(ctx, tenant)
	return err
}

func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	tenant, err := s.findTenantAllowDeleted(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, tenant.ID)
}

func (s *serviceImpl) findExistingTenant(ctx *appctx.Context, id uuid.UUID) (*domain.Tenant, error) {
	tenant, err := s.repo.FindById(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrTenantNotFound
		}
		return nil, err
	}

	if tenant.IsDeleted() {
		return nil, domain.ErrTenantDeleted
	}

	return tenant, nil
}

func (s *serviceImpl) findTenantAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.Tenant, error) {
	tenant, err := s.repo.FindById(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrTenantNotFound
		}
		return nil, err
	}
	return tenant, nil // ✅ don’t block deleted tenants
}
