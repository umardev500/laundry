package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/subscription/domain"
	"github.com/umardev500/laundry/internal/feature/subscription/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Service defines the business operations for managing Subscriptions.
type Service interface {
	// Create registers a new subscription for a tenant and plan.
	Create(ctx *appctx.Context, s *domain.Subscription) (*domain.Subscription, error)

	// Delete performs a soft delete (marks subscription as deleted).
	Delete(ctx *appctx.Context, id uuid.UUID) error

	// Purge permanently removes a subscription from the database.
	Purge(ctx *appctx.Context, id uuid.UUID) error

	// Restore reactivates a deleted subscription.
	Restore(ctx *appctx.Context, id uuid.UUID) (*domain.Subscription, error)

	// FindByID retrieves a subscription by its ID.
	FindByID(ctx *appctx.Context, id uuid.UUID, q *query.FindSubscriptionByIDQuery) (*domain.Subscription, error)

	// Update modifies an existing subscription (e.g. change status, end date, plan, etc.).
	Update(ctx *appctx.Context, s *domain.Subscription) (*domain.Subscription, error)

	// UpdateStatus updates the status of a subscription.
	UpdateStatus(ctx *appctx.Context, s *domain.Subscription) (*domain.Subscription, error)

	// List returns paginated subscriptions with optional filters, sorting, and inclusion of deleted ones.
	List(ctx *appctx.Context, q *query.ListSubscriptionQuery) (*pagination.PageData[domain.Subscription], error)
}
