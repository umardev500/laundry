package repository

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/orderitem/domain"
	"github.com/umardev500/laundry/internal/feature/orderitem/mapper"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

type entImpl struct {
	client *entdb.Client
}

func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{client: client}
}

// Create inserts one or more order items into the database.
func (r *entImpl) Create(ctx *appctx.Context, items []*domain.OrderItem) ([]*domain.OrderItem, error) {
	if len(items) == 0 {
		return []*domain.OrderItem{}, nil
	}

	conn := r.client.GetConn(ctx)
	bulk := make([]*ent.OrderItemCreate, 0, len(items))

	for _, item := range items {
		builder := conn.OrderItem.Create().
			SetOrderID(item.OrderID).
			SetServiceID(item.ServiceID).
			SetQuantity(item.Quantity).
			SetPrice(item.Price).
			SetSubtotal(item.Subtotal).
			SetTotalAmount(item.TotalAmount)

		bulk = append(bulk, builder)
	}

	created, err := conn.OrderItem.CreateBulk(bulk...).Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntList(created), nil
}
