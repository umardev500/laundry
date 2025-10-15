package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/tenant"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/tenant/domain"
	"github.com/umardev500/laundry/internal/feature/tenant/mapper"
	"github.com/umardev500/laundry/internal/feature/tenant/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

type entImpl struct {
	client *entdb.Client
}

func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{client: client}
}

// Create a new tenant.
func (e *entImpl) Create(ctx *appctx.Context, t *domain.Tenant) (*domain.Tenant, error) {
	conn := e.client.GetConn(ctx)

	entTenant, err := conn.Tenant.
		Create().
		SetName(t.Name).
		SetEmail(t.Email).
		SetPhone(t.Phone).
		SetStatus(tenant.Status(t.Status)).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntModel(entTenant), nil
}

func (e *entImpl) FindById(ctx *appctx.Context, id uuid.UUID) (*domain.Tenant, error) {
	conn := e.client.GetConn(ctx)
	entTenant, err := conn.Tenant.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entTenant), nil
}

func (e *entImpl) FindByEmail(ctx *appctx.Context, email string) (*domain.Tenant, error) {
	conn := e.client.GetConn(ctx)
	entTenant, err := conn.Tenant.
		Query().
		Where(tenant.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entTenant), nil
}

func (e *entImpl) Update(ctx *appctx.Context, t *domain.Tenant) (*domain.Tenant, error) {
	conn := e.client.GetConn(ctx)
	entTenant, err := conn.Tenant.
		UpdateOneID(t.ID).
		SetName(t.Name).
		SetEmail(t.Email).
		SetPhone(t.Phone).
		SetStatus(tenant.Status(t.Status)).
		SetNillableDeletedAt(t.DeletedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entTenant), nil
}

func (e *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := e.client.GetConn(ctx)
	return conn.Tenant.DeleteOneID(id).Exec(ctx)
}

func (e *entImpl) List(ctx *appctx.Context, q *query.ListTenantQuery) (*pagination.PageData[domain.Tenant], error) {
	var err error
	q.Normalize()

	conn := e.client.GetConn(ctx)
	qb := conn.Tenant.Query()

	qb, err = e.applyScope(ctx, qb)
	if err != nil {
		return nil, err
	}

	// 🔍 Apply search filter
	if q.Search != "" {
		qb = qb.Where(
			tenant.Or(
				tenant.NameContainsFold(q.Search),
				tenant.EmailContainsFold(q.Search),
			),
		)
	}

	// ⚙️ Apply status filter
	if q.Status != "" {
		qb = qb.Where(
			tenant.StatusEQ(tenant.Status(q.Status)),
		)
	}

	// 🗑️ Include/exclude deleted tenants
	if !q.IncludeDeleted {
		qb = qb.Where(tenant.DeletedAtIsNil())
	}

	// ↕️ Apply ordering
	switch q.Order {
	case query.OrderNameAsc:
		qb = qb.Order(ent.Asc(tenant.FieldName))
	case query.OrderNameDesc:
		qb = qb.Order(ent.Desc(tenant.FieldName))
	case query.OrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(tenant.FieldCreatedAt))
	case query.OrderCreatedAtDesc:
		qb = qb.Order(ent.Desc(tenant.FieldCreatedAt))
	case query.OrderUpdatedAtAsc:
		qb = qb.Order(ent.Asc(tenant.FieldUpdatedAt))
	case query.OrderUpdatedAtDesc:
		qb = qb.Order(ent.Desc(tenant.FieldUpdatedAt))
	default:
		qb = qb.Order(ent.Desc(tenant.FieldCreatedAt))
	}

	// 📊 Count total before pagination
	total, err := qb.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	// 📄 Paginated results
	entTenants, err := qb.
		Limit(q.Limit).
		Offset(q.Offset()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	tenants := mapper.FromEntModels(entTenants)

	return &pagination.PageData[domain.Tenant]{
		Data:  tenants,
		Total: total,
	}, nil
}

// -------------------------
// Helpers
// -------------------------

// applyScope ensures tenant-level filtering.
func (r *entImpl) applyScope(ctx *appctx.Context, qb *ent.TenantQuery) (*ent.TenantQuery, error) {
	switch ctx.Scope() {
	case appctx.ScopeTenant:
		qb = qb.Where(tenant.IDEQ(*ctx.TenantID()))
	case appctx.ScopeUser:
		// RBAC already handles access, no filtering needed here
	case appctx.ScopeAdmin:
	// no filtering for admin
	default:
		// Unknown scope, deny access by default
		return nil, domain.ErrUnauthorizedTenant
	}

	return qb, nil
}
