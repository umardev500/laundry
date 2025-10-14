package orderstatushistory

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/handler"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/repository"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/service"
)

var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewStatusHistoryService,
	repository.NewEntStatusHistoryRepository,
)
