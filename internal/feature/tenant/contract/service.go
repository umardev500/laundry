package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/tenant/domain"
	"github.com/umardev500/laundry/internal/feature/tenant/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Service interface {
	Create(ctx *appctx.Context, tenant *domain.Tenant) (*domain.Tenant, error)
	List(ctx *appctx.Context, query *query.ListTenantQuery) (*pagination.PageData[domain.Tenant], error)
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Tenant, error)
	GetByEmail(ctx *appctx.Context, email string) (*domain.Tenant, error)
	Update(ctx *appctx.Context, tenant *domain.Tenant) (*domain.Tenant, error)
	UpdateStatus(ctx *appctx.Context, tenant *domain.Tenant) (*domain.Tenant, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	Purge(ctx *appctx.Context, id uuid.UUID) error
}
