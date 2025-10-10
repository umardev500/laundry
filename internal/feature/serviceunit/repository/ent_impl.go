package repository

import (
	"github.com/google/uuid"

	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/serviceunit"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/serviceunit/domain"
	"github.com/umardev500/laundry/internal/feature/serviceunit/mapper"
	"github.com/umardev500/laundry/internal/feature/serviceunit/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

// entImpl implements the ServiceUnit repository using Ent.
type entImpl struct {
	client *entdb.Client
}

// NewEntRepository returns a new Ent-based repository for ServiceUnit.
func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{
		client: client,
	}
}

// Create inserts a new service unit record.
func (e *entImpl) Create(ctx *appctx.Context, s *domain.ServiceUnit) (*domain.ServiceUnit, error) {
	conn := e.client.GetConn(ctx)

	entModel, err := conn.ServiceUnit.
		Create().
		SetTenantID(s.TenantID).
		SetName(s.Name).
		SetNillableSymbol(&s.Symbol).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByID retrieves a service unit by its UUID.
func (e *entImpl) FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.ServiceUnit, error) {
	conn := e.client.GetConn(ctx)

	entModel, err := conn.ServiceUnit.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByName retrieves a service unit by name within the tenant scope.
func (e *entImpl) FindByName(ctx *appctx.Context, name string) (*domain.ServiceUnit, error) {
	conn := e.client.GetConn(ctx)

	qb := conn.ServiceUnit.Query().Where(serviceunit.NameEQ(name))
	qb = e.tenantScopedQuery(ctx, qb)

	entModel, err := qb.Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Update modifies an existing service unit record.
func (e *entImpl) Update(ctx *appctx.Context, s *domain.ServiceUnit) (*domain.ServiceUnit, error) {
	conn := e.client.GetConn(ctx)

	entModel, err := conn.ServiceUnit.
		UpdateOneID(s.ID).
		SetName(s.Name).
		SetNillableSymbol(&s.Symbol).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Delete removes a service unit by ID.
func (e *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := e.client.GetConn(ctx)
	return conn.ServiceUnit.DeleteOneID(id).Exec(ctx)
}

// List retrieves paginated service units with optional filters and ordering.
func (e *entImpl) List(ctx *appctx.Context, q *query.ListServiceUnitQuery) (*pagination.PageData[domain.ServiceUnit], error) {
	q.Normalize()

	conn := e.client.GetConn(ctx)
	qb := conn.ServiceUnit.Query()
	qb = e.tenantScopedQuery(ctx, qb)

	if q.Search != "" {
		qb = qb.Where(serviceunit.NameContainsFold(q.Search))
	}

	switch q.Order {
	case query.OrderNameAsc:
		qb = qb.Order(ent.Asc(serviceunit.FieldName))
	case query.OrderNameDesc:
		qb = qb.Order(ent.Desc(serviceunit.FieldName))
	case query.OrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(serviceunit.FieldCreatedAt))
	default:
		qb = qb.Order(ent.Desc(serviceunit.FieldCreatedAt))
	}

	total, err := qb.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	ents, err := qb.
		Limit(q.Limit).
		Offset(q.Offset()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	items := mapper.FromEntList(ents)
	return &pagination.PageData[domain.ServiceUnit]{
		Data:  items,
		Total: total,
	}, nil
}

// tenantScopedQuery ensures queries are filtered by tenant ID when available.
func (e *entImpl) tenantScopedQuery(ctx *appctx.Context, qb *ent.ServiceUnitQuery) *ent.ServiceUnitQuery {
	tenantID := ctx.TenantID()
	if tenantID == nil {
		return qb // fallback for system-wide access (like seeders or platform users)
	}
	return qb.Where(serviceunit.TenantIDEQ(*tenantID))
}
