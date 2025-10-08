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

// permissionServiceImpl implements contract.PermissionService.
type permissionServiceImpl struct {
	repo repository.PermissionRepository
}

// NewPermissionService creates a new Permission service.
func NewPermissionService(repo repository.PermissionRepository) contract.PermissionService {
	return &permissionServiceImpl{repo: repo}
}

// 🔍 GetByID retrieves a permission by ID.
func (s *permissionServiceImpl) GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Permission, error) {
	return s.findExistingPermission(ctx, id)
}

// 📄 List returns a paginated list of permissions.
func (s *permissionServiceImpl) List(ctx *appctx.Context, q *query.ListPermissionQuery) (*pagination.PageData[domain.Permission], error) {
	return s.repo.List(ctx, q)
}

// ⚙️ Update modifies an existing permission.
func (s *permissionServiceImpl) Update(ctx *appctx.Context, p *domain.Permission) (*domain.Permission, error) {
	existing, err := s.findExistingPermission(ctx, p.ID)
	if err != nil {
		return nil, err
	}

	existing.Update(p.Name, p.DisplayName, p.Description)
	return s.repo.Update(ctx, existing)
}

// 🚦 UpdateStatus changes the permission’s status (e.g. active/suspended).
func (s *permissionServiceImpl) UpdateStatus(ctx *appctx.Context, id uuid.UUID, p *domain.Permission) (*domain.Permission, error) {
	existing, err := s.findExistingPermission(ctx, id)
	if err != nil {
		return nil, err
	}

	if err := existing.SetStatus(p.Status); err != nil {
		return nil, err
	}

	return s.repo.Update(ctx, existing)
}

// 🗑️ Soft delete a permission.
func (s *permissionServiceImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	permission, err := s.findExistingPermission(ctx, id)
	if err != nil {
		return err
	}

	permission.SoftDelete()
	_, err = s.repo.Update(ctx, permission)
	return err
}

// 💣 Permanently delete (purge) a permission.
func (s *permissionServiceImpl) Purge(ctx *appctx.Context, id uuid.UUID) error {
	permission, err := s.findPermissionAllowDeleted(ctx, id)
	if err != nil {
		return err
	}
	return s.repo.Delete(ctx, permission.ID)
}

//
// 🔒 Internal helper methods
//

// findExistingPermission ensures the permission exists and is not soft-deleted.
func (s *permissionServiceImpl) findExistingPermission(ctx *appctx.Context, id uuid.UUID) (*domain.Permission, error) {
	permission, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrPermissionNotFound
		}
		return nil, err
	}

	if permission.IsDeleted() {
		return nil, domain.ErrPermissionDeleted
	}

	return permission, nil
}

// findPermissionAllowDeleted retrieves a permission even if it’s deleted (for purge).
func (s *permissionServiceImpl) findPermissionAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.Permission, error) {
	permission, err := s.repo.FindByID(ctx, id)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, domain.ErrPermissionNotFound
		}
		return nil, err
	}

	return permission, nil
}
