package plan

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/plan/handler"
	"github.com/umardev500/laundry/internal/feature/plan/repository"
	"github.com/umardev500/laundry/internal/feature/plan/service"
)

var ProviderSet = wire.NewSet(
	NewRoutes,
	handler.NewHandler,
	service.NewPlanService,
	repository.NewEntRepository,
)
