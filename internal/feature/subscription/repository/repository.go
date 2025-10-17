package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/subscription/domain"
	"github.com/umardev500/laundry/internal/feature/subscription/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Repository defines the interface for Subscription storage operations.
type Repository interface {
	Create(ctx *appctx.Context, s *domain.Subscription) (*domain.Subscription, error)
	FindByID(ctx *appctx.Context, id uuid.UUID, q *query.FindSubscriptionByIDQuery) (*domain.Subscription, error)
	Update(ctx *appctx.Context, s *domain.Subscription) (*domain.Subscription, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListSubscriptionQuery) (*pagination.PageData[domain.Subscription], error)
}
