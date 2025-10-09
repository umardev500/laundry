package seeder

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/infra/database/seeder"
)

type RBACSeeder seeder.Seeder

func NewRbacSeederSet(
	feature *FeatureSeeder,
	rolePlatform *PlatformRoleSeeder,
	roleTenant *TenantRoleSeeder,
	permission *PermissionSeeder,
) []RBACSeeder {
	return []RBACSeeder{
		feature,
		rolePlatform,
		roleTenant,
		permission,
	}
}

var ProviderSet = wire.NewSet(
	NewFeatureSeeder,
	NewPlatformRoleSeeder,
	NewTenantRoleSeeder,
	NewPermissionSeeder,
	NewRbacSeederSet,
)
