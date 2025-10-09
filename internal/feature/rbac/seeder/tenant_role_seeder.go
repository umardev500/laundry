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

// TenantRoleSeeder seeds tenant-specific roles
type TenantRoleSeeder struct {
	client *entdb.Client
}

var _ seeder.Seeder = (*TenantRoleSeeder)(nil)

func NewTenantRoleSeeder(client *entdb.Client) *TenantRoleSeeder {
	return &TenantRoleSeeder{client: client}
}

func (s *TenantRoleSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding tenant roles...")

	conn := s.client.GetConn(ctx)

	tenantRoles := []struct {
		ID       uuid.UUID
		Name     string
		TenantID uuid.UUID
	}{
		{uuid.MustParse("44444444-4444-4444-4444-444444444444"), "TenantAdmin", uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")},
		{uuid.MustParse("55555555-5555-5555-5555-555555555555"), "TenantUser", uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")},
	}

	for _, r := range tenantRoles {
		err := conn.Role.
			Create().
			SetID(r.ID).
			SetName(r.Name).
			SetTenantID(r.TenantID).
			OnConflict(
				sql.ConflictColumns(role.FieldName, role.FieldTenantID),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	log.Info().Msg("✅ Tenant roles seeded successfully.")
	return nil
}
