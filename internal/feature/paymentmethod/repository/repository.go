package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/domain"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Repository defines the interface for PaymentMethod persistence
type Repository interface {
	// Create inserts a new payment method
	Create(ctx *appctx.Context, pm *domain.PaymentMethod) (*domain.PaymentMethod, error)

	// FindByID returns a payment method by its ID
	FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.PaymentMethod, error)

	// FindByName returns a payment method by its name
	FindByName(ctx *appctx.Context, name string) (*domain.PaymentMethod, error)

	// Update modifies an existing payment method
	Update(ctx *appctx.Context, pm *domain.PaymentMethod) (*domain.PaymentMethod, error)

	// Delete performs a soft delete
	Delete(ctx *appctx.Context, id uuid.UUID) error

	// List returns a paginated list of payment methods
	List(ctx *appctx.Context, q *query.ListPaymentMethodQuery) (*pagination.PageData[domain.PaymentMethod], error)
}
