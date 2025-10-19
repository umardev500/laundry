package address

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/address/handler"
	"github.com/umardev500/laundry/internal/feature/address/repository"
	"github.com/umardev500/laundry/internal/feature/address/service"
)

// ProviderSet wires the dependencies for the Address feature.
var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewAddressService,
	repository.NewEntRepository,
)
