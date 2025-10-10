package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	domainPkg "github.com/umardev500/laundry/internal/feature/servicecategory/domain"
	"github.com/umardev500/laundry/internal/feature/servicecategory/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	Create(ctx *appctx.Context, s *domainPkg.ServiceCategory) (*domainPkg.ServiceCategory, error)
	FindByID(ctx *appctx.Context, id uuid.UUID) (*domainPkg.ServiceCategory, error)
	FindByName(ctx *appctx.Context, name string) (*domainPkg.ServiceCategory, error)
	Update(ctx *appctx.Context, s *domainPkg.ServiceCategory) (*domainPkg.ServiceCategory, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListServiceCategoryQuery) (*pagination.PageData[domainPkg.ServiceCategory], error)
}
