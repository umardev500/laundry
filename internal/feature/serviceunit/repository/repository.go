package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/serviceunit/domain"
	"github.com/umardev500/laundry/internal/feature/serviceunit/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	Create(ctx *appctx.Context, s *domain.ServiceUnit) (*domain.ServiceUnit, error)
	FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.ServiceUnit, error)
	FindByName(ctx *appctx.Context, name string) (*domain.ServiceUnit, error)
	Update(ctx *appctx.Context, s *domain.ServiceUnit) (*domain.ServiceUnit, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListServiceUnitQuery) (*pagination.PageData[domain.ServiceUnit], error)
}
