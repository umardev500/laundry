package service

import (
	"context"
	"time"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/contract"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/order/query"
	"github.com/umardev500/laundry/internal/feature/order/repository"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"

	orderItemContract "github.com/umardev500/laundry/internal/feature/orderitem/contract"
	paymentContract "github.com/umardev500/laundry/internal/feature/payment/contract"
	paymentDomain "github.com/umardev500/laundry/internal/feature/payment/domain"
	paymentMethodContract "github.com/umardev500/laundry/internal/feature/paymentmethod/contract"
	serviceContract "github.com/umardev500/laundry/internal/feature/service/contract"
)

// orderService implements OrderService interface.
type orderService struct {
	repo                 repository.Repository
	service              serviceContract.Service
	orderItemService     orderItemContract.Service
	client               *entdb.Client
	paymentService       paymentContract.Service
	paymentMethodService paymentMethodContract.Service
}

// CreatePayment implements contract.OrderService.
func (s *orderService) CreatePayment(ctx *appctx.Context, o *domain.Order) (*paymentDomain.Payment, error) {
	pymnt := o.Payment
	receivedAmount := pymnt.ReceivedAmount
	pm := pymnt.PaymentMethodID

	// Get payment method
	m, err := s.paymentMethodService.GetByID(ctx, pm)
	if err != nil {
		return nil, err
	}

	payment := o.Payment
	payment.TenantID = ctx.TenantID()
	payment.RefID = o.ID
	payment.RefType = types.PaymentTypeOrder
	payment.Amount = o.TotalAmount
	payment.Status = types.PaymentStatusPending

	if m.Type == types.PaymentMethodCash {
		if receivedAmount == nil || *receivedAmount < o.TotalAmount {
			return nil, paymentDomain.ErrInsufficientPayment
		}

		now := time.Now()
		payment.Status = types.PaymentStatusPaid
		change := *receivedAmount - o.TotalAmount
		payment.ChangeAmount = &change
		payment.PaidAt = &now
	}

	if err := payment.Validate(); err != nil {
		return nil, err
	}

	payment, err = s.paymentService.Create(ctx, payment)
	if err != nil {
		return nil, err
	}

	return payment, nil
}

// NewOrderService creates a new OrderService.
func NewOrderService(
	repo repository.Repository,
	service serviceContract.Service,
	orderItemService orderItemContract.Service,
	client *entdb.Client,
	paymentService paymentContract.Service,
	paymentMethodService paymentMethodContract.Service,
) contract.OrderService {
	return &orderService{
		repo:                 repo,
		service:              service,
		orderItemService:     orderItemService,
		client:               client,
		paymentService:       paymentService,
		paymentMethodService: paymentMethodService,
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

		result.Payment = o.Payment // Assign the payment details from the input order
		p, err := s.CreatePayment(newCtx, result)
		if err != nil {
			return err
		}

		// Assign the payment details from the created payment
		result.Payment = p

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
