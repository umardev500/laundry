package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/user/contract"
	"github.com/umardev500/laundry/internal/feature/user/domain"
	"github.com/umardev500/laundry/internal/feature/user/query"
	"github.com/umardev500/laundry/internal/feature/user/repository"
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

// Create implements contract.Service.
func (s *serviceImpl) Create(ctx *appctx.Context, user *domain.User) (*domain.User, error) {
	user, err := s.repo.FindByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if user != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	return s.repo.Create(ctx, user)
}

// Delete implements contract.Service.
func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	user, err := s.findExistingUser(ctx, id)
	if err != nil {
		return err
	}

	user.SoftDelete()

	_, err = s.repo.Update(ctx, user)
	return err
}

// GetByEmail implements contract.Service.
func (s *serviceImpl) GetByEmail(ctx *appctx.Context, email string) (*domain.User, error) {
	return s.repo.FindByEmail(ctx, email)
}

// GetByID implements contract.Service.
func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.User, error) {
	return s.repo.FindById(ctx, id)
}

// List implements contract.Service.
func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListUserQuery) (*pagination.PageData[domain.User], error) {
	return s.repo.List(ctx, q)
}

// Update implements contract.Service.
func (s *serviceImpl) Update(ctx *appctx.Context, u *domain.User) (*domain.User, error) {
	user, err := s.findExistingUser(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	user.Update(u.Email, u.Password)

	return s.repo.Update(ctx, user)
}

func (s *serviceImpl) UpdateStatus(ctx *appctx.Context, u *domain.User) (*domain.User, error) {
	user, err := s.findExistingUser(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	err = user.SetStatus(u.Status)
	if err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, user)
}

// Purge implements contract.Service.
func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	user, err := s.findExistingUser(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, user.ID)
}

// findExistingUser returns any user that is not deleted.
func (s *serviceImpl) findExistingUser(ctx *appctx.Context, id uuid.UUID) (*domain.User, error) {
	user, err := s.repo.FindById(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrUserNotFound
		}
		return nil, err
	}

	if user.IsDeleted() {
		return nil, domain.ErrUserDeleted
	}

	return user, nil
}
