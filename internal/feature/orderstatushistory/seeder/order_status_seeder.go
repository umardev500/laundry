package seeder

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/orderstatushistory"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/types"
)

type OrderStatusHistorySeeder struct {
	client *entdb.Client
}

func NewOrderStatusHistorySeeder(client *entdb.Client) *OrderStatusHistorySeeder {
	return &OrderStatusHistorySeeder{client: client}
}

func (s *OrderStatusHistorySeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding order status history...")

	conn := s.client.GetConn(ctx)

	data := []struct {
		ID      uuid.UUID
		OrderID uuid.UUID
		Status  types.OrderStatus
		Notes   *string
	}{
		// Order A has Pending → Confirmed → Completed
		{
			ID:      uuid.MustParse("11111111-aaaa-aaaa-aaaa-111111111111"),
			OrderID: uuid.MustParse("aaaaaaaa-1111-1111-1111-aaaaaaaa1111"),
			Status:  types.OrderStatusPending,
			Notes:   nil,
		},
		{
			ID:      uuid.MustParse("11111111-aaaa-aaaa-aaaa-111111111112"),
			OrderID: uuid.MustParse("aaaaaaaa-1111-1111-1111-aaaaaaaa1111"),
			Status:  types.OrderStatusConfirmed,
			Notes:   nil,
		},
		{
			ID:      uuid.MustParse("11111111-aaaa-aaaa-aaaa-111111111113"),
			OrderID: uuid.MustParse("aaaaaaaa-1111-1111-1111-aaaaaaaa1111"),
			Status:  types.OrderStatusCompleted,
			Notes:   nil,
		},

		// Order B has Confirmed → Completed
		{
			ID:      uuid.MustParse("22222222-bbbb-bbbb-bbbb-222222222222"),
			OrderID: uuid.MustParse("bbbbbbbb-2222-2222-2222-bbbbbbbb2222"),
			Status:  types.OrderStatusConfirmed,
			Notes:   nil,
		},
		{
			ID:      uuid.MustParse("22222222-bbbb-bbbb-bbbb-222222222223"),
			OrderID: uuid.MustParse("bbbbbbbb-2222-2222-2222-bbbbbbbb2222"),
			Status:  types.OrderStatusCompleted,
			Notes:   nil,
		},

		// Order C has Pending only
		{
			ID:      uuid.MustParse("33333333-cccc-cccc-cccc-333333333333"),
			OrderID: uuid.MustParse("cccccccc-3333-3333-3333-cccccccc3333"),
			Status:  types.OrderStatusPending,
			Notes:   nil,
		},

		// Order D has Pending → Cancelled
		{
			ID:      uuid.MustParse("44444444-dddd-dddd-dddd-444444444444"),
			OrderID: uuid.MustParse("dddddddd-4444-4444-4444-dddddddd4444"),
			Status:  types.OrderStatusPending,
			Notes:   nil,
		},
		{
			ID:      uuid.MustParse("44444444-dddd-dddd-dddd-444444444445"),
			OrderID: uuid.MustParse("dddddddd-4444-4444-4444-dddddddd4444"),
			Status:  types.OrderStatusCancelled,
			Notes:   nil,
		},
	}

	for _, d := range data {
		exists, _ := conn.OrderStatusHistory.Query().Where(
			orderstatushistory.IDEQ(d.ID),
		).Exist(ctx)

		if !exists {
			_, err := conn.OrderStatusHistory.Create().
				SetID(d.ID).
				SetOrderID(d.OrderID).
				SetStatus(orderstatushistory.Status(d.Status)).
				SetCreatedAt(time.Now().Add(1 * time.Hour)).
				SetNillableNotes(d.Notes).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to seed order status history %s: %w", d.ID, err)
			}
		} else {
			fmt.Println("exist")
		}
	}

	return nil
}
