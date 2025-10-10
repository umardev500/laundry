package servicecategory

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/servicecategory/handler"
	"github.com/umardev500/laundry/internal/feature/servicecategory/repository"
	"github.com/umardev500/laundry/internal/feature/servicecategory/service"
)

// ProviderSet defines all dependencies for the ServiceCategory feature.
var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewService,
	repository.NewEntRepository,
)
