package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/payment/domain"
	"github.com/umardev500/laundry/internal/feature/payment/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Service defines the business logic for payments.
type Service interface {
	// Create a new payment
	Create(ctx *appctx.Context, payment *domain.Payment) (*domain.Payment, error)

	// Update an existing payment
	Update(ctx *appctx.Context, payment *domain.Payment) (*domain.Payment, error)

	// Updte status
	UpdateStatus(ctx *appctx.Context, payment *domain.Payment) (*domain.Payment, error)

	// GetByID retrieves a payment by its ID
	GetByID(ctx *appctx.Context, id uuid.UUID, q *query.FindPaymentByIdQuery) (*domain.Payment, error)

	// Delete a payment by its ID
	Delete(ctx *appctx.Context, id uuid.UUID) error

	// List retrieves paginated payments with filters
	List(ctx *appctx.Context, q *query.ListPaymentQuery) (*pagination.PageData[domain.Payment], error)

	// MarkPaid completes a payment (cash or other methods)
	MarkPaid(ctx *appctx.Context, id uuid.UUID, receivedAmount float64) (*domain.Payment, error)

	// Purge a payment by its ID
	Purge(ctx *appctx.Context, id uuid.UUID) error
}
