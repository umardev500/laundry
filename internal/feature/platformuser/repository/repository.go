package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/platformuser/domain"
	"github.com/umardev500/laundry/internal/feature/platformuser/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	Create(ctx *appctx.Context, pu *domain.PlatformUser) (*domain.PlatformUser, error)
	FindById(ctx *appctx.Context, id uuid.UUID) (*domain.PlatformUser, error)
	FindByIdUserID(ctx *appctx.Context, userID uuid.UUID) (*domain.PlatformUser, error)
	Update(ctx *appctx.Context, pu *domain.PlatformUser) (*domain.PlatformUser, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListPlatformUserQuery) (*pagination.PageData[domain.PlatformUser], error)
}
