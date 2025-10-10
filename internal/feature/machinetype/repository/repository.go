package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/machinetype/domain"
	"github.com/umardev500/laundry/internal/feature/machinetype/query"
	"github.com/umardev500/laundry/pkg/pagination"
)

type Repository interface {
	Create(ctx *appctx.Context, t *domain.MachineType) (*domain.MachineType, error)
	FindById(ctx *appctx.Context, id uuid.UUID) (*domain.MachineType, error)
	FindByName(ctx *appctx.Context, name string) (*domain.MachineType, error)
	Update(ctx *appctx.Context, t *domain.MachineType) (*domain.MachineType, error)
	Delete(ctx *appctx.Context, id uuid.UUID) error
	List(ctx *appctx.Context, q *query.ListMachineTypeQuery) (*pagination.PageData[domain.MachineType], error)
}
