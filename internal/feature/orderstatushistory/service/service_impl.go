package service

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/contract"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/query"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/repository"
	"github.com/umardev500/laundry/pkg/pagination"

	domain "github.com/umardev500/laundry/internal/feature/orderstatushistory/domain"
)

type statusHistoryService struct {
	repo repository.StatusHistoryRepository
}

// NewStatusHistoryService creates a new instance of the service
func NewStatusHistoryService(repo repository.StatusHistoryRepository) contract.StatusHistoryService {
	return &statusHistoryService{repo: repo}
}

// Create a new status history record
func (s *statusHistoryService) Create(ctx *appctx.Context, sh *domain.OrderStatusHistory) (*domain.OrderStatusHistory, error) {
	return s.repo.Create(ctx, sh)
}

// FindByID retrieves a status history by ID
func (s *statusHistoryService) FindByID(ctx *appctx.Context, id uuid.UUID, q *query.StatusHistoryByIDQuery) (*domain.OrderStatusHistory, error) {
	return s.findExisting(ctx, id, q)
}

// List retrieves status history records with optional filters
func (s *statusHistoryService) List(ctx *appctx.Context, q *query.OrderStatusHistoryListQuery) (*pagination.PageData[domain.OrderStatusHistory], error) {
	return s.repo.List(ctx, q)
}

// -------------------------
// Helpers
// -------------------------

func (s *statusHistoryService) findExisting(ctx *appctx.Context, id uuid.UUID, q *query.StatusHistoryByIDQuery) (*domain.OrderStatusHistory, error) {
	sh, err := s.repo.FindById(ctx, id, q)
	if err != nil {
		// Map repository not found error to domain error
		if sh == nil {
			return nil, domain.ErrStatusHistoryNotFound
		}
		return nil, err
	}

	return sh, nil
}
