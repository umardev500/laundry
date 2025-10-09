package tenantuser

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/tenantuser/handler"
	"github.com/umardev500/laundry/internal/feature/tenantuser/repository"
	"github.com/umardev500/laundry/internal/feature/tenantuser/service"
)

// ProviderSet wires up the TenantUser module dependencies using Google Wire.
var ProviderSet = wire.NewSet(
	handler.NewTenantUserHandler,
	service.NewService,
	repository.NewRepository,
	NewRoutes,
)
