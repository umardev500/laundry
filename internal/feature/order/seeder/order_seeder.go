package seeder

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/order"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/types"
	"github.com/umardev500/laundry/pkg/utils"
)

type OrderSeeder struct {
	client *entdb.Client
}

func NewOrderSeeder(client *entdb.Client) *OrderSeeder {
	return &OrderSeeder{client: client}
}

func (s *OrderSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding orders...")

	conn := s.client.GetConn(ctx)
	userID1 := uuid.MustParse("22222222-1111-1111-1111-111111111111")
	userID2 := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	data := []struct {
		ID           uuid.UUID
		TenantID     uuid.UUID
		UserID       *uuid.UUID
		Status       types.OrderStatus
		TotalAmount  float64
		GuestName    *string
		GuestEmail   *string
		GuestPhone   *string
		GuestAddress *string
	}{
		{
			ID:          uuid.MustParse("aaaaaaaa-1111-1111-1111-aaaaaaaa1111"),
			TenantID:    uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			UserID:      utils.NilIfUUIDZero(userID1),
			Status:      types.OrderStatusPending,
			TotalAmount: 25.0,
		},
		{
			ID:           uuid.MustParse("bbbbbbbb-2222-2222-2222-bbbbbbbb2222"),
			TenantID:     uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			UserID:       nil,
			Status:       types.OrderStatusConfirmed,
			TotalAmount:  40.0,
			GuestName:    utils.NilIfZero("Jane Smith", ""),
			GuestEmail:   utils.NilIfZero("jane@example.com", ""),
			GuestPhone:   utils.NilIfZero("123-456-7890", ""),
			GuestAddress: utils.NilIfZero("123 Main St", ""),
		},
		{
			ID:          uuid.MustParse("cccccccc-3333-3333-3333-cccccccc3333"),
			TenantID:    uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			UserID:      utils.NilIfUUIDZero(userID2),
			Status:      types.OrderStatusPending,
			TotalAmount: 25.0,
		},
		{
			ID:          uuid.MustParse("dddddddd-4444-4444-4444-dddddddd4444"),
			TenantID:    uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
			UserID:      utils.NilIfUUIDZero(userID2),
			Status:      types.OrderStatusPending,
			TotalAmount: 25.0,
		},
	}

	for _, d := range data {
		exists, _ := conn.Order.Query().Where(
			order.IDEQ(d.ID),
		).Exist(ctx)

		if !exists {
			_, err := conn.Order.Create().
				SetID(d.ID).
				SetTenantID(d.TenantID).
				SetNillableUserID(d.UserID).
				SetStatus(order.Status(d.Status)).
				SetTotalAmount(d.TotalAmount).
				SetNillableGuestName(d.GuestName).
				SetNillableGuestEmail(d.GuestEmail).
				SetNillableGuestPhone(d.GuestPhone).
				SetNillableGuestAddress(d.GuestAddress).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to seed order %s: %w", d.ID, err)
			}
		}
	}

	return nil
}
