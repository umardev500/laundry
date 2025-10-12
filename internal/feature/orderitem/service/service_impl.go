package service

import (
	"fmt"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/orderitem/contract"
	"github.com/umardev500/laundry/internal/feature/orderitem/domain"
	"github.com/umardev500/laundry/internal/feature/orderitem/repository"
)

type orderItemService struct {
	repo repository.Repository
}

func New(repo repository.Repository) contract.Service {
	return &orderItemService{repo: repo}
}

func (s *orderItemService) Create(ctx *appctx.Context, items []*domain.OrderItem) ([]*domain.OrderItem, error) {
	if len(items) == 0 {
		return []*domain.OrderItem{}, fmt.Errorf("no order items provided")
	}
	// 💡 You could add future business logic here:
	// - price validation
	// - stock checks
	// - tax calculation, etc.

	// Validte and calculate each item before insert
	for _, item := range items {
		if err := item.Validate(); err != nil {
			return nil, err
		}

		item.CalculateTotals()
	}

	fmt.Println(items)

	return s.repo.Create(ctx, items)
}
