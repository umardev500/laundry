package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type PermissionRepository interface {
	FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.Permission, error)
	Update(ctx *appctx.Context, p *domain.Permission) (*domain.Permission, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListPermissionQuery) (*pagination.PageData[domain.Permission], error)
}
