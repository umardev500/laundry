package platformuser

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/platformuser/handler"
	"github.com/umardev500/laundry/internal/feature/platformuser/repository"
	"github.com/umardev500/laundry/internal/feature/platformuser/service"
)

var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewService,
	repository.NewEntRepository,
)
