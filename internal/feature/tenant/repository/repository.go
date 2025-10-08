package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/tenant/domain"
	"github.com/umardev500/laundry/internal/feature/tenant/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	Create(ctx *appctx.Context, tenant *domain.Tenant) (*domain.Tenant, error)
	FindById(ctx *appctx.Context, id uuid.UUID) (*domain.Tenant, error)
	FindByEmail(ctx *appctx.Context, email string) (*domain.Tenant, error)
	Update(ctx *appctx.Context, tenant *domain.Tenant) (*domain.Tenant, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListTenantQuery) (*pagination.PageData[domain.Tenant], error)
}
