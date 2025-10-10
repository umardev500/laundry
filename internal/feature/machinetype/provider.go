package machinetype

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/machinetype/handler"
	"github.com/umardev500/laundry/internal/feature/machinetype/repository"
	"github.com/umardev500/laundry/internal/feature/machinetype/service"
)

var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewService,
	repository.NewEntRepository,
)
