package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type FeatureService interface {
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Feature, error)
	List(ctx *appctx.Context, q *query.ListFeatureQuery) (*pagination.PageData[domain.Feature], error)
	Update(ctx *appctx.Context, feature *domain.Feature) (*domain.Feature, error)
	UpdateStatus(ctx *appctx.Context, feature *domain.Feature) (*domain.Feature, error)
}
