package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/permission"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/mapper"
	"github.com/umardev500/laundry/internal/feature/rbac/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

type permissionRepoEnt struct {
	client *entdb.Client
}

func NewPermissionRepository(client *entdb.Client) PermissionRepository {
	return &permissionRepoEnt{client: client}
}

func (r *permissionRepoEnt) FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.Permission, error) {
	conn := r.client.GetConn(ctx)
	entPerm, err := conn.Permission.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntPermission(entPerm), nil
}

func (r *permissionRepoEnt) Update(ctx *appctx.Context, p *domain.Permission) (*domain.Permission, error) {
	conn := r.client.GetConn(ctx)
	entPerm, err := conn.Permission.
		UpdateOneID(p.ID).
		SetName(p.Name).
		SetDisplayName(p.DisplayName).
		SetDescription(p.Description).
		SetStatus(permission.Status(p.Status)).
		SetNillableDeletedAt(p.DeletedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEntPermission(entPerm), nil
}

func (r *permissionRepoEnt) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)
	return conn.Permission.DeleteOneID(id).Exec(ctx)
}

func (r *permissionRepoEnt) List(ctx *appctx.Context, q *query.ListPermissionQuery) (*pagination.PageData[domain.Permission], error) {
	q.Normalize()

	conn := r.client.GetConn(ctx)
	builder := conn.Permission.Query()

	if q.Search != "" {
		builder = builder.Where(
			permission.Or(
				permission.NameContainsFold(q.Search),
				permission.DisplayNameContainsFold(q.Search),
			),
		)
	}
	if q.Status != "" {
		builder = builder.Where(permission.StatusEQ(permission.Status(q.Status)))
	}
	if !q.IncludeDeleted {
		builder = builder.Where(permission.DeletedAtIsNil())
	}

	switch q.Order {
	case query.PermissionOrderNameAsc:
		builder = builder.Order(ent.Asc(permission.FieldName))
	case query.PermissionOrderNameDesc:
		builder = builder.Order(ent.Desc(permission.FieldName))
	case query.PermissionOrderDisplayNameAsc:
		builder = builder.Order(ent.Asc(permission.FieldDisplayName))
	case query.PermissionOrderDisplayNameDesc:
		builder = builder.Order(ent.Desc(permission.FieldDisplayName))
	case query.PermissionOrderUpdatedAtAsc:
		builder = builder.Order(ent.Asc(permission.FieldUpdatedAt))
	case query.PermissionOrderUpdatedAtDesc:
		builder = builder.Order(ent.Desc(permission.FieldUpdatedAt))
	default:
		builder = builder.Order(ent.Desc(permission.FieldCreatedAt))
	}

	total, err := builder.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	items, err := builder.Limit(q.Limit).Offset(q.Offset()).All(ctx)
	if err != nil {
		return nil, err
	}

	return &pagination.PageData[domain.Permission]{
		Data:  mapper.FromEntPermissionList(items),
		Total: total,
	}, nil
}
