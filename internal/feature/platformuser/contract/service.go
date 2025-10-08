package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/platformuser/domain"
	"github.com/umardev500/laundry/internal/feature/platformuser/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Service defines the business logic for PlatformUser.
type Service interface {
	// Create a new PlatformUser for an existing user
	Create(ctx *appctx.Context, pu *domain.PlatformUser) (*domain.PlatformUser, error)

	// GetByID retrieves a PlatformUser by its ID
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.PlatformUser, error)

	// GetByUserID retrieves a PlatformUser by the linked UserID
	GetByUserID(ctx *appctx.Context, userID uuid.UUID) (*domain.PlatformUser, error)

	// List returns paginated PlatformUsers based on filters and search
	List(ctx *appctx.Context, q *query.ListPlatformUserQuery) (*pagination.PageData[domain.PlatformUser], error)

	// UpdateStatus updates only the status of a PlatformUser
	UpdateStatus(ctx *appctx.Context, pu *domain.PlatformUser) (*domain.PlatformUser, error)

	// Delete soft deletes a PlatformUser
	Delete(ctx *appctx.Context, id uuid.UUID) error

	// Purge permanently deletes a PlatformUser
	Purge(ctx *appctx.Context, id uuid.UUID) error
}
