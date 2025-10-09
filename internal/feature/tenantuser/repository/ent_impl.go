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
	"github.com/umardev500/laundry/pkg/types"
)

// tenantUserRepository implements the Repository interface.
type tenantUserRepository struct {
	client *entdb.Client
}

// NewRepository creates a new TenantUser repository.
func NewRepository(client *entdb.Client) Repository {
	return &tenantUserRepository{client: client}
}

// Create a new TenantUser entry.
func (r *tenantUserRepository) Create(ctx *appctx.Context, tu *domain.TenantUser) (*domain.TenantUser, error) {
	conn := r.client.GetConn(ctx)

	entTenantUser, err := conn.TenantUser.Create().
		SetID(tu.ID).
		SetUserID(tu.UserID).
		SetTenantID(tu.TenantID).
		SetStatus(tenantuser.Status(tu.Status)).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to create tenant user: %w", err)
	}

	return mapToDomain(entTenantUser), nil
}

// FindByID retrieves a TenantUser by ID.
func (r *tenantUserRepository) FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.TenantUser, error) {
	conn := r.client.GetConn(ctx)

	entTenantUser, err := conn.TenantUser.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapToDomain(entTenantUser), nil
}

// FindByUserAndTenant retrieves a TenantUser by user_id and tenant_id.
func (r *tenantUserRepository) FindByUserAndTenant(ctx *appctx.Context, userID, tenantID uuid.UUID) (*domain.TenantUser, error) {
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
	return mapToDomain(entTenantUser), nil
}

// List retrieves tenant users with pagination and optional filters.
func (r *tenantUserRepository) List(ctx *appctx.Context, q *query.ListTenantUserQuery) (*pagination.PageData[domain.TenantUser], error) {
	q.Normalize()
	conn := r.client.GetConn(ctx)

	db := conn.TenantUser.Query()

	if q.Search != "" {
		db = db.Where(
			tenantuser.HasTenantWith(tenant.NameContainsFold(q.Search)),
			tenantuser.HasUserWith(user.EmailContainsFold(q.Search)),
		)
	}

	if q.Status != "" {
		db = db.Where(tenantuser.StatusEQ(tenantuser.Status(q.Status)))
	}

	if !q.IncludeDeleted {
		db = db.Where(tenantuser.DeletedAtIsNil())
	}

	total, err := db.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	// Ordering
	switch q.Order {
	case query.OrderCreatedAtAsc:
		db = db.Order(ent.Asc(tenantuser.FieldCreatedAt))
	case query.OrderCreatedAtDesc:
		db = db.Order(ent.Desc(tenantuser.FieldCreatedAt))
	case query.OrderUpdatedAtAsc:
		db = db.Order(ent.Asc(tenantuser.FieldUpdatedAt))
	case query.OrderUpdatedAtDesc:
		db = db.Order(ent.Desc(tenantuser.FieldUpdatedAt))
	default:
		db = db.Order(ent.Desc(tenantuser.FieldCreatedAt))
	}

	// Pagination
	items, err := db.
		Limit(q.Limit).
		Offset(q.Offset()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	results := make([]*domain.TenantUser, len(items))
	for i, item := range items {
		results[i] = mapToDomain(item)
	}

	return pagination.NewPageData(results, total), nil
}

// Update updates an existing TenantUser record.
func (r *tenantUserRepository) Update(ctx *appctx.Context, tu *domain.TenantUser) (*domain.TenantUser, error) {
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

	return mapToDomain(entTenantUser), nil
}

// Delete permanently removes a tenant user (purge).
func (r *tenantUserRepository) Delete(ctx *appctx.Context, id uuid.UUID) error {
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

// mapToDomain converts Ent model to domain model.
func mapToDomain(tu *ent.TenantUser) *domain.TenantUser {
	return &domain.TenantUser{
		ID:        tu.ID,
		UserID:    tu.UserID,
		TenantID:  tu.TenantID,
		Status:    types.Status(*tu.Status),
		CreatedAt: tu.CreatedAt,
		UpdatedAt: tu.UpdatedAt,
		DeletedAt: tu.DeletedAt,
	}
}
