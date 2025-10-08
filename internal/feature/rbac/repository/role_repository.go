package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Repository defines the data-access contract for roles.
type RoleRepository interface {
	Create(ctx *appctx.Context, role *domain.Role) (*domain.Role, error)
	FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.Role, error)
	FindByName(ctx *appctx.Context, name string) (*domain.Role, error)
	Update(ctx *appctx.Context, role *domain.Role) (*domain.Role, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListRoleQuery) (*pagination.PageData[domain.Role], error)
}
