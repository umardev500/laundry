package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/serviceunit/domain"
	"github.com/umardev500/laundry/internal/feature/serviceunit/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Service interface {
	// Create a new service unit
	Create(ctx *appctx.Context, s *domain.ServiceUnit) (*domain.ServiceUnit, error)

	// List service units with pagination and filtering
	List(ctx *appctx.Context, q *query.ListServiceUnitQuery) (*pagination.PageData[domain.ServiceUnit], error)

	// GetByID fetches a single service unit by its ID
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.ServiceUnit, error)

	// Update modifies an existing service unit
	Update(ctx *appctx.Context, s *domain.ServiceUnit) (*domain.ServiceUnit, error)

	// Delete removes a service unit by ID
	Delete(ctx *appctx.Context, id uuid.UUID) error
}
