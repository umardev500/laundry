package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/service/domain"
	"github.com/umardev500/laundry/internal/feature/service/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Service interface {
	Create(ctx *appctx.Context, s *domain.Service) (*domain.Service, error)
	List(ctx *appctx.Context, q *query.ListServiceQuery) (*pagination.PageData[domain.Service], error)
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Service, error)
	GetByName(ctx *appctx.Context, name string) (*domain.Service, error)
	Update(ctx *appctx.Context, s *domain.Service) (*domain.Service, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	Purge(ctx *appctx.Context, id uuid.UUID) error

	AreItemsAvailable(ctx *appctx.Context, ids []uuid.UUID) (*domain.AvailabilityResult, error)
}
