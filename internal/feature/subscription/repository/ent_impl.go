package repository

import (
	"time"

	"github.com/google/uuid"

	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/subscription"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/subscription/domain"
	"github.com/umardev500/laundry/internal/feature/subscription/mapper"
	"github.com/umardev500/laundry/internal/feature/subscription/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

// entImpl implements the Subscription repository using Ent.
type entImpl struct {
	client *entdb.Client
}

// NewEntRepository returns a new Ent-based repository for Subscription.
func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{
		client: client,
	}
}

// Create inserts a new subscription record.
func (r *entImpl) Create(ctx *appctx.Context, s *domain.Subscription) (*domain.Subscription, error) {
	conn := r.client.GetConn(ctx)

	entModel, err := conn.Subscription.
		Create().
		SetTenantID(s.TenantID).
		SetPlanID(s.PlanID).
		SetStatus(subscription.Status(s.Status)).
		SetNillableStartDate(s.StartDate).
		SetNillableEndDate(s.EndDate).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindByID retrieves a subscription by its UUID (optionally including soft-deleted).
func (r *entImpl) FindByID(ctx *appctx.Context, id uuid.UUID, q *query.FindSubscriptionByIDQuery) (*domain.Subscription, error) {
	conn := r.client.GetConn(ctx)

	qb := conn.Subscription.Query().Where(subscription.IDEQ(id))
	if q == nil || !q.IncludeDeleted {
		qb = qb.Where(subscription.DeletedAtIsNil())
	}

	entModel, err := qb.Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// FindActiveByTenantID retrieves the active subscription for a tenant.
func (r *entImpl) FindActiveByTenantID(ctx *appctx.Context, tenantID uuid.UUID) (*domain.Subscription, error) {
	conn := r.client.GetConn(ctx)
	entModel, err := conn.Subscription.
		Query().
		Where(subscription.TenantIDEQ(tenantID)).
		Where(subscription.DeletedAtIsNil()).
		Where(subscription.StatusEQ(subscription.Status(types.SubscriptionStatusActive))).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEnt(entModel), nil
}

// Update modifies an existing subscription record.
func (r *entImpl) Update(ctx *appctx.Context, s *domain.Subscription) (*domain.Subscription, error) {
	conn := r.client.GetConn(ctx)

	builder := conn.Subscription.UpdateOneID(s.ID).
		SetStatus(subscription.Status(s.Status)).
		SetNillableStartDate(s.StartDate).
		SetNillableEndDate(s.EndDate).
		SetUpdatedAt(time.Now())

	if s.DeletedAt != nil {
		builder.SetDeletedAt(*s.DeletedAt)
	} else {
		builder.ClearDeletedAt()
	}

	entModel, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entModel), nil
}

// Delete removes a subscription by ID (hard delete).
func (r *entImpl) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)
	return conn.Subscription.DeleteOneID(id).Exec(ctx)
}

// List retrieves paginated subscriptions with filtering, ordering, and soft-deleted inclusion.
func (r *entImpl) List(ctx *appctx.Context, q *query.ListSubscriptionQuery) (*pagination.PageData[domain.Subscription], error) {
	conn := r.client.GetConn(ctx)
	q.Normalize()

	qb := conn.Subscription.Query()

	// Filtering
	if q.TenantID != "" {
		if tid, err := uuid.Parse(q.TenantID); err == nil {
			qb = qb.Where(subscription.TenantIDEQ(tid))
		}
	}
	if q.PlanID != "" {
		if pid, err := uuid.Parse(q.PlanID); err == nil {
			qb = qb.Where(subscription.PlanIDEQ(pid))
		}
	}
	if q.Status != nil {
		qb = qb.Where(subscription.StatusEQ(subscription.Status(*q.Status)))
	}
	if q.ActiveOnly != nil && *q.ActiveOnly {
		qb = qb.Where(subscription.StatusEQ(subscription.Status(types.SubscriptionStatusActive)))
	}
	if !q.IncludeDeleted {
		qb = qb.Where(subscription.DeletedAtIsNil())
	}

	// Date range filters
	if q.StartDateFrom != nil {
		qb = qb.Where(subscription.StartDateGTE(*q.StartDateFrom))
	}
	if q.StartDateTo != nil {
		qb = qb.Where(subscription.StartDateLTE(*q.StartDateTo))
	}
	if q.EndDateFrom != nil {
		qb = qb.Where(subscription.EndDateGTE(*q.EndDateFrom))
	}
	if q.EndDateTo != nil {
		qb = qb.Where(subscription.EndDateLTE(*q.EndDateTo))
	}

	// Sorting
	switch q.Order {
	case query.SubscriptionOrderStartDateAsc:
		qb = qb.Order(ent.Asc(subscription.FieldStartDate))
	case query.SubscriptionOrderStartDateDesc:
		qb = qb.Order(ent.Desc(subscription.FieldStartDate))
	case query.SubscriptionOrderEndDateAsc:
		qb = qb.Order(ent.Asc(subscription.FieldEndDate))
	case query.SubscriptionOrderEndDateDesc:
		qb = qb.Order(ent.Desc(subscription.FieldEndDate))
	case query.SubscriptionOrderUpdatedAtAsc:
		qb = qb.Order(ent.Asc(subscription.FieldUpdatedAt))
	case query.SubscriptionOrderUpdatedAtDesc:
		qb = qb.Order(ent.Desc(subscription.FieldUpdatedAt))
	case query.SubscriptionOrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(subscription.FieldCreatedAt))
	default:
		qb = qb.Order(ent.Desc(subscription.FieldCreatedAt))
	}

	// Pagination
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

	return &pagination.PageData[domain.Subscription]{
		Data:  items,
		Total: total,
	}, nil
}
