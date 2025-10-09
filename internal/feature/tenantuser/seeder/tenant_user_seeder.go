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
	log.Info().Msg("🌿 Seeding tenant user...")

	conn := s.client.GetConn(ctx)

	tenantID := uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa")
	userID := uuid.MustParse("33333333-1111-1111-1111-111111111111")

	err := conn.TenantUser.
		Create().
		SetTenantID(tenantID).
		SetUserID(userID).
		SetStatus(tenantuser.Status(types.StatusActive)).
		OnConflict(
			sql.ConflictColumns(tenantuser.FieldTenantID, tenantuser.FieldUserID),
		).
		UpdateNewValues().
		Exec(ctx)

	if err != nil {
		return err
	}

	log.Info().Msg("✅ Tenant user seeded successfully.")
	return nil
}
