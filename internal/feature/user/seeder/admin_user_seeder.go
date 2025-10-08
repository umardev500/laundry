package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
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

	err = conn.User.
		Create().
		SetID(uuid.MustParse("11111111-1111-1111-1111-111111111111")).
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

	return nil
}
