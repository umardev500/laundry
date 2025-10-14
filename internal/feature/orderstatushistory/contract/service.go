package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/domain"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

// StatusHistoryService defines service methods for order status milestones.
type StatusHistoryService interface {
	Create(ctx *appctx.Context, sh *domain.OrderStatusHistory) (*domain.OrderStatusHistory, error)
	FindByID(ctx *appctx.Context, id uuid.UUID, q *query.StatusHistoryByIDQuery) (*domain.OrderStatusHistory, error)
	List(ctx *appctx.Context, q *query.OrderStatusHistoryListQuery) (*pagination.PageData[domain.OrderStatusHistory], error)
}
