package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/service/domain"
	"github.com/umardev500/laundry/internal/feature/service/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	Create(ctx *appctx.Context, s *domain.Service) (*domain.Service, error)
	FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.Service, error)
	FindByName(ctx *appctx.Context, name string) (*domain.Service, error)
	Update(ctx *appctx.Context, s *domain.Service) (*domain.Service, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListServiceQuery) (*pagination.PageData[domain.Service], error)
	AreItemsAvailable(ctx *appctx.Context, ids []uuid.UUID) (*domain.AvailabilityResult, error)
}
