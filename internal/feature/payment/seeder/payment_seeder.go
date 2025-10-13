package seeder

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/payment"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/types"
	"github.com/umardev500/laundry/pkg/utils"
)

type PaymentSeeder struct {
	client *entdb.Client
}

func NewPaymentSeeder(client *entdb.Client) *PaymentSeeder {
	return &PaymentSeeder{client: client}
}

func (s *PaymentSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding payments...")

	conn := s.client.GetConn(ctx)

	userID1 := uuid.MustParse("11111111-1111-1111-1111-111111111111")
	userID4 := uuid.MustParse("44444444-1111-1111-1111-111111111111")
	tenantA := uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	tenantB := uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb")
	cashID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	data := []struct {
		ID              uuid.UUID
		UserID          *uuid.UUID
		TenantID        uuid.UUID
		RefID           uuid.UUID
		RefType         types.PaymentType
		PaymentMethodID uuid.UUID
		Amount          float64
		Status          types.PaymentStatus
		PaidAt          *time.Time
	}{
		{
			ID:              uuid.MustParse("11111111-aaaa-aaaa-aaaa-111111111111"),
			UserID:          utils.NilIfUUIDZero(userID1),
			TenantID:        tenantA,
			RefID:           uuid.MustParse("cccccccc-3333-3333-3333-cccccccc3333"), // order ID
			RefType:         types.PaymentTypeOrder,
			PaymentMethodID: cashID,
			Amount:          25.0,
			Status:          types.PaymentStatusPending,
		},
		{
			ID:              uuid.MustParse("22222222-bbbb-bbbb-bbbb-222222222222"),
			UserID:          nil, // guest payment
			TenantID:        tenantA,
			RefID:           uuid.MustParse("bbbbbbbb-2222-2222-2222-bbbbbbbb2222"),
			RefType:         types.PaymentTypeOrder,
			PaymentMethodID: cashID,
			Amount:          40.0,
			Status:          types.PaymentStatusPaid,
			PaidAt:          ptrTime(time.Now().Add(-24 * time.Hour)),
		},
		{
			ID:              uuid.MustParse("33333333-cccc-cccc-cccc-333333333333"),
			UserID:          utils.NilIfUUIDZero(userID4),
			TenantID:        tenantB,
			RefID:           uuid.MustParse("dddddddd-4444-4444-4444-dddddddd4444"),
			RefType:         types.PaymentTypeOrder,
			PaymentMethodID: cashID,
			Amount:          25.0,
			Status:          types.PaymentStatusFailed,
		},
	}

	for _, d := range data {
		exists, _ := conn.Payment.Query().
			Where(payment.IDEQ(d.ID)).
			Exist(ctx)

		if !exists {
			q := conn.Payment.Create().
				SetID(d.ID).
				SetNillableUserID(d.UserID).
				SetTenantID(d.TenantID).
				SetRefID(d.RefID).
				SetRefType(payment.RefType(d.RefType)).
				SetPaymentMethodID(d.PaymentMethodID).
				SetAmount(d.Amount).
				SetStatus(payment.Status(d.Status)).
				SetNillablePaidAt(d.PaidAt)

			if d.RefType == types.PaymentTypeOrder {
				q.SetOrderID(d.RefID)
			}

			_, err := q.Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to seed payment %s: %w", d.ID, err)
			}
		}
	}

	return nil
}

// helper to get a pointer to time.Time
func ptrTime(t time.Time) *time.Time {
	return &t
}
