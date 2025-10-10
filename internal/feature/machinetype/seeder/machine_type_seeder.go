package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/machinetype"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

type MachineTypeSeeder struct {
	client *entdb.Client
}

func NewMachineTypeSeeder(client *entdb.Client) *MachineTypeSeeder {
	return &MachineTypeSeeder{client: client}
}

func (s *MachineTypeSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding machine types...")
	conn := s.client.GetConn(ctx)

	types := []struct {
		ID          uuid.UUID
		TenantID    uuid.UUID
		Name        string
		Description string
		Capacity    int
	}{
		{uuid.MustParse("11111111-1111-1111-1111-111111111111"), uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), "Washer", "Front load washer", 1},
		{uuid.MustParse("22222222-2222-2222-2222-222222222222"), uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), "Dryer", "Commercial dryer", 1},
	}

	for _, t := range types {
		err := conn.MachineType.
			Create().
			SetID(t.ID).
			SetTenantID(t.TenantID).
			SetName(t.Name).
			SetDescription(t.Description).
			SetCapacity(t.Capacity).
			OnConflict(
				sql.ConflictColumns(machinetype.FieldName),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	log.Info().Msg("✅ Machine types seeded successfully.")
	return nil
}
