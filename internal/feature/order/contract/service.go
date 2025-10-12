package contract

import (
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/order/query"
	"github.com/umardev500/laundry/pkg/pagination"

	paymentDomain "github.com/umardev500/laundry/internal/feature/payment/domain"
)

// OrderService defines the business logic for orders.
// For now, we only expose the List method.
type OrderService interface {
	CreatePayment(ctx *appctx.Context, o *domain.Order) (*paymentDomain.Payment, error)
	GuestOrder(ctx *appctx.Context, o *domain.Order) (*domain.Order, error)
	List(ctx *appctx.Context, q *query.ListOrderQuery) (*pagination.PageData[domain.Order], error)
	Preview(ctx *appctx.Context, o *domain.Order) (*domain.Order, error)
}
