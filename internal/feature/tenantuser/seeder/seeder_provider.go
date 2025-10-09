package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type TenantUserSeederSet seeder.Seeder

func NewUserSeederSet(
	user *TenantUserSeeder,
) []TenantUserSeederSet {
	return []TenantUserSeederSet{
		user,
	}
}

var ProviderSet = wire.NewSet(
	NewTenantUserSeeder,
	NewUserSeederSet,
)
