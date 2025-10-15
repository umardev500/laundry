package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/machine"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/machine/domain"
	"github.com/umardev500/laundry/internal/feature/machine/mapper"
	"github.com/umardev500/laundry/internal/feature/machine/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

type entImpl struct {
	client *entdb.Client
}

func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{client: client}
}

func (e *entImpl) Create(ctx *appctx.Context, m *domain.Machine) (*domain.Machine, error) {
	conn := e.client.GetConn(ctx)
	entModel, err := conn.Machine.
		Create().
		SetTenantID(m.TenantID).
		SetNillableMachineTypeID(m.MachineTypeID).
		SetName(m.Name).
		SetDescription(m.Description).
		SetStatus(machine.Status(m.Status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entModel), nil
}

func (e *entImpl) FindById(ctx *appctx.Context, id uuid.UUID) (*domain.Machine, error) {
	conn := e.client.GetConn(ctx)
	entModel, err := conn.Machine.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entModel), nil
}

func (e *entImpl) FindByName(ctx *appctx.Context, name string) (*domain.Machine, error) {
	var err error
	conn := e.client.GetConn(ctx)
	qb := conn.Machine.Query().Where(machine.NameEQ(name))

	qb, err = e.applyScope(ctx, qb)
	if err != nil {
		return nil, err
	}

	entModel, err := qb.Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entModel), nil
}

func (e *entImpl) Update(ctx *appctx.Context, m *domain.Machine) (*domain.Machine, error) {
	conn := e.client.GetConn(ctx)
	entModel, err := conn.Machine.
		UpdateOneID(m.ID).
		SetName(m.Name).
		SetDescription(m.Description).
		SetStatus(machine.Status(m.Status)).
		SetNillableDeletedAt(m.DeletedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entModel), nil
}

func (e *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := e.client.GetConn(ctx)
	return conn.Machine.DeleteOneID(id).Exec(ctx)
}

func (e *entImpl) List(ctx *appctx.Context, q *query.ListMachineQuery) (*pagination.PageData[domain.Machine], error) {
	var err error
	q.Normalize()

	conn := e.client.GetConn(ctx)
	qb := conn.Machine.Query()

	qb, err = e.applyScope(ctx, qb)
	if err != nil {
		return nil, err
	}

	if q.Search != "" {
		qb = qb.Where(
			machine.Or(
				machine.NameContainsFold(q.Search),
				machine.DescriptionContainsFold(q.Search),
			),
		)
	}

	if q.Status != "" {
		qb = qb.Where(machine.StatusEQ(machine.Status(q.Status)))
	}

	if !q.IncludeDeleted {
		qb = qb.Where(machine.DeletedAtIsNil())
	}

	switch q.Order {
	case query.OrderNameAsc:
		qb = qb.Order(ent.Asc(machine.FieldName))
	case query.OrderNameDesc:
		qb = qb.Order(ent.Desc(machine.FieldName))
	case query.OrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(machine.FieldCreatedAt))
	case query.OrderCreatedAtDesc:
		qb = qb.Order(ent.Desc(machine.FieldCreatedAt))
	case query.OrderUpdatedAtAsc:
		qb = qb.Order(ent.Asc(machine.FieldUpdatedAt))
	case query.OrderUpdatedAtDesc:
		qb = qb.Order(ent.Desc(machine.FieldUpdatedAt))
	default:
		qb = qb.Order(ent.Desc(machine.FieldCreatedAt))
	}

	total, err := qb.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	ents, err := qb.Limit(q.Limit).Offset(q.Offset()).All(ctx)
	if err != nil {
		return nil, err
	}

	machines := mapper.FromEntModels(ents)
	return &pagination.PageData[domain.Machine]{
		Data:  machines,
		Total: total,
	}, nil
}

// -------------------------
// Helpers
// -------------------------

// applyScope ensures tenant-level filtering.
func (r *entImpl) applyScope(ctx *appctx.Context, qb *ent.MachineQuery) (*ent.MachineQuery, error) {
	switch ctx.Scope() {
	case appctx.ScopeTenant:
		qb = qb.Where(machine.TenantIDEQ(*ctx.TenantID()))
	case appctx.ScopeUser:
		// RBAC already handles access, no filtering needed here
	case appctx.ScopeAdmin:
	// no filtering for admin
	default:
		// Unknown scope, deny access by default
		return nil, domain.ErrUnauthorizedMachineAccess
	}

	return qb, nil
}
