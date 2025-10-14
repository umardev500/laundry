package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/order"
	"github.com/umardev500/laundry/ent/orderstatushistory"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/domain"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/mapper"
	"github.com/umardev500/laundry/internal/feature/orderstatushistory/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

// entStatusHistoryImpl implements StatusHistoryRepository using Ent.
type entStatusHistoryImpl struct {
	client *entdb.Client
}

// NewEntStatusHistoryRepository creates a new repository.
func NewEntStatusHistoryRepository(client *entdb.Client) StatusHistoryRepository {
	return &entStatusHistoryImpl{client: client}
}

// Create inserts a new status milestone.
func (r *entStatusHistoryImpl) Create(ctx *appctx.Context, sh *domain.OrderStatusHistory) (*domain.OrderStatusHistory, error) {
	conn := r.client.GetConn(ctx)

	builder := conn.OrderStatusHistory.Create().
		SetOrderID(sh.OrderID).
		SetStatus(orderstatushistory.Status(sh.Status)).
		SetNillableNotes(sh.Notes)

	if sh.Notes != nil {
		builder.SetNillableNotes(sh.Notes)
	}

	shEnt, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntStatusHistory(shEnt), nil
}

// FindById retrieves a status milestone by ID.
func (r *entStatusHistoryImpl) FindById(ctx *appctx.Context, id uuid.UUID, q *query.StatusHistoryByIDQuery) (*domain.OrderStatusHistory, error) {
	conn := r.client.GetConn(ctx)
	qb := conn.OrderStatusHistory.
		Query().
		Where(orderstatushistory.IDEQ(id))
	qb = r.applyScope(ctx, qb)

	if q.IncludeOrder {
		qb = qb.WithOrder(func(oq *ent.OrderQuery) {
			if q.IncludeOrderRef {
				oq.WithItems().WithPayment()
			}
		})
	}

	shEnt, err := qb.Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEntStatusHistory(shEnt), nil
}

// List retrieves paginated status milestones with filtering.
func (r *entStatusHistoryImpl) List(ctx *appctx.Context, q *query.OrderStatusHistoryListQuery) (*pagination.PageData[domain.OrderStatusHistory], error) {
	q.Normalize()
	conn := r.client.GetConn(ctx)

	qb := conn.OrderStatusHistory.Query()
	qb = r.applyScope(ctx, qb)

	// Filter by order ID
	if q.OrderID != nil && *q.OrderID != uuid.Nil {
		qb = qb.Where(orderstatushistory.OrderIDEQ(*q.OrderID))
	}

	// Filter by status
	if q.Status != "" {
		qb = qb.Where(orderstatushistory.StatusEQ(orderstatushistory.Status(q.Status)))
	}

	if q.IncludeOrder {
		qb = qb.WithOrder(func(oq *ent.OrderQuery) {
			if q.IncludeOrderRef {
				oq.WithItems().WithPayment()
			}
		})
	}

	// Apply ordering
	switch q.Order {
	case query.StatusCreatedAtAsc:
		qb = qb.Order(ent.Asc(orderstatushistory.FieldCreatedAt))
	case query.StatusCreatedAtDesc:
		qb = qb.Order(ent.Desc(orderstatushistory.FieldCreatedAt))
	default:
		qb = qb.Order(ent.Desc(orderstatushistory.FieldCreatedAt))
	}

	// Count total results
	total, err := qb.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	// Apply pagination
	ents, err := qb.Limit(q.Limit).Offset(q.Offset()).All(ctx)
	if err != nil {
		return nil, err
	}

	items := mapper.FromEntStatusHistoryList(ents)

	return &pagination.PageData[domain.OrderStatusHistory]{
		Data:  items,
		Total: total,
	}, nil
}

// -------------------------
// Helpers
// -------------------------

// applyScope ensures tenant-level filtering.
func (r *entStatusHistoryImpl) applyScope(ctx *appctx.Context, qb *ent.OrderStatusHistoryQuery) *ent.OrderStatusHistoryQuery {
	switch ctx.Scope() {
	case appctx.ScopeTenant:
		qb = qb.Where(orderstatushistory.HasOrderWith(order.TenantIDEQ(*ctx.TenantID())))
	case appctx.ScopeUser:
		qb = qb.Where(orderstatushistory.HasOrderWith(order.UserIDEQ(*ctx.UserID())))
	case appctx.ScopeAdmin:
		// no filtering for admin
	}

	return qb
}
