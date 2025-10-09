//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/seeder"

	platformUserSeeder "github.com/umardev500/laundry/internal/feature/platformuser/seeder"
	rbacSeeder "github.com/umardev500/laundry/internal/feature/rbac/seeder"
	tenantSeeder "github.com/umardev500/laundry/internal/feature/tenant/seeder"
	userSeeder "github.com/umardev500/laundry/internal/feature/user/seeder"
)

// func toSeeders[T any](src []T) []seeder.Seeder {
// 	out := make([]seeder.Seeder, len(src))
// 	for i, s := range src {
// 		out[i] = any(s).(seeder.Seeder)
// 	}
// 	return out
// }

func NewSeederSet(
	rbac []rbacSeeder.RBACSeeder,
	user []userSeeder.UserSeederSet,
	tenant []tenantSeeder.TenantSeederSet,
	platformUser []platformUserSeeder.PlatformUserSeederSet,
) []seeder.Seeder {
	var all []seeder.Seeder

	// append all tenant seeder
	for _, s := range tenant {
		all = append(all, s)
	}

	// append all rbac seeder
	for _, s := range rbac {
		all = append(all, s)
	}

	// append all user seeder
	for _, s := range user {
		all = append(all, s)
	}

	// append all platform user seeder
	for _, s := range platformUser {
		all = append(all, s)
	}

	return all
}

func InitialzeSeeder(cfg *config.Config) ([]seeder.Seeder, error) {
	wire.Build(
		entdb.NewEntClient,
		userSeeder.ProviderSet,
		rbacSeeder.ProviderSet,
		tenantSeeder.ProviderSet,
		platformUserSeeder.ProviderSet,
		NewSeederSet,
	)

	return nil, nil
}
