package contract

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/machine/domain"
	"github.com/umardev500/laundry/internal/feature/machine/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Service interface {
	Create(ctx *appctx.Context, m *domain.Machine) (*domain.Machine, error)
	List(ctx *appctx.Context, q *query.ListMachineQuery) (*pagination.PageData[domain.Machine], error)
	GetByID(ctx *appctx.Context, id uuid.UUID) (*domain.Machine, error)
	GetByName(ctx *appctx.Context, name string) (*domain.Machine, error)
	Update(ctx *appctx.Context, m *domain.Machine) (*domain.Machine, error)
	UpdateStatus(ctx *appctx.Context, m *domain.Machine) (*domain.Machine, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	Purge(ctx *appctx.Context, id uuid.UUID) error
}
