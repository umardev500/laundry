package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/machine/contract"
	"github.com/umardev500/laundry/internal/feature/machine/domain"
	"github.com/umardev500/laundry/internal/feature/machine/query"
	"github.com/umardev500/laundry/internal/feature/machine/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

type serviceImpl struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) contract.Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) Create(ctx *appctx.Context, m *domain.Machine) (*domain.Machine, error) {
	existing, err := s.repo.FindByName(ctx, m.Name)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if existing != nil {
		return nil, domain.ErrMachineAlreadyExists
	}

	return s.repo.Create(ctx, m)
}

func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListMachineQuery) (*pagination.PageData[domain.Machine], error) {
	return s.repo.List(ctx, q)
}

func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Machine, error) {
	return s.findExisting(ctx, id)
}

func (s *serviceImpl) GetByName(ctx *appctx.Context, name string) (*domain.Machine, error) {
	return s.repo.FindByName(ctx, name)
}

func (s *serviceImpl) Update(ctx *appctx.Context, m *domain.Machine) (*domain.Machine, error) {
	machine, err := s.findExisting(ctx, m.ID)
	if err != nil {
		return nil, err
	}

	machine.Update(m.Name, m.Description)
	return s.repo.Update(ctx, machine)
}

func (s *serviceImpl) UpdateStatus(ctx *appctx.Context, m *domain.Machine) (*domain.Machine, error) {
	machine, err := s.findExisting(ctx, m.ID)
	if err != nil {
		return nil, err
	}

	err = machine.SetStatus(m.Status)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, machine)
}

func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	machine, err := s.findExisting(ctx, id)
	if err != nil {
		return err
	}

	machine.SoftDelete()
	_, err = s.repo.Update(ctx, machine)
	return err
}

func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	m, err := s.findAllowDeleted(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, m.ID)
}

func (s *serviceImpl) findExisting(ctx *appctx.Context, id uuid.UUID) (*domain.Machine, error) {
	m, err := s.repo.FindById(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrMachineNotFound
		}
		return nil, err
	}

	if m.IsDeleted() {
		return nil, domain.ErrMachineDeleted
	}

	if !m.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedMachineAccess
	}

	return m, nil
}

func (s *serviceImpl) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.Machine, error) {
	m, err := s.repo.FindById(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrMachineNotFound
		}
		return nil, err
	}

	if !m.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedMachineAccess
	}

	return m, nil
}
