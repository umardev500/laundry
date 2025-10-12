package service

import (
	"context"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/contract"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/order/query"
	"github.com/umardev500/laundry/internal/feature/order/repository"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"

	orderItemContract "github.com/umardev500/laundry/internal/feature/orderitem/contract"
	serviceContract "github.com/umardev500/laundry/internal/feature/service/contract"
)

// orderService implements OrderService interface.
type orderService struct {
	repo             repository.Repository
	service          serviceContract.Service
	orderItemService orderItemContract.Service
	client           *entdb.Client
}

// NewOrderService creates a new OrderService.
func NewOrderService(
	repo repository.Repository,
	service serviceContract.Service,
	orderItemService orderItemContract.Service,
	client *entdb.Client,
) contract.OrderService {
	return &orderService{
		repo:             repo,
		service:          service,
		orderItemService: orderItemService,
		client:           client,
	}
}

// GuestOrder implements contract.OrderService.
func (s *orderService) GuestOrder(ctx *appctx.Context, o *domain.Order) (*domain.Order, error) {
	serviceIDs := o.GetServiceIDs()
	availability, err := s.service.AreItemsAvailable(ctx, serviceIDs)
	if err != nil {
		return nil, err
	}

	if !availability.AllAvailable() {
		return nil, domain.NewServiceUnavailableError(availability.UnavailableIDs())
	}

	// Init defaults
	o.InitDefaults()

	// Validate the input
	if err := o.Validate(); err != nil {
		return nil, err
	}

	if err := o.Place(availability.AvailableServices); err != nil {
		return nil, err
	}

	var result *domain.Order
	err = s.client.WithTransaction(ctx, func(txCtx context.Context) error {
		var err error
		newCtx := appctx.New(txCtx)

		result, err = s.repo.Create(newCtx, o)
		if err != nil {
			return err
		}

		// 🔹 Attach the generated OrderID to each order item before creating them
		o.AttachOrderID(result.ID)

		createdItems, err := s.orderItemService.Create(newCtx, o.Items)
		if err != nil {
			return err
		}

		result.Items = createdItems
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// List returns a paginated list of orders.
func (s *orderService) List(ctx *appctx.Context, q *query.ListOrderQuery) (*pagination.PageData[domain.Order], error) {
	q.Normalize()

	// The repository already handles scope filtering (tenant/user/admin)
	return s.repo.List(ctx, q)
}
