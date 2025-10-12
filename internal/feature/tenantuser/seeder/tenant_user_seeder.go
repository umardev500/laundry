package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/tenantuser"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
	"github.com/umardev500/laundry/pkg/types"
)

type TenantUserSeeder struct {
	client *entdb.Client
}

var _ seeder.Seeder = (*TenantUserSeeder)(nil)

func NewTenantUserSeeder(client *entdb.Client) *TenantUserSeeder {
	return &TenantUserSeeder{client: client}
}

func (s *TenantUserSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding tenant users...")

	conn := s.client.GetConn(ctx)

	// Example dataset — you can adjust or expand this as needed.
	data := []struct {
		TenantID uuid.UUID
		UserID   uuid.UUID
		Status   types.Status
	}{
		{
			TenantID: uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			UserID:   uuid.MustParse("33333333-1111-1111-1111-111111111111"),
			Status:   types.StatusActive,
		},
		{
			TenantID: uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
			UserID:   uuid.MustParse("22222222-1111-1111-1111-111111111111"),
			Status:   types.StatusActive,
		},
	}

	for _, d := range data {
		err := conn.TenantUser.
			Create().
			SetTenantID(d.TenantID).
			SetUserID(d.UserID).
			SetStatus(tenantuser.Status(d.Status)).
			OnConflict(
				sql.ConflictColumns(tenantuser.FieldTenantID, tenantuser.FieldUserID),
			).
			UpdateNewValues().
			Exec(ctx)

		if err != nil {
			log.Error().Err(err).Msgf("❌ Failed to seed tenant user for tenant %s and user %s", d.TenantID, d.UserID)
			return err
		}
	}

	log.Info().Msg("✅ Tenant users seeded successfully.")
	return nil
}
