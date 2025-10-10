package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/machinetype"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/machinetype/domain"
	"github.com/umardev500/laundry/internal/feature/machinetype/mapper"
	"github.com/umardev500/laundry/internal/feature/machinetype/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

type entImpl struct {
	client *entdb.Client
}

func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{client: client}
}

func (e *entImpl) Create(ctx *appctx.Context, t *domain.MachineType) (*domain.MachineType, error) {
	conn := e.client.GetConn(ctx)
	entModel, err := conn.MachineType.
		Create().
		SetTenantID(t.TenantID).
		SetName(t.Name).
		SetDescription(t.Description).
		SetCapacity(t.Capacity).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entModel), nil
}

func (e *entImpl) FindById(ctx *appctx.Context, id uuid.UUID) (*domain.MachineType, error) {
	conn := e.client.GetConn(ctx)
	entModel, err := conn.MachineType.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entModel), nil
}

func (e *entImpl) FindByName(ctx *appctx.Context, name string) (*domain.MachineType, error) {
	conn := e.client.GetConn(ctx)
	qb := conn.MachineType.Query().Where(machinetype.NameEQ(name))
	qb = e.tenantScopedQuery(ctx, qb)

	entModel, err := qb.Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entModel), nil
}

func (e *entImpl) Update(ctx *appctx.Context, t *domain.MachineType) (*domain.MachineType, error) {
	conn := e.client.GetConn(ctx)
	entModel, err := conn.MachineType.
		UpdateOneID(t.ID).
		SetNillableDescription(&t.Description).
		SetNillableCapacity(&t.Capacity).
		SetName(t.Name).
		SetNillableDeletedAt(t.DeletedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entModel), nil
}

func (e *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := e.client.GetConn(ctx)

	err := conn.MachineType.DeleteOneID(id).Exec(ctx)

	return err
}

func (e *entImpl) List(ctx *appctx.Context, q *query.ListMachineTypeQuery) (*pagination.PageData[domain.MachineType], error) {
	q.Normalize()

	conn := e.client.GetConn(ctx)
	qb := conn.MachineType.Query()
	qb = e.tenantScopedQuery(ctx, qb)

	if q.Search != "" {
		qb = qb.Where(
			machinetype.Or(
				machinetype.NameContainsFold(q.Search),
				machinetype.DescriptionContainsFold(q.Search),
			),
		)
	}

	if !q.IncludeDeleted {
		qb = qb.Where(machinetype.DeletedAtIsNil())
	}

	switch q.Order {
	case query.OrderNameAsc:
		qb = qb.Order(ent.Asc(machinetype.FieldName))
	case query.OrderNameDesc:
		qb = qb.Order(ent.Desc(machinetype.FieldName))
	case query.OrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(machinetype.FieldCreatedAt))
	case query.OrderCreatedAtDesc:
		qb = qb.Order(ent.Desc(machinetype.FieldCreatedAt))
	case query.OrderUpdatedAtAsc:
		qb = qb.Order(ent.Asc(machinetype.FieldUpdatedAt))
	case query.OrderUpdatedAtDesc:
		qb = qb.Order(ent.Desc(machinetype.FieldUpdatedAt))
	default:
		qb = qb.Order(ent.Desc(machinetype.FieldCreatedAt))
	}

	total, err := qb.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	ents, err := qb.Limit(q.Limit).Offset(q.Offset()).All(ctx)
	if err != nil {
		return nil, err
	}

	items := mapper.FromEntModels(ents)
	return &pagination.PageData[domain.MachineType]{
		Data:  items,
		Total: total,
	}, nil
}

// tenantScopedQuery filters queries by tenant id when applicable.
func (e *entImpl) tenantScopedQuery(ctx *appctx.Context, qb *ent.MachineTypeQuery) *ent.MachineTypeQuery {
	tenantID := ctx.TenantID()
	if tenantID == nil {
		return qb
	}
	return qb.Where(machinetype.TenantIDEQ(*tenantID))
}
