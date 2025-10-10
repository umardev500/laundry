package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/servicecategory"
	"github.com/umardev500/laundry/internal/app/appctx"
	domainPkg "github.com/umardev500/laundry/internal/feature/servicecategory/domain"
	"github.com/umardev500/laundry/internal/feature/servicecategory/mapper"
	"github.com/umardev500/laundry/internal/feature/servicecategory/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

// entImpl implements Repository using Ent.
type entImpl struct {
	client *entdb.Client
}

// NewEntRepository creates a new Ent repository.
func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{client: client}
}

// Create inserts a new ServiceCategory.
func (r *entImpl) Create(ctx *appctx.Context, s *domainPkg.ServiceCategory) (*domainPkg.ServiceCategory, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.ServiceCategory.
		Create().
		SetTenantID(s.TenantID).
		SetName(s.Name).
		SetNillableDescription(&s.Description).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByID retrieves a ServiceCategory by ID.
func (r *entImpl) FindByID(ctx *appctx.Context, id uuid.UUID) (*domainPkg.ServiceCategory, error) {
	conn := r.client.GetConn(ctx)
	entModel, err := conn.ServiceCategory.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.FromEnt(entModel), nil
}

// FindByName retrieves a ServiceCategory by name (tenant scoped).
func (r *entImpl) FindByName(ctx *appctx.Context, name string) (*domainPkg.ServiceCategory, error) {
	conn := r.client.GetConn(ctx)
	qb := conn.ServiceCategory.Query().Where(servicecategory.NameEQ(name))
	qb = r.tenantScopedQuery(ctx, qb)

	entModel, err := qb.Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEnt(entModel), nil
}

// Update modifies an existing ServiceCategory.
func (r *entImpl) Update(ctx *appctx.Context, s *domainPkg.ServiceCategory) (*domainPkg.ServiceCategory, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.ServiceCategory.
		UpdateOneID(s.ID).
		SetName(s.Name).
		SetNillableDescription(&s.Description).
		SetNillableDeletedAt(s.DeletedAt).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Delete permanently deletes a ServiceCategory.
func (r *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)
	return conn.ServiceCategory.DeleteOneID(id).Exec(ctx)
}

// List returns paginated results with filtering and ordering.
func (r *entImpl) List(ctx *appctx.Context, q *query.ListServiceCategoryQuery) (*pagination.PageData[domainPkg.ServiceCategory], error) {
	q.Normalize()

	conn := r.client.GetConn(ctx)
	qb := conn.ServiceCategory.Query()
	qb = r.tenantScopedQuery(ctx, qb)

	if q.Search != "" {
		qb = qb.Where(servicecategory.NameContainsFold(q.Search))
	}

	if !q.IncludeDeleted {
		qb = qb.Where(servicecategory.DeletedAtIsNil())
	}

	switch q.Order {
	case query.OrderNameAsc:
		qb = qb.Order(ent.Asc(servicecategory.FieldName))
	case query.OrderNameDesc:
		qb = qb.Order(ent.Desc(servicecategory.FieldName))
	case query.OrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(servicecategory.FieldCreatedAt))
	default:
		qb = qb.Order(ent.Desc(servicecategory.FieldCreatedAt))
	}

	total, err := qb.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	ents, err := qb.Limit(q.Limit).Offset(q.Offset()).All(ctx)
	if err != nil {
		return nil, err
	}

	items := mapper.FromEntList(ents)

	return &pagination.PageData[domainPkg.ServiceCategory]{
		Data:  items,
		Total: total,
	}, nil
}

// tenantScopedQuery ensures tenant-level filtering.
func (r *entImpl) tenantScopedQuery(ctx *appctx.Context, qb *ent.ServiceCategoryQuery) *ent.ServiceCategoryQuery {
	if ctx.TenantID() != nil {
		qb = qb.Where(servicecategory.TenantIDEQ(*ctx.TenantID()))
	}
	return qb
}
