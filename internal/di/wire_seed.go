//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/seeder"

	userSeeder "github.com/umardev500/laundry/internal/feature/user/seeder"
)

func InitialzeSeeder(cfg *config.Config) ([]seeder.Seeder, error) {
	wire.Build(
		entdb.NewEntClient,
		userSeeder.ProviderSet,
	)

	return nil, nil
}
