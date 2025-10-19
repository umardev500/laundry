package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/address/domain"
	"github.com/umardev500/laundry/internal/feature/address/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Address defines the business service interface for managing Addresses.
// This interface mirrors repository operations but allows domain-level
// validations, rules, and orchestration (e.g., ensuring only one primary address per user).
type Address interface {
	// CRUD operations
	Create(ctx *appctx.Context, a *domain.Address) (*domain.Address, error)
	List(ctx *appctx.Context, q *query.ListAddressQuery) (*pagination.PageData[domain.Address], error)
	GetByID(ctx *appctx.Context, id uuid.UUID, q *query.FindAddressByIDQuery) (*domain.Address, error)
	Update(ctx *appctx.Context, a *domain.Address) (*domain.Address, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error // soft delete
	Purge(ctx *appctx.Context, id uuid.UUID) error  // hard delete (permanent)

	// Custom business operations
	GetPrimaryByUserID(ctx *appctx.Context, userID uuid.UUID) (*domain.Address, error)
	SetPrimary(ctx *appctx.Context, id uuid.UUID, userID uuid.UUID) (*domain.Address, error)
	Restore(ctx *appctx.Context, id uuid.UUID) (*domain.Address, error)
}
