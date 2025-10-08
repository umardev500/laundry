package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/role"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/mapper"
	"github.com/umardev500/laundry/internal/feature/rbac/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

// entImpl implements RoleRepository using Ent as the persistence layer.
type entImpl struct {
	client *entdb.Client
}

// NewEntRepository returns a new Ent-based role repository.
func NewEntRepository(client *entdb.Client) RoleRepository {
	return &entImpl{client: client}
}

// Create inserts a new role record into the database.
func (e *entImpl) Create(ctx *appctx.Context, r *domain.Role) (*domain.Role, error) {
	conn := e.client.GetConn(ctx)

	entRole, err := conn.Role.
		Create().
		SetName(r.Name).
		SetDescription(r.Description).
		SetNillableTenantID(r.TenantID).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntModel(entRole), nil
}

// FindByID retrieves a role by its unique ID.
func (e *entImpl) FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.Role, error) {
	conn := e.client.GetConn(ctx)

	entRole, err := conn.Role.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntModel(entRole), nil
}

// FindByName retrieves a role by name, scoped by tenant if applicable.
func (e *entImpl) FindByName(ctx *appctx.Context, name string) (*domain.Role, error) {
	conn := e.client.GetConn(ctx)
	builder := conn.Role.Query().Where(role.NameEQ(name))

	if tenantID := ctx.TenantID(); tenantID != nil {
		builder = builder.Where(role.TenantIDEQ(*tenantID))
	}

	entRole, err := builder.Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntModel(entRole), nil
}

// Update modifies an existing role record.
func (e *entImpl) Update(ctx *appctx.Context, r *domain.Role) (*domain.Role, error) {
	conn := e.client.GetConn(ctx)

	entRole, err := conn.Role.
		UpdateOneID(r.ID).
		SetName(r.Name).
		SetDescription(r.Description).
		SetNillableDeletedAt(r.DeletedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntModel(entRole), nil
}

// Delete permanently removes a role record.
func (e *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := e.client.GetConn(ctx)
	return conn.Role.DeleteOneID(id).Exec(ctx)
}

// List returns paginated and filtered roles.
func (e *entImpl) List(ctx *appctx.Context, q *query.ListRoleQuery) (*pagination.PageData[domain.Role], error) {
	q.Normalize()

	conn := e.client.GetConn(ctx)
	builder := conn.Role.Query()

	// --- Tenant scoping ---
	if tenantID := ctx.TenantID(); tenantID != nil {
		// Force tenant restriction if context is tenant-scoped
		builder = builder.Where(role.TenantIDEQ(*tenantID))
	} else if q.TenantID != "" {
		// Allow filtering by tenantID if explicitly requested (admin context)
		if tid, err := uuid.Parse(q.TenantID); err == nil {
			builder = builder.Where(role.TenantIDEQ(tid))
		}
	}

	// --- Search filtering ---
	if q.Search != "" {
		builder = builder.Where(
			role.Or(
				role.NameContainsFold(q.Search),
				role.DescriptionContainsFold(q.Search),
			),
		)
	}

	// --- Deleted filtering ---
	if !q.IncludeDeleted {
		builder = builder.Where(role.DeletedAtIsNil())
	}

	// --- Ordering ---
	switch q.Order {
	case query.OrderNameAsc:
		builder = builder.Order(ent.Asc(role.FieldName))
	case query.OrderNameDesc:
		builder = builder.Order(ent.Desc(role.FieldName))
	case query.OrderCreatedAtAsc:
		builder = builder.Order(ent.Asc(role.FieldCreatedAt))
	case query.OrderCreatedAtDesc:
		builder = builder.Order(ent.Desc(role.FieldCreatedAt))
	case query.OrderUpdatedAtAsc:
		builder = builder.Order(ent.Asc(role.FieldUpdatedAt))
	case query.OrderUpdatedAtDesc:
		builder = builder.Order(ent.Desc(role.FieldUpdatedAt))
	default:
		builder = builder.Order(ent.Desc(role.FieldCreatedAt))
	}

	// --- Pagination ---
	total, err := builder.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	entRoles, err := builder.
		Limit(q.Limit).
		Offset(q.Offset()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	return &pagination.PageData[domain.Role]{
		Data:  mapper.FromEntModels(entRoles),
		Total: total,
	}, nil
}
