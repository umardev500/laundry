package seeder

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"

	"github.com/umardev500/laundry/ent/service"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

type ServiceSeeder struct {
	client *entdb.Client
}

func NewServiceSeeder(client *entdb.Client) *ServiceSeeder {
	return &ServiceSeeder{client: client}
}

func (s *ServiceSeeder) Seed(ctx context.Context) error {
	log.Info().Msg("🌱 Seeding services...")

	conn := s.client.GetConn(ctx)

	services := []struct {
		ID                uuid.UUID
		TenantID          uuid.UUID
		ServiceUnitID     uuid.UUID
		ServiceCategoryID uuid.UUID
		Name              string
		Price             float64
		Description       string
	}{
		{
			ID:                uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			TenantID:          uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			ServiceUnitID:     uuid.MustParse("11111111-1111-1111-1111-111111111111"), // Per Piece
			ServiceCategoryID: uuid.MustParse("11111111-1111-1111-1111-111111111111"),
			Name:              "Wash Small",
			Price:             2.5,
			Description:       "Small single load wash",
		},
		{
			ID:                uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			TenantID:          uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"),
			ServiceUnitID:     uuid.MustParse("22222222-2222-2222-2222-222222222222"), // Per Kilogram
			ServiceCategoryID: uuid.MustParse("22222222-2222-2222-2222-222222222222"),
			Name:              "Wash Large",
			Price:             5.0,
			Description:       "Large double load wash",
		},
	}

	for _, svc := range services {
		err := conn.Service.
			Create().
			SetID(svc.ID).
			SetTenantID(svc.TenantID).
			SetNillableServiceUnitID(&svc.ServiceUnitID).
			SetNillableServiceCategoryID(&svc.ServiceCategoryID).
			SetName(svc.Name).
			SetBasePrice(svc.Price).
			SetNillableDescription(&svc.Description).
			OnConflict(
				sql.ConflictColumns(service.FieldName),
			).
			UpdateNewValues().
			Exec(ctx)
		if err != nil {
			return err
		}
	}

	log.Info().Msg("✅ Services seeded successfully.")
	return nil
}
