package tenant

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/tenant/handler"
	"github.com/umardev500/laundry/internal/feature/tenant/repository"
	"github.com/umardev500/laundry/internal/feature/tenant/service"
)

// ProviderSet wires up the dependencies for the tenant feature.
var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewService,
	repository.NewEntRepository,
)
