package seeder

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/umardev500/laundry/ent/subscription"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/types"
)

type SubscriptionSeeder struct {
	client *entdb.Client
}

func NewSubscriptionSeeder(client *entdb.Client) *SubscriptionSeeder {
	return &SubscriptionSeeder{client: client}
}

func (s *SubscriptionSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌱 Seeding subscriptions...")

	conn := s.client.GetConn(ctx)

	subscriptions := []struct {
		ID        uuid.UUID
		TenantID  uuid.UUID
		PlanID    uuid.UUID
		Status    types.SubscriptionStatus
		StartDate *time.Time
		EndDate   *time.Time
	}{
		{
			ID:       uuid.MustParse("cccccccc-cccc-cccc-cccc-cccccccccccc"),
			TenantID: uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			PlanID:   uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), // Basic
			Status:   types.SubscriptionStatusActive,
			StartDate: func() *time.Time {
				t := time.Now().AddDate(0, 0, -10)
				return &t
			}(),
			EndDate: func() *time.Time {
				t := time.Now().AddDate(0, 1, -10)
				return &t
			}(),
		},
		{
			ID:       uuid.MustParse("dddddddd-dddd-dddd-dddd-dddddddddddd"),
			TenantID: uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
			PlanID:   uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"), // Premium
			Status:   types.SubscriptionStatusPending,
			StartDate: func() *time.Time {
				t := time.Now().AddDate(0, 0, -5)
				return &t
			}(),
			EndDate: func() *time.Time {
				t := time.Now().AddDate(0, 1, -5)
				return &t
			}(),
		},
	}

	for _, sub := range subscriptions {
		err := conn.Subscription.
			Create().
			SetID(sub.ID).
			SetTenantID(sub.TenantID).
			SetPlanID(sub.PlanID).
			SetStatus(subscription.Status(sub.Status)).
			SetNillableStartDate(sub.StartDate).
			SetNillableEndDate(sub.EndDate).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			OnConflict(
				sql.ConflictColumns(subscription.FieldID),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	log.Info().Msg("✅ Subscriptions seeded successfully.")
	return nil
}
