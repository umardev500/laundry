package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent/machine"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

type MachineSeeder struct {
	client *entdb.Client
}

func NewMachineSeeder(client *entdb.Client) *MachineSeeder { return &MachineSeeder{client: client} }

func (s *MachineSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌿 Seeding machines...")
	conn := s.client.GetConn(ctx)
	machines := []struct {
		ID            uuid.UUID
		TenantID      uuid.UUID
		MachineTypeID uuid.UUID
		Name          string
		Description   string
	}{
		{uuid.MustParse("11111111-1111-1111-1111-111111111111"), uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), uuid.MustParse("11111111-1111-1111-1111-111111111111"), "Washer A", "Front load washer"},
		{uuid.MustParse("22222222-2222-2222-2222-222222222222"), uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), uuid.MustParse("22222222-2222-2222-2222-222222222222"), "Dryer B", "Commercial dryer"},
	}

	for _, m := range machines {
		err := conn.Machine.
			Create().
			SetID(m.ID).
			SetTenantID(m.TenantID).
			SetMachineTypeID(m.MachineTypeID).
			SetName(m.Name).
			SetDescription(m.Description).
			OnConflict(
				sql.ConflictColumns(machine.FieldName),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}
	log.Info().Msg("✅ Machines seeded successfully.")
	return nil
}
