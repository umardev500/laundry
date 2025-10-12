package seeder

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent/orderitem"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

type OrderItemSeeder struct {
	client *entdb.Client
}

func NewOrderItemSeeder(client *entdb.Client) *OrderItemSeeder {
	return &OrderItemSeeder{client: client}
}

func (s *OrderItemSeeder) Seed(ctx context.Context) error {
	conn := s.client.GetConn(ctx)

	data := []struct {
		ID        uuid.UUID
		OrderID   uuid.UUID
		ServiceID uuid.UUID
		Quantity  float64
		Price     float64
		Subtotal  float64
		Total     float64
	}{
		{
			ID:        uuid.MustParse("11111111-aaaa-1111-aaaa-111111111111"),
			OrderID:   uuid.MustParse("aaaaaaaa-1111-1111-1111-aaaaaaaa1111"),
			ServiceID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Quantity:  2,
			Price:     2.5,
			Subtotal:  5.0,
			Total:     5.0,
		},
		{
			ID:        uuid.MustParse("22222222-bbbb-2222-bbbb-222222222222"),
			OrderID:   uuid.MustParse("bbbbbbbb-2222-2222-2222-bbbbbbbb2222"),
			ServiceID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Quantity:  1,
			Price:     5.0,
			Subtotal:  5.0,
			Total:     5.0,
		},
		{
			ID:        uuid.MustParse("33333333-cccc-3333-cccc-333333333333"),
			OrderID:   uuid.MustParse("cccccccc-3333-3333-3333-cccccccc3333"),
			ServiceID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Quantity:  1,
			Price:     15.0,
			Subtotal:  15.0,
			Total:     15.0,
		},
		{
			ID:        uuid.MustParse("44444444-dddd-4444-dddd-444444444444"),
			OrderID:   uuid.MustParse("dddddddd-4444-4444-4444-dddddddd4444"),
			ServiceID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Quantity:  1,
			Price:     15.0,
			Subtotal:  15.0,
			Total:     15.0,
		},
	}

	for _, d := range data {
		exists, _ := conn.OrderItem.Query().Where(orderitem.IDEQ(d.ID)).Exist(ctx)
		if !exists {
			_, err := conn.OrderItem.Create().
				SetID(d.ID).
				SetOrderID(d.OrderID).
				SetServiceID(d.ServiceID).
				SetQuantity(d.Quantity).
				SetPrice(d.Price).
				SetSubtotal(d.Subtotal).
				SetTotalAmount(d.Total).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to seed order item %s: %w", d.ID, err)
			}
		}
	}

	return nil
}
