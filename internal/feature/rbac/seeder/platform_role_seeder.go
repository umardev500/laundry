package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/role"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

// PlatformRoleSeeder seeds platform-wide roles (tenant_id = nil)
type PlatformRoleSeeder struct {
	client *entdb.Client
}

var _ seeder.Seeder = (*PlatformRoleSeeder)(nil)

func NewPlatformRoleSeeder(client *entdb.Client) *PlatformRoleSeeder {
	return &PlatformRoleSeeder{client: client}
}

func (s *PlatformRoleSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding platform roles...")

	conn := s.client.GetConn(ctx)

	roles := []struct {
		ID   uuid.UUID
		Name string
	}{
		{uuid.MustParse("22222222-2222-2222-2222-222222222222"), "Admin"},
		{uuid.MustParse("33333333-3333-3333-3333-333333333333"), "Moderator"},
	}

	for _, r := range roles {
		err := conn.Role.
			Create().
			SetID(r.ID).
			SetName(r.Name).
			OnConflict(
				sql.ConflictColumns(role.FieldID),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	log.Info().Msg("✅ Platform roles seeded successfully.")
	return nil
}
