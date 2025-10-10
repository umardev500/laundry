package seeder

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent/servicecategory"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

type ServiceCategorySeeder struct {
	client *entdb.Client
}

func NewServiceCategorySeeder(client *entdb.Client) *ServiceCategorySeeder {
	return &ServiceCategorySeeder{client: client}
}

func (s *ServiceCategorySeeder) Seed(ctx context.Context) error {
	conn := s.client.GetConn(ctx)

	data := []struct {
		ID          uuid.UUID
		TenantID    uuid.UUID
		Name        string
		Description string
	}{
		{uuid.MustParse("11111111-1111-1111-1111-111111111111"), uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), "Laundry", "Basic laundry services"},
		{uuid.MustParse("22222222-2222-2222-2222-222222222222"), uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), "Dry Cleaning", "Premium dry cleaning"},
		{uuid.MustParse("33333333-3333-3333-3333-333333333333"), uuid.MustParse("aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"), "Ironing", "Ironing and folding"},
	}

	for _, d := range data {
		exists, _ := conn.ServiceCategory.Query().Where(
			servicecategory.NameEQ(d.Name),
			servicecategory.TenantIDEQ(d.TenantID),
		).Exist(ctx)

		if !exists {
			_, err := conn.ServiceCategory.Create().
				SetID(d.ID).
				SetTenantID(d.TenantID).
				SetName(d.Name).
				SetDescription(d.Description).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("failed to seed category %s: %w", d.Name, err)
			}
		}
	}

	return nil
}
