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

	// Combined machines for multiple tenants
	machines := []struct {
		ID            uuid.UUID
		TenantID      uuid.UUID
		MachineTypeID uuid.UUID
		Name          string
		Description   string
	}{
		// Tenant A
		{
			ID:            uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			TenantID:      uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			MachineTypeID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Name:          "Washer A",
			Description:   "Front load washer",
		},
		{
			ID:            uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			TenantID:      uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			MachineTypeID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Name:          "Dryer B",
			Description:   "Commercial dryer",
		},

		// Tenant B
		{
			ID:            uuid.MustParse("33333333-3333-3333-3333-333333333333"),
			TenantID:      uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
			MachineTypeID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Name:          "Washer C",
			Description:   "Front load washer",
		},
		{
			ID:            uuid.MustParse("44444444-4444-4444-4444-444444444444"),
			TenantID:      uuid.MustParse("bbbbbbbb-bbbb-bbbb-bbbb-bbbbbbbbbbbb"),
			MachineTypeID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Name:          "Dryer D",
			Description:   "Commercial dryer",
		},
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
