package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/servicecategory/domain"
	"github.com/umardev500/laundry/internal/feature/servicecategory/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	Create(ctx *appctx.Context, s *domain.ServiceCategory) (*domain.ServiceCategory, error)
	FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.ServiceCategory, error)
	FindByName(ctx *appctx.Context, name string) (*domain.ServiceCategory, error)
	Update(ctx *appctx.Context, s *domain.ServiceCategory) (*domain.ServiceCategory, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListServiceCategoryQuery) (*pagination.PageData[domain.ServiceCategory], error)
}
