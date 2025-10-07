package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent/user"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/user/domain"
	"github.com/umardev500/laundry/internal/feature/user/mapper"
	"github.com/umardev500/laundry/internal/feature/user/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

type entImpl struct {
	client *entdb.Client
}

func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{
		client: client,
	}
}

// Create implements Repository.
func (e *entImpl) Create(ctx *appctx.Context, u *domain.User) (*domain.User, error) {
	conn := e.client.GetConn(ctx)
	builder := conn.User.
		Create().
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetStatus(user.Status(u.Status))
	entUser, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.ToDomainUser(entUser), nil
}

// Delete implements Repository.
func (e *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := e.client.GetConn(ctx)
	return conn.User.DeleteOneID(id).Exec(ctx)
}

// FindByEmail implements Repository.
func (e *entImpl) FindByEmail(ctx *appctx.Context, email string) (*domain.User, error) {
	conn := e.client.GetConn(ctx)

	entUser, err := conn.User.
		Query().
		Where(user.EmailEQ(email)).
		Only(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.ToDomainUser(entUser), nil
}

// FindById implements Repository.
func (e *entImpl) FindById(ctx *appctx.Context, id uuid.UUID) (*domain.User, error) {
	conn := e.client.GetConn(ctx)

	entUser, err := conn.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.ToDomainUser(entUser), nil
}

// List implements Repository.
func (e *entImpl) List(ctx *appctx.Context, q *query.ListUserQuery) (*pagination.PageData[domain.User], error) {
	conn := e.client.GetConn(ctx)
	queryBuilder := conn.User.Query()

	// Applu optional filters
	if q.Search != "" {
		queryBuilder = queryBuilder.Where(user.EmailContainsFold(q.Search))
	}

	// Get total count first
	total, err := queryBuilder.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	// Get paginated results
	entUsers, err := queryBuilder.
		Limit(q.Limit).
		Offset(q.Offset()).
		All(ctx)

	if err != nil {
		return nil, err
	}

	users := mapper.ToDomainUsers(entUsers)

	return &pagination.PageData[domain.User]{
		Data:  users,
		Total: total,
	}, nil
}

// Update implements Repository.
func (e *entImpl) Update(ctx *appctx.Context, u *domain.User) (*domain.User, error) {
	conn := e.client.GetConn(ctx)

	entUser, err := conn.User.
		UpdateOneID(u.ID).
		SetEmail(u.Email).
		SetPassword(u.Password).
		SetStatus(user.Status(u.Status)).
		SetNillableDeletedAt(u.DeletedAt).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.ToDomainUser(entUser), nil
}
