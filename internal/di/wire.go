//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	fiberApp "github.com/umardev500/laundry/internal/app/fiber"
	"github.com/umardev500/laundry/internal/app/router"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/feature/auth"
	"github.com/umardev500/laundry/internal/feature/machine"
	"github.com/umardev500/laundry/internal/feature/machinetype"
	"github.com/umardev500/laundry/internal/feature/order"
	"github.com/umardev500/laundry/internal/feature/orderitem"
	"github.com/umardev500/laundry/internal/feature/payment"
	"github.com/umardev500/laundry/internal/feature/paymentmethod"
	"github.com/umardev500/laundry/internal/feature/platformuser"
	"github.com/umardev500/laundry/internal/feature/rbac"
	"github.com/umardev500/laundry/internal/feature/service"
	"github.com/umardev500/laundry/internal/feature/servicecategory"
	"github.com/umardev500/laundry/internal/feature/serviceunit"
	"github.com/umardev500/laundry/internal/feature/tenant"
	"github.com/umardev500/laundry/internal/feature/tenantuser"
	"github.com/umardev500/laundry/internal/feature/user"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/redis"
	"github.com/umardev500/laundry/pkg/validator"
)

var AppSet = wire.NewSet(
	entdb.NewEntClient,
	fiberApp.NewFiberApp,
	user.ProviderSet,
	platformuser.ProviderSet,
	auth.ProviderSet,
	tenant.ProviderSet,
	router.NewRouter,
	tenantuser.ProviderSet,
	rbac.ProviderSet,
	redis.NewRedisClient,
	machine.ProviderSet,
	machinetype.ProviderSet,
	serviceunit.ProviderSet,
	service.ProviderSet,
	servicecategory.ProviderSet,
	orderitem.ProviderSet,
	payment.ProviderSet,
	paymentmethod.ProviderSet,
	order.ProviderSet,
	validator.New,
	newRegistrars,
)

func Initialize(cfg *config.Config) (*router.Router, error) {
	wire.Build(AppSet)
	return nil, nil
}

func newRegistrars(
	userReg *user.Routes,
	platformUserReg *platformuser.Routes,
	authReg *auth.Routes,
	tenantReg *tenant.Routes,
	rbacReg *rbac.Routes,
	tenantUserReg *tenantuser.Routes,
	machineReg *machine.Routes,
	machineTypeReg *machinetype.Routes,
	serviceUnitReg *serviceunit.Routes,
	serviceReg *service.Routes,
	serviceCategoryReg *servicecategory.Routes,
	orderReg *order.Routes,
	paymentMethodReg *paymentmethod.Routes,
) []router.RouteRegistrar {
	return []router.RouteRegistrar{
		userReg,
		platformUserReg,
		authReg,
		tenantReg,
		rbacReg,
		tenantReg,
		tenantUserReg,
		machineReg,
		machineTypeReg,
		serviceUnitReg,
		serviceReg,
		serviceCategoryReg,
		orderReg,
		paymentMethodReg,
	}
}
