package contract

import (
	"github.com/umardev500/laundry/internal/app/appctx"
	orderDomain "github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/payment/domain"
)

type Orchestrator interface {
	// SyncOrder syncs payment details with related entities like orders.
	SyncOrder(ctx *appctx.Context, ord *orderDomain.Order, pay *domain.Payment) error
}
