package service

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/service/handler"
	"github.com/umardev500/laundry/internal/feature/service/repository"
	"github.com/umardev500/laundry/internal/feature/service/service"
)

var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewService,
	repository.NewEntRepository,
)
