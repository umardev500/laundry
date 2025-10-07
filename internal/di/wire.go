//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	fiberApp "github.com/umardev500/laundry/internal/app/fiber"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/user"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
)

var AppSet = wire.NewSet(
	entdb.NewEntClient,
	fiberApp.NewFiberApp,
	user.ProviderSet,
	router.NewRouter,
	newRegistrars,
)

func Initialize(cfg *config.Config) (*router.Router, error) {
	wire.Build(AppSet)
	return nil, nil
}

func newRegistrars(
	userReg router.RouteRegistrar,
) []router.RouteRegistrar {
	return []router.RouteRegistrar{
		userReg,
	}
}
