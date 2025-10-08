package auth

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/auth/handler"
	"github.com/umardev500/laundry/internal/feature/auth/repository"
	"github.com/umardev500/laundry/internal/feature/auth/service"
)

var ProviderSet = wire.NewSet(
	handler.NewHandler,
	service.NewService,
	repository.NewRedisRefreshTokenRepository,
	NewRoutes,
)
