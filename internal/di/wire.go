//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	fiberApp "github.com/umardev500/laundry/internal/app/fiber"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/auth"
	"github.com/umardev500/laundry/internal/feature/platformuser"
	"github.com/umardev500/laundry/internal/feature/user"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/redis"
)

var AppSet = wire.NewSet(
	entdb.NewEntClient,
	fiberApp.NewFiberApp,
	user.ProviderSet,
	platformuser.ProviderSet,
	auth.ProviderSet,
	router.NewRouter,
	redis.NewRedisClient,
	newRegistrars,
)

func Initialize(cfg *config.Config) (*router.Router, error) {
	wire.Build(AppSet)
	return nil, nil
}

func newRegistrars(
	userReg *user.Routes,
	platformUserReg *platformuser.Routes,
	authReg *auth.Routes,
) []router.RouteRegistrar {
	return []router.RouteRegistrar{
		userReg,
		platformUserReg,
		authReg,
	}
}
