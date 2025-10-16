package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/plan/domain"
	"github.com/umardev500/laundry/internal/feature/plan/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Plan defines the business service interface for managing Plans.
// This interface mirrors the repository operations but allows for
// domain-level validations, rules, and orchestration.
type Plan interface {
	Create(ctx *appctx.Context, p *domain.Plan) (*domain.Plan, error)
	List(ctx *appctx.Context, q *query.ListPlanQuery) (*pagination.PageData[domain.Plan], error)
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error)
	GetByName(ctx *appctx.Context, name string) (*domain.Plan, error)
	Update(ctx *appctx.Context, p *domain.Plan) (*domain.Plan, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error // soft delete
	Purge(ctx *appctx.Context, id uuid.UUID) error  // hard delete

	// Status operations
	Activate(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error)
	Deactivate(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error)
	Restore(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error)
}
