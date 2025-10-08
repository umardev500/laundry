package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type PermissionService interface {
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Permission, error)
	List(ctx *appctx.Context, q *query.ListPermissionQuery) (*pagination.PageData[domain.Permission], error)
	Update(ctx *appctx.Context, permission *domain.Permission) (*domain.Permission, error)
	UpdateStatus(ctx *appctx.Context, id uuid.UUID, permission *domain.Permission) (*domain.Permission, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	Purge(ctx *appctx.Context, id uuid.UUID) error
}
