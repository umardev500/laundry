package repository

import (
	"github.com/umardev500/laundry/pkg/pagination"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/order/query"
)

type Repository interface {
	List(ctx *appctx.Context, q *query.ListOrderQuery) (*pagination.PageData[domain.Order], error)
}
