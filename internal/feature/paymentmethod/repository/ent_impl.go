package repository

import (
	"github.com/google/uuid"

	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/paymentmethod"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/domain"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/mapper"
	"github.com/umardev500/laundry/internal/feature/paymentmethod/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

// entImpl implements the PaymentMethod repository using Ent.
type entImpl struct {
	client *entdb.Client
}

// NewEntRepository returns a new Ent-based repository for PaymentMethod.
func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{
		client: client,
	}
}

// Create inserts a new payment method
func (r *entImpl) Create(ctx *appctx.Context, pm *domain.PaymentMethod) (*domain.PaymentMethod, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.PaymentMethod.
		Create().
		SetName(pm.Name).
		SetNillableDescription(pm.Description).
		SetCreatedAt(pm.CreatedAt).
		SetUpdatedAt(pm.UpdatedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByID retrieves a payment method by its UUID
func (r *entImpl) FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.PaymentMethod, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.PaymentMethod.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByName retrieves a payment method by name
func (r *entImpl) FindByName(ctx *appctx.Context, name string) (*domain.PaymentMethod, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.PaymentMethod.
		Query().
		Where(paymentmethod.NameEQ(name)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Update modifies an existing payment method
func (r *entImpl) Update(ctx *appctx.Context, pm *domain.PaymentMethod) (*domain.PaymentMethod, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.PaymentMethod.
		UpdateOneID(pm.ID).
		SetName(pm.Name).
		SetNillableDescription(pm.Description).
		SetNillableDeletedAt(pm.DeletedAt).
		SetUpdatedAt(pm.UpdatedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Delete performs a hard delete
func (r *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)
	err := conn.PaymentMethod.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

// List returns a paginated list of payment methods
func (r *entImpl) List(ctx *appctx.Context, q *query.ListPaymentMethodQuery) (*pagination.PageData[domain.PaymentMethod], error) {
	q.Normalize()
	conn := r.client.GetConn(ctx)

	qb := conn.PaymentMethod.Query()

	if q.Search != "" {
		qb = qb.Where(
			paymentmethod.Or(
				paymentmethod.NameContainsFold(q.Search),
				paymentmethod.DescriptionContainsFold(q.Search),
			),
		)
	}

	if !q.IncludeDeleted {
		qb = qb.Where(paymentmethod.DeletedAtIsNil())
	}

	switch q.Order {
	case query.OrderNameAsc:
		qb = qb.Order(ent.Asc(paymentmethod.FieldName))
	case query.OrderNameDesc:
		qb = qb.Order(ent.Desc(paymentmethod.FieldName))
	case query.OrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(paymentmethod.FieldCreatedAt))
	default:
		qb = qb.Order(ent.Desc(paymentmethod.FieldCreatedAt))
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

	return &pagination.PageData[domain.PaymentMethod]{
		Data:  items,
		Total: total,
	}, nil
}
