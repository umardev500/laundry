package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/pagination"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/order/query"
)

type Repository interface {
	Create(ctx *appctx.Context, o *domain.Order) (*domain.Order, error)
	FindById(ctx *appctx.Context, id uuid.UUID, q *query.OrderQuery) (*domain.Order, error)
	List(ctx *appctx.Context, q *query.ListOrderQuery) (*pagination.PageData[domain.Order], error)
	Update(ctx *appctx.Context, o *domain.Order) (*domain.Order, error)
}
