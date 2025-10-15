package repository

import (
	"github.com/google/uuid"

	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/service"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/service/domain"
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

func (r *entImpl) AreItemsAvailable(ctx *appctx.Context, ids []uuid.UUID) (*domain.AvailabilityResult, error) {
	conn := r.client.GetConn(ctx)
	services, err := conn.Service.
		Query().
		Where(service.IDIn(ids...)).
		All(ctx)

	if err != nil {
		return nil, err
	}

	return &domain.AvailabilityResult{
		RequestedIDs:      ids,
		AvailableServices: mapper.FromEntList(services),
	}, nil
}

// Create inserts a new service record.
func (r *entImpl) Create(ctx *appctx.Context, s *domain.Service) (*domain.Service, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.Service.
		Create().
		SetTenantID(s.TenantID).
		SetNillableServiceUnitID(s.ServiceUnitID).
		SetNillableServiceCategoryID(s.ServiceCategoryID).
		SetName(s.Name).
		SetBasePrice(s.BasePrice).
		SetNillableDescription(&s.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByID retrieves a service by its UUID.
func (r *entImpl) FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.Service, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.Service.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByName retrieves a service by name within tenant scope.
func (r *entImpl) FindByName(ctx *appctx.Context, name string) (*domain.Service, error) {
	var err error
	conn := r.client.GetConn(ctx)

	qb := conn.Service.Query().Where(service.NameEQ(name))

	qb, err = r.applyScope(ctx, qb)
	if err != nil {
		return nil, err
	}

	entModel, err := qb.Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Update modifies an existing service record.
func (r *entImpl) Update(ctx *appctx.Context, s *domain.Service) (*domain.Service, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.Service.
		UpdateOneID(s.ID).
		SetName(s.Name).
		SetBasePrice(s.BasePrice).
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
func (r *entImpl) List(ctx *appctx.Context, q *query.ListServiceQuery) (*pagination.PageData[domain.Service], error) {
	var err error

	q.Normalize()

	conn := r.client.GetConn(ctx)
	qb := conn.Service.Query()

	qb, err = r.applyScope(ctx, qb)
	if err != nil {
		return nil, err
	}

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
		qb = qb.Order(ent.Asc(service.FieldBasePrice))
	case query.OrderPriceDesc:
		qb = qb.Order(ent.Desc(service.FieldBasePrice))
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

	return &pagination.PageData[domain.Service]{
		Data:  items,
		Total: total,
	}, nil
}

// -------------------------
// Helpers
// -------------------------

// applyScope ensures tenant-level filtering.
func (r *entImpl) applyScope(ctx *appctx.Context, qb *ent.ServiceQuery) (*ent.ServiceQuery, error) {
	switch ctx.Scope() {
	case appctx.ScopeTenant:
		qb = qb.Where(service.TenantIDEQ(*ctx.TenantID()))
	case appctx.ScopeUser:
		// RBAC already handles access, no filtering needed here
	case appctx.ScopeAdmin:
	// no filtering for admin
	default:
		// Unknown scope, deny access by default
		return nil, domain.ErrUnauthorizedServiceAccess
	}

	return qb, nil
}
