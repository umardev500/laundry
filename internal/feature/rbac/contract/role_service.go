package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Service defines all use cases for role management.
type Service interface {
	Create(ctx *appctx.Context, role *domain.Role) (*domain.Role, error)
	List(ctx *appctx.Context, query *query.ListRoleQuery) (*pagination.PageData[domain.Role], error)
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Role, error)
	GetByName(ctx *appctx.Context, name string) (*domain.Role, error)
	Update(ctx *appctx.Context, role *domain.Role) (*domain.Role, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	Purge(ctx *appctx.Context, id uuid.UUID) error
}
