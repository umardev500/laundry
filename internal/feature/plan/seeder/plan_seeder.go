package seeder

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/umardev500/laundry/ent/plan"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/types"
)

type PlanSeeder struct {
	client *entdb.Client
}

func NewPlanSeeder(client *entdb.Client) *PlanSeeder {
	return &PlanSeeder{client: client}
}

func (s *PlanSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌱 Seeding plans...")

	conn := s.client.GetConn(ctx)

	plans := []struct {
		ID              uuid.UUID
		Name            string
		Description     string
		Price           float64
		BillingInterval types.BillingInterval
		Features        map[string]any
		Active          bool
		Permissions     []uuid.UUID
	}{
		{
			ID:              uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			Name:            "Basic",
			Description:     "Basic plan with limited features",
			Price:           9.99,
			BillingInterval: types.BillingIntervalMonthly,
			Features: map[string]any{
				"max_users": 1,
			},
			Active:      true,
			Permissions: []uuid.UUID{uuid.MustParse("aaaaaaaa-1111-1111-1111-aaaaaaaaaaaa")}, // Example permission ID
		},
		{
			ID:              uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
			Name:            "Premium",
			Description:     "Premium plan with full features",
			Price:           29.99,
			BillingInterval: types.BillingIntervalMonthly,
			Features: map[string]any{
				"max_users": 10,
			},
			Active:      false,
			Permissions: []uuid.UUID{uuid.MustParse("aaaaaaaa-1111-1111-1111-aaaaaaaaaaaa")}, // Example permission ID
		},
	}

	for _, p := range plans {
		// Insert plan
		err := conn.Plan.
			Create().
			SetID(p.ID).
			SetName(p.Name).
			SetNillableDescription(&p.Description).
			SetPrice(p.Price).
			SetBillingInterval(plan.BillingInterval(p.BillingInterval)).
			SetFeatures(p.Features).
			SetActive(p.Active).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			AddPermissionIDs(p.Permissions...). // Attach permissions
			OnConflict(
				sql.ConflictColumns(plan.FieldName),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}

	}

	log.Info().Msg("✅ Plans seeded successfully.")
	return nil
}
