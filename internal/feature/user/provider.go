package user

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/user/handler"
	"github.com/umardev500/laundry/internal/feature/user/repository"
	"github.com/umardev500/laundry/internal/feature/user/service"
)

var ProviderSet = wire.NewSet(
	handler.NewHandler,
	service.NewService,
	repository.NewEntRepository,
	NewRoutes,
)
