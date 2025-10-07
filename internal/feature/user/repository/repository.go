package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/user/domain"
	"github.com/umardev500/laundry/internal/feature/user/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	Create(ctx *appctx.Context, user *domain.User) (*domain.User, error)
	FindById(ctx *appctx.Context, id uuid.UUID) (*domain.User, error)
	FindByEmail(ctx *appctx.Context, email string) (*domain.User, error)
	Update(ctx *appctx.Context, user *domain.User) (*domain.User, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListUserQuery) (*pagination.PageData[domain.User], error)
}
