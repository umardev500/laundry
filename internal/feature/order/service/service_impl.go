package service

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
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

// Preview implements contract.OrderService.
// It calculates the total amount and details for an order before payment or creation.
func (s *orderService) Preview(ctx *appctx.Context, o *domain.Order) (*domain.Order, error) {
	// 1️⃣ Get service availability
	serviceIDs := o.GetServiceIDs()
	availability, err := s.service.AreItemsAvailable(ctx, serviceIDs)
	if err != nil {
		return nil, err
	}

	// 2️⃣ If some services are not available, return which ones
	if !availability.AllAvailable() {
		return nil, domain.NewServiceUnavailableError(availability.UnavailableIDs())
	}

	// 3️⃣ Place the order temporarily (calculate totals but don’t persist)
	if err := o.Place(availability.AvailableServices); err != nil {
		return nil, err
	}

	// 4️⃣ You can set a default status for the preview
	o.Status = types.OrderStatusPreview

	return o, nil
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

	// Validate
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

// FindByID implements contract.OrderService.
func (s *orderService) FindByID(ctx *appctx.Context, id uuid.UUID, q *query.OrderQuery) (*domain.Order, error) {
	return s.findExisting(ctx, id, q)
}

// UpdateStatus implements contract.OrderService.
func (s *orderService) UpdateStatus(ctx *appctx.Context, o *domain.Order) (*domain.Order, error) {
	order, err := s.findExisting(ctx, o.ID, nil)
	if err != nil {
		return nil, err
	}

	order.UpdateStatus(o.Status)

	return s.repo.Update(ctx, order)
}

// -----------------------
// Helper methods
// -----------------------

// findExisting ensures the payment exists, is not soft-deleted, and belongs to tenant
func (s *orderService) findExisting(ctx *appctx.Context, id uuid.UUID, q *query.OrderQuery) (*domain.Order, error) {
	if q == nil {
		q = &query.OrderQuery{}
	}

	p, err := s.repo.FindById(ctx, id, q)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}

	if p.IsDeleted() {
		return nil, domain.ErrOrderDeleted
	}

	if !p.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedOrderAccess
	}

	return p, nil
}

// findAllowDeleted fetches a payment regardless of deleted status but checks tenant ownership
func (s *orderService) findAllowDeleted(ctx *appctx.Context, id uuid.UUID) (*domain.Order, error) {
	p, err := s.repo.FindById(ctx, id, nil)
	if err != nil {
		if !ent.IsNotFound(err) {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}

	if !p.BelongsToTenant(ctx) {
		return nil, domain.ErrUnauthorizedOrderAccess
	}

	return p, nil
}
