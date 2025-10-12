package paymentmethod

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/handler"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/repository"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/service"
)

var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewService,
	repository.NewEntRepository,
)
