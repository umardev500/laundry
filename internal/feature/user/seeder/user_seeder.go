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

// UserSeeder seeds multiple base users (e.g., Admin, Moderator, Operator)
type UserSeeder struct {
	client *entdb.Client
}

var _ seeder.Seeder = (*UserSeeder)(nil)

func NewUserSeeder(client *entdb.Client) *UserSeeder {
	return &UserSeeder{client: client}
}

func (s *UserSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding base users...")
	conn := s.client.GetConn(ctx)

	users := []struct {
		ID       uuid.UUID
		Email    string
		Password string
	}{
		{
			ID:       uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Email:    "admin@example.com",
			Password: "Admin123!",
		},
		{
			ID:       uuid.MustParse("22222222-1111-1111-1111-111111111111"),
			Email:    "moderator@example.com",
			Password: "Moderator123!",
		},
		{
			ID:       uuid.MustParse("33333333-1111-1111-1111-111111111111"),
			Email:    "operator@example.com",
			Password: "Operator123!",
		},
		{
			ID:       uuid.MustParse("44444444-1111-1111-1111-111111111111"),
			Email:    "user@example.com",
			Password: "User123!",
		},
	}

	for _, u := range users {
		hashed, err := security.Hash(u.Password)
		if err != nil {
			return err
		}

		err = conn.User.
			Create().
			SetID(u.ID).
			SetEmail(u.Email).
			SetPassword(hashed).
			OnConflict(
				sql.ConflictColumns(user.FieldEmail),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	log.Info().Msg("✅ Base users seeded successfully.")
	return nil
}
