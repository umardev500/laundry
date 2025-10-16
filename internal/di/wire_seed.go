//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/umardev500/laundry/internal/config"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/internal/infra/database/seeder"

	machineSeeder "github.com/umardev500/laundry/internal/feature/machine/seeder"
	machineTypeSeeder "github.com/umardev500/laundry/internal/feature/machinetype/seeder"
	orderSeeder "github.com/umardev500/laundry/internal/feature/order/seeder"
	orderStatusHistorySeeder "github.com/umardev500/laundry/internal/feature/orderstatushistory/seeder"
	paymentSeeder "github.com/umardev500/laundry/internal/feature/payment/seeder"
	paymentMethodSeeder "github.com/umardev500/laundry/internal/feature/paymentmethod/seeder"
	planSeeder "github.com/umardev500/laundry/internal/feature/plan/seeder"
	platformUserSeeder "github.com/umardev500/laundry/internal/feature/platformuser/seeder"
	rbacSeeder "github.com/umardev500/laundry/internal/feature/rbac/seeder"
	serviceSeeder "github.com/umardev500/laundry/internal/feature/service/seeder"
	serviceCategorySeeder "github.com/umardev500/laundry/internal/feature/servicecategory/seeder"
	serviceUnitSeeder "github.com/umardev500/laundry/internal/feature/serviceunit/seeder"
	tenantSeeder "github.com/umardev500/laundry/internal/feature/tenant/seeder"
	tenantUserSeeder "github.com/umardev500/laundry/internal/feature/tenantuser/seeder"
	userSeeder "github.com/umardev500/laundry/internal/feature/user/seeder"
)

func NewSeederSet(
	rbac []rbacSeeder.RBACSeeder,
	user []userSeeder.UserSeederSet,
	tenant []tenantSeeder.TenantSeederSet,
	platformUser []platformUserSeeder.PlatformUserSeederSet,
	tenantUser []tenantUserSeeder.TenantUserSeederSet,
	machine []machineSeeder.MachineSeederSet,
	machineType []machineTypeSeeder.MachineTypeSeederSet,
	serviceUnit []serviceUnitSeeder.ServiceUnitSeederSet,
	service []serviceSeeder.ServiceSeederSet,
	serviceCategory []serviceCategorySeeder.ServiceCategorySeederSet,
	orderSeeder []orderSeeder.OrderSeederSet,
	payment []paymentSeeder.PaymentSeederSet,
	paymentMethod []paymentMethodSeeder.PaymentMethodSeederSet,
	orderStatusHistory []orderStatusHistorySeeder.OrderStatusHistorySeederSet,
	plan []planSeeder.PlanSeederSet,
) []seeder.Seeder {
	var all []seeder.Seeder

	// append all tenant seeder
	for _, s := range tenant {
		all = append(all, s)
	}

	// append all rbac seeder
	for _, s := range rbac {
		all = append(all, s)
	}

	// append all user seeder
	for _, s := range user {
		all = append(all, s)
	}

	// append all platform user seeder
	for _, s := range platformUser {
		all = append(all, s)
	}

	// append all tenant user seeder
	for _, s := range tenantUser {
		all = append(all, s)
	}

	// append all machine type seeder
	for _, s := range machineType {
		all = append(all, s)
	}

	// append all machine seeder
	for _, s := range machine {
		all = append(all, s)
	}

	// append all service unit seeder
	for _, s := range serviceUnit {
		all = append(all, s)
	}

	// append all service category seeder
	for _, s := range serviceCategory {
		all = append(all, s)
	}

	// append all service seeder
	for _, s := range service {
		all = append(all, s)
	}

	// append all order seeder
	for _, s := range orderSeeder {
		all = append(all, s)
	}

	// append all payment method seeder
	for _, s := range paymentMethod {
		all = append(all, s)
	}

	// append all payment seeder
	for _, s := range payment {
		all = append(all, s)
	}

	// append all order status history seeder
	for _, s := range orderStatusHistory {
		all = append(all, s)
	}

	// append all plan seeder
	for _, s := range plan {
		all = append(all, s)
	}

	return all
}

func InitialzeSeeder(cfg *config.Config) ([]seeder.Seeder, error) {
	wire.Build(
		entdb.NewEntClient,
		userSeeder.ProviderSet,
		rbacSeeder.ProviderSet,
		tenantSeeder.ProviderSet,
		platformUserSeeder.ProviderSet,
		tenantUserSeeder.ProviderSet,
		machineSeeder.ProviderSet,
		machineTypeSeeder.ProviderSet,
		serviceUnitSeeder.ProviderSet,
		serviceSeeder.ProviderSet,
		serviceCategorySeeder.ProviderSet,
		orderSeeder.ProviderSet,
		paymentSeeder.ProviderSet,
		paymentMethodSeeder.ProviderSet,
		orderStatusHistorySeeder.ProviderSet,
		planSeeder.ProviderSet,
		NewSeederSet,
	)

	return nil, nil
}
