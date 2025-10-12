package repository

import (
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/orderitem/domain"
)

type Repository interface {
	// Create inserts a new order item or list of order items into the database.
	Create(ctx *appctx.Context, items []*domain.OrderItem) ([]*domain.OrderItem, error)
}
