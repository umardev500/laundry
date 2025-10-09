package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/tenant"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

// TenantSeeder seeds tenants into the database
type TenantSeeder struct {
	client *entdb.Client
}

var _ seeder.Seeder = (*TenantSeeder)(nil)

func NewTenantSeeder(client *entdb.Client) *TenantSeeder {
	return &TenantSeeder{client: client}
}

func (s *TenantSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding tenants...")

	conn := s.client.GetConn(ctx)

	tenants := []struct {
		ID    uuid.UUID
		Name  string
		Email string
		Phone string
	}{
		{uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), "Tenant One", "tenant1@example.com", "+1234567890"},
		{uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"), "Tenant Two", "tenant2@example.com", "+0987654321"},
	}

	for _, t := range tenants {
		err := conn.Tenant.
			Create().
			SetID(t.ID).
			SetName(t.Name).
			SetEmail(t.Email).
			SetPhone(t.Phone).
			OnConflict(
				sql.ConflictColumns(tenant.FieldEmail),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	log.Info().Msg("✅ Tenants seeded successfully.")
	return nil
}
