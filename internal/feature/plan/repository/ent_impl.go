package repository

import (
	"github.com/google/uuid"

	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/plan"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/plan/domain"
	"github.com/umardev500/laundry/internal/feature/plan/mapper"
	"github.com/umardev500/laundry/internal/feature/plan/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

// entImpl implements the Plan repository using Ent.
type entImpl struct {
	client *entdb.Client
}

// NewEntRepository returns a new Ent-based repository for Plan.
func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{
		client: client,
	}
}

// Create inserts a new plan record.
func (r *entImpl) Create(ctx *appctx.Context, p *domain.Plan) (*domain.Plan, error) {
	conn := r.client.GetConn(ctx)

	featuresMap, err := p.Features.ToMap()
	if err != nil {
		return nil, err
	}
	entModel, err := conn.Plan.
		Create().
		SetName(p.Name).
		SetNillableDescription(p.Description).
		SetPrice(p.Price).
		SetBillingInterval(plan.BillingInterval(p.BillingInterval)).
		SetFeatures(featuresMap).
		SetActive(p.Active).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByID retrieves a plan by its UUID.
func (r *entImpl) FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.Plan, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.Plan.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByName retrieves a plan by its name.
func (r *entImpl) FindByName(ctx *appctx.Context, name string) (*domain.Plan, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.Plan.Query().
		Where(plan.NameEQ(name)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Update modifies an existing plan record.
func (r *entImpl) Update(ctx *appctx.Context, p *domain.Plan) (*domain.Plan, error) {
	conn := r.client.GetConn(ctx)

	featuresMap, err := p.Features.ToMap()
	if err != nil {
		return nil, err
	}

	builder := conn.Plan.UpdateOneID(p.ID).
		SetName(p.Name).
		SetNillableDescription(p.Description).
		SetPrice(p.Price).
		SetBillingInterval(plan.BillingInterval(p.BillingInterval)).
		SetFeatures(featuresMap).
		SetActive(p.Active)

	if p.DeletedAt != nil {
		builder.SetDeletedAt(*p.DeletedAt)
	} else {
		builder.ClearDeletedAt() // reset column to NULL
	}

	entModel, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Delete removes a plan by ID (hard delete).
func (r *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)
	return conn.Plan.DeleteOneID(id).Exec(ctx)
}

// List retrieves paginated plans with filtering, ordering, and soft-deleted inclusion.
func (r *entImpl) List(ctx *appctx.Context, q *query.ListPlanQuery) (*pagination.PageData[domain.Plan], error) {
	conn := r.client.GetConn(ctx)
	q.Normalize()

	qb := conn.Plan.Query()

	if q.Search != "" {
		qb = qb.Where(
			plan.Or(
				plan.NameContainsFold(q.Search),
				plan.DescriptionContainsFold(q.Search),
			),
		)
	}

	if !q.IncludeDeleted {
		qb = qb.Where(plan.DeletedAtIsNil())
	}

	switch q.Order {
	case query.PlanOrderNameAsc:
		qb = qb.Order(ent.Asc(plan.FieldName))
	case query.PlanOrderNameDesc:
		qb = qb.Order(ent.Desc(plan.FieldName))
	case query.PlanOrderPriceAsc:
		qb = qb.Order(ent.Asc(plan.FieldPrice))
	case query.PlanOrderPriceDesc:
		qb = qb.Order(ent.Desc(plan.FieldPrice))
	case query.PlanOrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(plan.FieldCreatedAt))
	case query.PlanOrderCreatedAtDesc:
		qb = qb.Order(ent.Desc(plan.FieldCreatedAt))
	case query.PlanOrderUpdatedAtAsc:
		qb = qb.Order(ent.Asc(plan.FieldUpdatedAt))
	case query.PlanOrderUpdatedAtDesc:
		qb = qb.Order(ent.Desc(plan.FieldUpdatedAt))
	default:
		qb = qb.Order(ent.Desc(plan.FieldUpdatedAt))
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

	return &pagination.PageData[domain.Plan]{
		Data:  items,
		Total: total,
	}, nil
}
