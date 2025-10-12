package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/domain"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Service defines the PaymentMethod service interface
type Service interface {
	Create(ctx *appctx.Context, pm *domain.PaymentMethod) (*domain.PaymentMethod, error)
	Update(ctx *appctx.Context, pm *domain.PaymentMethod) (*domain.PaymentMethod, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	Purge(ctx *appctx.Context, id uuid.UUID) error
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.PaymentMethod, error)
	GetByName(ctx *appctx.Context, name string) (*domain.PaymentMethod, error)
	List(ctx *appctx.Context, q *query.ListPaymentMethodQuery) (*pagination.PageData[domain.PaymentMethod], error)
}
