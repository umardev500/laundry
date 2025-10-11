package order

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/order/handler"
	"github.com/umardev500/laundry/internal/feature/order/repository"
	"github.com/umardev500/laundry/internal/feature/order/service"
)

// ProviderSet defines all dependencies for the Order feature.
var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewOrderService,
	repository.NewEntRepository,
)
