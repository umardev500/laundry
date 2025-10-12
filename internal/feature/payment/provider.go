package payment

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/payment/repository"
	"github.com/umardev500/laundry/internal/feature/payment/service"
)

// ProviderSet wires Payment module dependencies
var ProviderSet = wire.NewSet(
	repository.NewEntPaymentRepository,
	service.NewPaymentService,
)
