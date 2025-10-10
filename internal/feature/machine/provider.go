package machine

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/machine/handler"
	"github.com/umardev500/laundry/internal/feature/machine/repository"
	"github.com/umardev500/laundry/internal/feature/machine/service"
)

var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewService,
	repository.NewEntRepository,
)
