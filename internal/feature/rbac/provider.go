package rbac

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/feature/rbac/handler"
	"github.com/umardev500/laundry/internal/feature/rbac/repository"
	"github.com/umardev500/laundry/internal/feature/rbac/service"
)

// ProviderSet wires up the dependencies for the RBAC (Role) feature.
var ProviderSet = wire.NewSet(
	NewRoutes,                   // Routes
	handler.NewHandler,          // HTTP handler
	service.NewService,          // Business logic
	repository.NewEntRepository, // Repository (Ent)

	handler.NewFeatureHandler,
	service.NewFeatureService,
	repository.NewFeatureEntRepository,

	handler.NewPermissionHandler,
	service.NewPermissionService,
	repository.NewPermissionRepository,
)
