package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/machinetype/contract"
	"github.com/umardev500/laundry/internal/feature/machinetype/domain"
	"github.com/umardev500/laundry/internal/feature/machinetype/query"
	"github.com/umardev500/laundry/internal/feature/machinetype/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) contract.Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) Create(ctx *appctx.Context, t *domain.MachineType) (*domain.MachineType, error) {
	existing, err := s.repo.FindByName(ctx, t.Name)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}
	if existing != nil {
		return nil, domain.ErrMachineTypeAlreadyExists
	}
	return s.repo.Create(ctx, t)
}

func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListMachineTypeQuery) (*pagination.PageData[domain.MachineType], error) {
	return s.repo.List(ctx, q)
}

func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.MachineType, error) {
	return s.findExisting(ctx, id)
}

func (s *serviceImpl) GetByName(ctx *appctx.Context, name string) (*domain.MachineType, error) {
	return s.repo.FindByName(ctx, name)
}

func (s *serviceImpl) Update(ctx *appctx.Context, t *domain.MachineType) (*domain.MachineType, error) {
	mt, err := s.findExisting(ctx, t.ID)
	if err != nil {
		return nil, err
	}
	mt.Update(t.Name, t.Description, t.Capacity)
	return s.repo.Update(ctx, mt)
}

func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	mt, err := s.findExisting(ctx, id)
	if err != nil {
		return err
	}
	mt.SoftDelete()
	_, err = s.repo.Update(ctx, mt)
	return err
}

func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	mt, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, mt.ID)
}

func (s *serviceImpl) findExisting(ctx *appctx.Context, id uuid.UUID) (*domain.MachineType, error) {
	m, err := s.repo.FindById(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrMachineTypeNotFound
		}
		return nil, err
	}
	if m.IsDeleted() {
		return nil, domain.ErrMachineTypeDeleted
	}
	if !m.BelongsToTenant(ctx) {
		return nil, domain.ErrMachineTypeDeleted // reuse, but better to have unauthorized; keep consistent with machine pattern
	}
	return m, nil
}

func (s *serviceImpl) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.MachineType, error) {
	m, err := s.repo.FindById(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrMachineTypeNotFound
		}
		return nil, err
	}
	if !m.BelongsToTenant(ctx) {
		return nil, domain.ErrMachineTypeDeleted
	}
	return m, nil
}
