package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/tenantuser/domain"
	"github.com/umardev500/laundry/internal/feature/tenantuser/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	Create(ctx *appctx.Context, tu *domain.TenantUser) (*domain.TenantUser, error)
	FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.TenantUser, error)
	FindByUser(ctx *appctx.Context, userID uuid.UUID) ([]*domain.TenantUser, error)
	FindByUserAndTenant(ctx *appctx.Context, userID, tenantID uuid.UUID) (*domain.TenantUser, error)
	List(ctx *appctx.Context, q *query.ListTenantUserQuery) (*pagination.PageData[domain.TenantUser], error)
	Update(ctx *appctx.Context, tu *domain.TenantUser) (*domain.TenantUser, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
}
