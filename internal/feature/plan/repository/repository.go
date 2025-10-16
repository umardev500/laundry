package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/plan/domain"
	"github.com/umardev500/laundry/internal/feature/plan/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Repository defines the interface for Plan storage operations.
type Repository interface {
	Create(ctx *appctx.Context, p *domain.Plan) (*domain.Plan, error)
	FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error)
	FindByName(ctx *appctx.Context, name string) (*domain.Plan, error)
	Update(ctx *appctx.Context, p *domain.Plan) (*domain.Plan, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListPlanQuery) (*pagination.PageData[domain.Plan], error)
}
