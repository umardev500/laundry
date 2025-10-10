package repository

import (
	"github.com/google/uuid"

	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/service"
	"github.com/umardev500/laundry/internal/app/appctx"
	domainPkg "github.com/umardev500/laundry/internal/feature/service/domain"
	"github.com/umardev500/laundry/internal/feature/service/mapper"
	"github.com/umardev500/laundry/internal/feature/service/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

// entImpl implements the Service repository using Ent.
type entImpl struct {
	client *entdb.Client
}

// NewEntRepository returns a new Ent-based repository for Service.
func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{
		client: client,
	}
}

// Create inserts a new service record.
func (r *entImpl) Create(ctx *appctx.Context, s *domainPkg.Service) (*domainPkg.Service, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.Service.
		Create().
		SetTenantID(s.TenantID).
		SetNillableServiceUnitID(s.ServiceUnitID).
		SetNillableServiceCategoryID(s.ServiceCategoryID).
		SetName(s.Name).
		SetPrice(s.Price).
		SetNillableDescription(&s.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByID retrieves a service by its UUID.
func (r *entImpl) FindByID(ctx *appctx.Context, id uuid.UUID) (*domainPkg.Service, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.Service.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByName retrieves a service by name within tenant scope.
func (r *entImpl) FindByName(ctx *appctx.Context, name string) (*domainPkg.Service, error) {
	conn := r.client.GetConn(ctx)

	qb := conn.Service.Query().Where(service.NameEQ(name))
	qb = r.tenantScopedQuery(ctx, qb)

	entModel, err := qb.Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Update modifies an existing service record.
func (r *entImpl) Update(ctx *appctx.Context, s *domainPkg.Service) (*domainPkg.Service, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.Service.
		UpdateOneID(s.ID).
		SetName(s.Name).
		SetPrice(s.Price).
		SetNillableDescription(&s.Description).
		SetNillableServiceUnitID(s.ServiceUnitID).
		SetNillableServiceCategoryID(s.ServiceCategoryID).
		SetNillableDeletedAt(s.DeletedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Delete removes a service by ID (hard delete).
func (r *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)
	return conn.Service.DeleteOneID(id).Exec(ctx)
}

// List retrieves paginated services with filtering, ordering and tenant scoping.
func (r *entImpl) List(ctx *appctx.Context, q *query.ListServiceQuery) (*pagination.PageData[domainPkg.Service], error) {
	q.Normalize()

	conn := r.client.GetConn(ctx)
	qb := conn.Service.Query()
	qb = r.tenantScopedQuery(ctx, qb)

	if q.Search != "" {
		qb = qb.Where(
			service.Or(
				service.NameContainsFold(q.Search),
				service.DescriptionContainsFold(q.Search),
			),
		)
	}

	if !q.IncludeDeleted {
		qb = qb.Where(service.DeletedAtIsNil())
	}

	switch q.Order {
	case query.OrderNameAsc:
		qb = qb.Order(ent.Asc(service.FieldName))
	case query.OrderNameDesc:
		qb = qb.Order(ent.Desc(service.FieldName))
	case query.OrderPriceAsc:
		qb = qb.Order(ent.Asc(service.FieldPrice))
	case query.OrderPriceDesc:
		qb = qb.Order(ent.Desc(service.FieldPrice))
	case query.OrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(service.FieldCreatedAt))
	default:
		qb = qb.Order(ent.Desc(service.FieldCreatedAt))
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

	return &pagination.PageData[domainPkg.Service]{
		Data:  items,
		Total: total,
	}, nil
}

// tenantScopedQuery ensures queries are filtered by tenant ID when available.
func (r *entImpl) tenantScopedQuery(ctx *appctx.Context, qb *ent.ServiceQuery) *ent.ServiceQuery {
	tenantID := ctx.TenantID()
	if tenantID == nil {
		return qb // fallback for system-wide/migrations/seeders
	}
	return qb.Where(service.TenantIDEQ(*tenantID))
}
