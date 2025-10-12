package contract

import (
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/orderitem/domain"
)

type Service interface {
	Create(ctx *appctx.Context, items []*domain.OrderItem) ([]*domain.OrderItem, error)
}
