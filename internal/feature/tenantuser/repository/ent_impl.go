package repository

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/tenant"
	"github.com/umardev500/laundry/ent/tenantuser"
	"github.com/umardev500/laundry/ent/user"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/tenantuser/domain"
	"github.com/umardev500/laundry/internal/feature/tenantuser/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"

	mapper "github.com/umardev500/laundry/internal/feature/tenantuser/mapper"
)

// entImpl implements the Repository interface.
type entImpl struct {
	client *entdb.Client
}

// NewRepository creates a new TenantUser repository.
func NewRepository(client *entdb.Client) Repository {
	return &entImpl{client: client}
}

// Create a new TenantUser entry.
func (r *entImpl) Create(ctx *appctx.Context, tu *domain.TenantUser) (*domain.TenantUser, error) {
	conn := r.client.GetConn(ctx)

	entTenantUser, err := conn.TenantUser.Create().
		SetUserID(tu.UserID).
		SetTenantID(tu.TenantID).
		SetStatus(tenantuser.Status(tu.Status)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create tenant user: %w", err)
	}

	return mapper.FromEntModel(entTenantUser), nil
}

// FindByID retrieves a TenantUser by ID.
func (r *entImpl) FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.TenantUser, error) {
	conn := r.client.GetConn(ctx)

	entTenantUser, err := conn.TenantUser.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entTenantUser), nil
}

// FindByUser implements Repository.
func (r *entImpl) FindByUser(ctx *appctx.Context, userID uuid.UUID) ([]*domain.TenantUser, error) {
	conn := r.client.GetConn(ctx)

	db := conn.TenantUser.Query()
	items, err := db.
		Where(tenantuser.UserIDEQ(userID)).
		Where(tenantuser.DeletedAtIsNil()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntModels(items), nil
}

// FindByUserAndTenant retrieves a TenantUser by user_id and tenant_id.
func (r *entImpl) FindByUserAndTenant(ctx *appctx.Context, userID, tenantID uuid.UUID) (*domain.TenantUser, error) {
	conn := r.client.GetConn(ctx)

	entTenantUser, err := conn.TenantUser.
		Query().
		Where(
			tenantuser.UserIDEQ(userID),
			tenantuser.TenantIDEQ(tenantID),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntModel(entTenantUser), nil
}

// List retrieves tenant users with pagination and optional filters.
func (r *entImpl) List(ctx *appctx.Context, q *query.ListTenantUserQuery) (*pagination.PageData[domain.TenantUser], error) {
	var err error

	q.Normalize()
	conn := r.client.GetConn(ctx)

	qb := conn.TenantUser.Query()
	qb, err = r.applyScope(ctx, qb)
	if err != nil {
		return nil, err
	}

	if q.Search != "" {
		qb = qb.Where(
			tenantuser.HasTenantWith(tenant.NameContainsFold(q.Search)),
			tenantuser.HasUserWith(user.EmailContainsFold(q.Search)),
		)
	}

	if q.Status != "" {
		qb = qb.Where(tenantuser.StatusEQ(tenantuser.Status(q.Status)))
	}

	if !q.IncludeDeleted {
		qb = qb.Where(tenantuser.DeletedAtIsNil())
	}

	total, err := qb.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	// Ordering
	switch q.Order {
	case query.OrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(tenantuser.FieldCreatedAt))
	case query.OrderCreatedAtDesc:
		qb = qb.Order(ent.Desc(tenantuser.FieldCreatedAt))
	case query.OrderUpdatedAtAsc:
		qb = qb.Order(ent.Asc(tenantuser.FieldUpdatedAt))
	case query.OrderUpdatedAtDesc:
		qb = qb.Order(ent.Desc(tenantuser.FieldUpdatedAt))
	default:
		qb = qb.Order(ent.Desc(tenantuser.FieldCreatedAt))
	}

	// Pagination
	items, err := qb.
		Limit(q.Limit).
		Offset(q.Offset()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	results := mapper.FromEntModels(items)

	return pagination.NewPageData(results, total), nil
}

// Update updates an existing TenantUser record.
func (r *entImpl) Update(ctx *appctx.Context, tu *domain.TenantUser) (*domain.TenantUser, error) {
	conn := r.client.GetConn(ctx)

	entTenantUser, err := conn.TenantUser.
		UpdateOneID(tu.ID).
		SetStatus(tenantuser.Status(tu.Status)).
		SetUpdatedAt(tu.UpdatedAt).
		SetNillableDeletedAt(tu.DeletedAt).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed to update tenant user: %w", err)
	}

	return mapper.FromEntModel(entTenantUser), nil
}

// Delete permanently removes a tenant user (purge).
func (r *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	err := conn.TenantUser.
		DeleteOneID(id).
		Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete tenant user: %w", err)
	}

	log.Debug().Str("tenant_user_id", id.String()).Msg("Tenant user purged")
	return nil
}

// -------------------------
// Helpers
// -------------------------

// applyScope ensures tenant-level filtering.
func (r *entImpl) applyScope(ctx *appctx.Context, qb *ent.TenantUserQuery) (*ent.TenantUserQuery, error) {
	switch ctx.Scope() {
	case appctx.ScopeTenant:
		qb = qb.Where(tenantuser.TenantIDEQ(*ctx.TenantID()))
	case appctx.ScopeUser:
		// RBAC already handles access, no filtering needed here
	case appctx.ScopeAdmin:
	// no filtering for admin
	default:
		// Unknown scope, deny access by default
		return nil, domain.ErrUnauthorizedUserAccess
	}

	return qb, nil
}
