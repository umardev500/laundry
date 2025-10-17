package subscription

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/subscription/handler"
	"github.com/umardev500/laundry/internal/feature/subscription/repository"
	"github.com/umardev500/laundry/internal/feature/subscription/service"
)

var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewSubscriptionService,
	repository.NewEntRepository,
)
