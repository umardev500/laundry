package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/address/domain"
	"github.com/umardev500/laundry/internal/feature/address/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Repository defines the interface for Address storage operations.
type Repository interface {
	// Create inserts a new address record.
	Create(ctx *appctx.Context, a *domain.Address) (*domain.Address, error)

	// FindByID retrieves an address by its ID.
	FindByID(ctx *appctx.Context, id uuid.UUID, q *query.FindAddressByIDQuery) (*domain.Address, error)

	// FindPrimaryByUserID retrieves the primary address for a given user.
	FindPrimaryByUserID(ctx *appctx.Context, userID uuid.UUID) (*domain.Address, error)

	// UnsetPrimary marks all addresses of a user as non-primary.
	UnsetPrimary(ctx *appctx.Context, userID uuid.UUID) error

	// Update modifies an existing address record.
	Update(ctx *appctx.Context, a *domain.Address) (*domain.Address, error)

	// Delete performs a soft or hard delete of an address by ID.
	Delete(ctx *appctx.Context, id uuid.UUID) error

	// List returns a paginated list of addresses filtered by query parameters.
	List(ctx *appctx.Context, q *query.ListAddressQuery) (*pagination.PageData[domain.Address], error)
}
