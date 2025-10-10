package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/umardev500/laundry/ent/serviceunit"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

// ServiceUnitSeeder seeds default service units for testing and local environments.
type ServiceUnitSeeder struct {
	client *entdb.Client
}

// NewServiceUnitSeeder creates a new instance of ServiceUnitSeeder.
func NewServiceUnitSeeder(client *entdb.Client) *ServiceUnitSeeder {
	return &ServiceUnitSeeder{client: client}
}

// Seed inserts predefined service units into the database.
// This seeder uses hardcoded UUIDs for predictable test data.
func (s *ServiceUnitSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌱 Seeding service units...")

	conn := s.client.GetConn(ctx)

	units := []struct {
		ID       uuid.UUID
		TenantID uuid.UUID
		Name     string
		Symbol   string
	}{
		{
			ID:       uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			TenantID: uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			Name:     "Per Piece",
			Symbol:   "pc",
		},
		{
			ID:       uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			TenantID: uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			Name:     "Per Kilogram",
			Symbol:   "kg",
		},
		{
			ID:       uuid.MustParse("33333333-3333-3333-3333-333333333333"),
			TenantID: uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			Name:     "Per Set",
			Symbol:   "set",
		},
	}

	for _, u := range units {
		err := conn.ServiceUnit.
			Create().
			SetID(u.ID).
			SetTenantID(u.TenantID).
			SetName(u.Name).
			SetSymbol(u.Symbol).
			OnConflict(sql.ConflictColumns(serviceunit.FieldName)).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	log.Info().Msg("✅ Service units seeded successfully.")
	return nil
}
