package orderitem

import (
	"github.com/google/wire"

	"github.com/umardev500/laundry/internal/feature/orderitem/repository"
	service "github.com/umardev500/laundry/internal/feature/orderitem/service"
)

var ProviderSet = wire.NewSet(
	service.New,
	repository.NewEntRepository,
)
