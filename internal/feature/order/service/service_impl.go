package service

import (
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/contract"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/order/query"
	"github.com/umardev500/laundry/internal/feature/order/repository"
	"github.com/umardev500/laundry/pkg/pagination"
)

// orderService implements OrderService interface.
type orderService struct {
	repo repository.Repository
}

// NewOrderService creates a new OrderService.
func NewOrderService(repo repository.Repository) contract.OrderService {
	return &orderService{
		repo: repo,
	}
}

// List returns a paginated list of orders.
func (s *orderService) List(ctx *appctx.Context, q *query.ListOrderQuery) (*pagination.PageData[domain.Order], error) {
	q.Normalize()

	// The repository already handles scope filtering (tenant/user/admin)
	return s.repo.List(ctx, q)
}
