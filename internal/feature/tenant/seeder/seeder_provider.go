package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type TenantSeederSet seeder.Seeder

func NewTenantSeederSet(
	tenant *TenantSeeder,
) []TenantSeederSet {
	return []TenantSeederSet{
		tenant,
	}
}

var ProviderSet = wire.NewSet(
	NewTenantSeederSet,
	NewTenantSeeder,
)
