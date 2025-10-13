package orchestrator

import (
	"context"
	"fmt"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/payment/contract"
	"github.com/umardev500/laundry/internal/feature/payment/domain"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/types"

	orderServiceContract "github.com/umardev500/laundry/internal/feature/order/contract"
	orderDomain "github.com/umardev500/laundry/internal/feature/order/domain"
)

type paymentToOrder struct {
	orderService orderServiceContract.OrderService
	service      contract.Service
	client       *entdb.Client
}

func NewPaymentToOrder(orderService orderServiceContract.OrderService, service contract.Service, client *entdb.Client) contract.Orchestrator {
	return &paymentToOrder{
		orderService: orderService,
		service:      service,
		client:       client,
	}
}

func (p *paymentToOrder) SyncOrder(ctx *appctx.Context, ord *orderDomain.Order, pay *domain.Payment) error {
	if ord == nil || pay == nil {
		return fmt.Errorf("order or payment is nil")
	}

	var newStatus types.OrderStatus

	switch pay.Status {
	case types.PaymentStatusPending:
		newStatus = types.OrderStatusPending
	case types.PaymentStatusPaid:
		newStatus = types.OrderStatusConfirmed
	case types.PaymentStatusFailed:
		newStatus = types.OrderStatusCancelled
	default:
		return fmt.Errorf("unsupported payment status: %s", pay.Status)
	}

	err := p.client.WithTransaction(ctx, func(ctx context.Context) error {
		newCtx := appctx.New(ctx)

		// Call OrderService to update the status
		ord.Status = newStatus
		_, err := p.orderService.UpdateStatus(newCtx, ord)
		if err != nil {
			return fmt.Errorf("failed to update order status: %w", err)
		}

		// Call PaymentService to update the payment status if needed
		_, err = p.service.UpdateStatus(newCtx, pay)
		if err != nil {
			return fmt.Errorf("failed to update payment status: %w", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	return nil
}
