package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/user/domain"
	"github.com/umardev500/laundry/internal/feature/user/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Service interface {
	// Create a new user
	Create(ctx *appctx.Context, user *domain.User) (*domain.User, error)

	// List users with pagination and filters
	List(ctx *appctx.Context, query *query.ListUserQuery) (*pagination.PageData[domain.User], error)

	// Get a single users by ID
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.User, error)

	// Get a single user by email
	GetByEmail(ctx *appctx.Context, email string) (*domain.User, error)

	// Update an existing user
	Update(ctx *appctx.Context, user *domain.User) (*domain.User, error)

	// UpdateStatus updates the status of a user
	UpdateStatus(ctx *appctx.Context, user *domain.User) (*domain.User, error)

	// Soft delete(or hard delete, depending on repo) a user by ID
	Delete(ctx *appctx.Context, id uuid.UUID) error

	// Purge is a hard delete of a user by ID
	Purge(ctx *appctx.Context, id uuid.UUID) error
}
