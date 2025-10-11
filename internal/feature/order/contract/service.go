package contract

import (
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/order/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// OrderService defines the business logic for orders.
// For now, we only expose the List method.
type OrderService interface {
	List(ctx *appctx.Context, q *query.ListOrderQuery) (*pagination.PageData[domain.Order], error)
}
