package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/platformuser"
	"github.com/umardev500/laundry/ent/user"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/platformuser/domain"
	"github.com/umardev500/laundry/internal/feature/platformuser/mapper"
	"github.com/umardev500/laundry/internal/feature/platformuser/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

type entImpl struct {
	client *entdb.Client
}

// Create implements Repository.
func (e *entImpl) Create(ctx *appctx.Context, pu *domain.PlatformUser) (*domain.PlatformUser, error) {
	conn := e.client.GetConn(ctx)
	builder := conn.PlatformUser.Create().
		SetUserID(pu.UserID).
		SetStatus(platformuser.Status(pu.Status))

	entPu, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntModel(entPu), nil
}

// Delete implements Repository.
func (e *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := e.client.GetConn(ctx)
	return conn.PlatformUser.DeleteOneID(id).Exec(ctx)
}

// FindById implements Repository.
func (e *entImpl) FindById(ctx *appctx.Context, id uuid.UUID) (*domain.PlatformUser, error) {
	conn := e.client.GetConn(ctx)

	entPu, err := conn.PlatformUser.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntModel(entPu), nil
}

// FindByIdUserID implements Repository.
func (e *entImpl) FindByIdUserID(ctx *appctx.Context, userID uuid.UUID) (*domain.PlatformUser, error) {
	conn := e.client.GetConn(ctx)

	entPu, err := conn.PlatformUser.
		Query().
		Where(platformuser.UserID(userID)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromEntModel(entPu), nil
}

// List implements Repository.
func (e *entImpl) List(ctx *appctx.Context, q *query.ListPlatformUserQuery) (*pagination.PageData[domain.PlatformUser], error) {
	q.Normalize()

	conn := e.client.GetConn(ctx)

	queryBuilder := conn.PlatformUser.Query()

	// Apply optional filter
	if q.Status != "" {
		queryBuilder = queryBuilder.Where(platformuser.StatusEQ(platformuser.Status(q.Status)))
	}

	if q.IncludeDeleted {
		queryBuilder = queryBuilder.Where(platformuser.DeletedAtIsNil())
	}

	if q.Search != "" {
		queryBuilder = queryBuilder.Where(
			platformuser.HasUserWith(
				user.Or(
					user.EmailContainsFold(q.Search),
				),
			),
		)
	}

	// Apply ordering
	switch q.Order {
	case query.OrderCreatedAtAsc:
		queryBuilder = queryBuilder.Order(ent.Asc(platformuser.FieldCreatedAt))
	case query.OrderCreatedAtDesc:
		queryBuilder = queryBuilder.Order(ent.Desc(platformuser.FieldCreatedAt))
	case query.OrderUpdatedAtAsc:
		queryBuilder = queryBuilder.Order(ent.Asc(platformuser.FieldUpdatedAt))
	case query.OrderUpdatedAtDesc:
		queryBuilder = queryBuilder.Order(ent.Desc(platformuser.FieldUpdatedAt))
	default:
		queryBuilder = queryBuilder.Order(ent.Asc(platformuser.FieldCreatedAt))
	}

	// Get total count first
	total, err := queryBuilder.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	// Paginate
	entPUs, err := queryBuilder.
		Limit(q.Limit).
		Offset(q.Offset()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	pus := mapper.FromEntModels(entPUs)

	return &pagination.PageData[domain.PlatformUser]{
		Data:  pus,
		Total: total,
	}, nil
}

// Update implements Repository.
func (e *entImpl) Update(ctx *appctx.Context, pu *domain.PlatformUser) (*domain.PlatformUser, error) {
	conn := e.client.GetConn(ctx)

	entPu, err := conn.PlatformUser.
		UpdateOneID(pu.ID).
		SetStatus(platformuser.Status(pu.Status)).
		SetNillableDeletedAt(pu.DeletedAt).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromEntModel(entPu), nil
}

// NewEntRepository creates a new PlatformUser repository.
func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{
		client: client,
	}
}
