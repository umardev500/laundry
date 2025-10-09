package seeder

import (
	"context"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/feature"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
	"github.com/umardev500/laundry/pkg/types"
)

// FeatureSeeder seeds system features (e.g., Orders, Users, Payments, etc.)
type FeatureSeeder struct {
	client *entdb.Client
}

// Ensure FeatureSeeder implements seeder.Seeder interface.
var _ seeder.Seeder = (*FeatureSeeder)(nil)

// NewFeatureSeeder creates a new instance of FeatureSeeder.
func NewFeatureSeeder(client *entdb.Client) *FeatureSeeder {
	return &FeatureSeeder{client: client}
}

// Seed seeds predefined features.
func (s *FeatureSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding system features...")
	conn := s.client.GetConn(ctx)

	features := []struct {
		ID          uuid.UUID
		Name        string
		Description string
	}{
		{
			ID:          uuid.MustParse("22222222-1111-1111-1111-111111111111"),
			Name:        "users",
			Description: "Manage system users and their access",
		},
		{
			ID:          uuid.MustParse("22222222-1111-1111-1111-222222222222"),
			Name:        "roles",
			Description: "Manage roles and permissions",
		},
		{
			ID:          uuid.MustParse("22222222-1111-1111-1111-333333333333"),
			Name:        "permissions",
			Description: "Manage granular permission access",
		},
		{
			ID:          uuid.MustParse("22222222-1111-1111-1111-444444444444"),
			Name:        "tenants",
			Description: "Tenant management and configurations",
		},
		{
			ID:          uuid.MustParse("22222222-1111-1111-1111-555555555555"),
			Name:        "laundry_orders",
			Description: "Handle laundry order lifecycle",
		},
	}

	for _, f := range features {
		err := conn.Feature.
			Create().
			SetID(f.ID).
			SetName(f.Name).
			SetDescription(f.Description).
			SetStatus(feature.Status(types.StatusActive)).
			SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			OnConflict(
				sql.ConflictColumns(feature.FieldName),
			).
			UpdateNewValues().
			Exec(ctx)

		if err != nil {
			return err
		}
	}

	log.Info().Msg("✅ System features seeded successfully.")
	return nil
}
