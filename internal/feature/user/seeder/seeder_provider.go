package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

func NewUserSeederSet(
	admin *AdminUserSeeder,
) []seeder.Seeder {
	return []seeder.Seeder{
		admin,
	}
}

var ProviderSet = wire.NewSet(
	NewAdminUserSeeder,
	NewUserSeederSet,
)
