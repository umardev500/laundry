package serviceunit

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/serviceunit/handler"
	"github.com/umardev500/laundry/internal/feature/serviceunit/repository"
	"github.com/umardev500/laundry/internal/feature/serviceunit/service"
)

var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewService,
	repository.NewEntRepository,
)
