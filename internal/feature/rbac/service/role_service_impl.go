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

type serviceImpl struct {
	repo repository.RoleRepository
}

func NewService(repo repository.RoleRepository) contract.Service {
	return &serviceImpl{repo: repo}
}

// Create adds a new role after checking for duplicates.
func (s *serviceImpl) Create(ctx *appctx.Context, r *domain.Role) (*domain.Role, error) {
	existing, err := s.repo.FindByName(ctx, r.Name)
	if err != nil && !ent.IsNotFound(err) {
		return nil, err
	}

	if existing != nil {
		return nil, domain.ErrRoleAlreadyExists
	}

	return s.repo.Create(ctx, r)
}

// List returns a paginated list of roles.
func (s *serviceImpl) List(ctx *appctx.Context, q *query.ListRoleQuery) (*pagination.PageData[domain.Role], error) {
	return s.repo.List(ctx, q)
}

// GetByID retrieves a role by its ID.
func (s *serviceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Role, error) {
	return s.findExistingRole(ctx, id)
}

// GetByName retrieves a role by its name and tenant ID.
func (s *serviceImpl) GetByName(ctx *appctx.Context, name string) (*domain.Role, error) {
	role, err := s.repo.FindByName(ctx, name)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrRoleNotFound
		}
		return nil, err
	}

	return role, nil
}

// Update changes the name or description of a role.
func (s *serviceImpl) Update(ctx *appctx.Context, r *domain.Role) (*domain.Role, error) {
	role, err := s.findExistingRole(ctx, r.ID)
	if err != nil {
		return nil, err
	}

	role.Update(r.Name, r.Description)
	return s.repo.Update(ctx, role)
}

// Delete performs a soft-delete on a role.
func (s *serviceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	role, err := s.findExistingRole(ctx, id)
	if err != nil {
		return err
	}

	role.SoftDelete()
	_, err = s.repo.Update(ctx, role)
	return err
}

// Purge permanently deletes a role (bypass soft-delete).
func (s *serviceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	role, err := s.findRoleAllowDeleted(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, role.ID)
}

// findExistingRole ensures the role exists and is not soft-deleted.
func (s *serviceImpl) findExistingRole(ctx *appctx.Context, id uuid.UUID) (*domain.Role, error) {
	role, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrRoleNotFound
		}
		return nil, err
	}

	if role.IsDeleted() {
		return nil, domain.ErrRoleDeleted
	}

	if !role.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedRoleAccess
	}

	return role, nil
}

// findRoleAllowDeleted retrieves a role even if it’s deleted (for purge).
func (s *serviceImpl) findRoleAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.Role, error) {
	role, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrRoleNotFound
		}
		return nil, err
	}

	if !role.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedRoleAccess
	}

	return role, nil
}
