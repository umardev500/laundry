package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/platformuser"
	"github.com/umardev500/laundry/ent/user"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
	"github.com/umardev500/laundry/pkg/security"
)

type AdminUserSeeder struct {
	client *entdb.Client
}

// Ensure UserSeeder implements the Seeder interface
var _ seeder.Seeder = (*AdminUserSeeder)(nil)

func NewAdminUserSeeder(clinet *entdb.Client) *AdminUserSeeder {
	return &AdminUserSeeder{
		client: clinet,
	}
}

func (s *AdminUserSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding admin user...")
	conn := s.client.GetConn(ctx)

	hashed, err := security.Hash("Admin123!")
	if err != nil {
		return err
	}

	adminID := uuid.MustParse("11111111-1111-1111-1111-111111111111")

	// Upsert the User
	err = conn.User.
		Create().
		SetID(adminID).
		SetEmail("admin@example.com").
		SetPassword(hashed).
		OnConflict(
			sql.ConflictColumns(user.FieldEmail),
		).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}

	// Upsert the Platformuser for this admin
	err = conn.PlatformUser.
		Create().
		SetID(uuid.MustParse("11111111-1111-1111-1111-111111111111")).
		SetUserID(adminID).
		SetStatus(platformuser.Status(user.StatusActive)).
		OnConflict(
			sql.ConflictColumns(platformuser.FieldUserID),
		).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return err
	}

	log.Info().Msg("✅ Admin user and platform user seeded successfully.")
	return nil
}
