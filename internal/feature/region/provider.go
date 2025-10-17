package region

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/region/handler"
	"github.com/umardev500/laundry/internal/feature/region/repository"
	"github.com/umardev500/laundry/internal/feature/region/service"
)

var ProviderSet = wire.NewSet(
	handler.NewHandler,
	service.NewService,
	repository.NewRepository,
	NewRoutes,
)
