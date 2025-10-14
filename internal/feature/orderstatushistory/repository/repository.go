package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/pagination"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/domain"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/query"
)

// StatusHistoryRepository defines repository methods for order status milestones.
type StatusHistoryRepository interface {
	// Create a new status history record for an order
	Create(ctx *appctx.Context, sh *domain.OrderStatusHistory) (*domain.OrderStatusHistory, error)

	// FindById retrieves a status history record by its ID
	FindById(ctx *appctx.Context, id uuid.UUID, q *query.StatusHistoryByIDQuery) (*domain.OrderStatusHistory, error)

	// List retrieves status history records with optional filters
	List(ctx *appctx.Context, q *query.OrderStatusHistoryListQuery) (*pagination.PageData[domain.OrderStatusHistory], error)
}
