package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/order"
	"github.com/umardev500/laundry/ent/orderstatushistory"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/order/domain"
	"github.com/umardev500/laundry/internal/feature/order/mapper"
	"github.com/umardev500/laundry/internal/feature/order/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

// entImpl implements Repository using Ent.
type entImpl struct {
	client *entdb.Client
}

// FindById implements Repository.
func (r *entImpl) FindById(ctx *appctx.Context, id uuid.UUID, q *query.OrderQuery) (*domain.Order, error) {
	conn := r.client.GetConn(ctx)
	qb := conn.Order.Query().
		Where(order.IDEQ(id))

	// Conditionally preload items
	if q.IncludeItems {
		qb = qb.WithItems()
	}

	if q.IncludePayment {
		qb = qb.WithPayment(func(pq *ent.PaymentQuery) {
			if q.IncludePaymentMethod {
				pq.WithPaymentMethod()
			}
		})
	}

	// Conditionally preload status
	// Conditionally preload status
	if q.IncludeStatuses {
		orderFunc := ent.Desc(orderstatushistory.FieldCreatedAt)
		if q.StatusOrder == query.StatusesOrderAsc {
			orderFunc = ent.Asc(orderstatushistory.FieldCreatedAt)
		}
		qb = qb.WithStatusHistory(func(shq *ent.OrderStatusHistoryQuery) {
			shq.Order(orderFunc)
		})
	}

	// Exclude deleted orders
	if !q.IncludeDeleted {
		qb = qb.Where(order.DeletedAtIsNil())
	}

	orderObj, err := qb.Only(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(orderObj), nil
}

// Update implements Repository.
func (r *entImpl) Update(ctx *appctx.Context, o *domain.Order) (*domain.Order, error) {
	conn := r.client.GetConn(ctx)
	builder := conn.Order.UpdateOneID(o.ID).
		SetStatus(order.Status(o.Status)).
		SetNillableNotes(o.Notes).
		SetTotalAmount(o.TotalAmount).
		SetNillableGuestName(o.GuestName).
		SetNillableGuestEmail(o.GuestEmail).
		SetNillableGuestPhone(o.GuestPhone).
		SetNillableGuestAddress(o.GuestAddress).
		SetNillablePaymentID(o.PaymentID)

	orderObj, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(orderObj), nil
}

// NewEntRepository creates a new Ent repository.
func NewEntRepository(client *entdb.Client) Repository {
	return &entImpl{client: client}
}

// Create implements Repository.
func (r *entImpl) Create(ctx *appctx.Context, o *domain.Order) (*domain.Order, error) {
	conn := r.client.GetConn(ctx)

	builder := conn.Order.Create().
		SetTenantID(o.TenantID).
		SetNillableNotes(o.Notes).
		SetStatus(order.Status(o.Status)).
		SetTotalAmount(o.TotalAmount).
		SetNillableGuestName(o.GuestName).
		SetNillableGuestEmail(o.GuestEmail).
		SetNillableGuestPhone(o.GuestPhone).
		SetNillableGuestAddress(o.GuestAddress)

	orderObj, err := builder.Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(orderObj), err
}

// List returns paginated orders with filtering and ordering.
func (r *entImpl) List(ctx *appctx.Context, q *query.ListOrderQuery) (*pagination.PageData[domain.Order], error) {
	q.Normalize()

	conn := r.client.GetConn(ctx)
	qb := conn.Order.Query()
	qb = r.applyScope(ctx, qb)

	// Filter by status if provided
	if q.Status != nil && *q.Status != "" {
		qb = qb.Where(order.StatusEQ(order.Status(*q.Status)))
	}

	// Conditionally preload items
	if q.IncludeItems {
		qb = qb.WithItems()
	}

	if q.IncludePayment {
		qb = qb.WithPayment(func(pq *ent.PaymentQuery) {
			if q.IncludePaymentMethod {
				pq.WithPaymentMethod()
			}
		})
	}

	// Conditionally preload status
	if q.IncludeStatuses {
		orderFunc := ent.Desc(orderstatushistory.FieldCreatedAt)
		if q.StatusOrder == query.StatusesOrderAsc {
			orderFunc = ent.Asc(orderstatushistory.FieldCreatedAt)
		}
		qb = qb.WithStatusHistory(func(shq *ent.OrderStatusHistoryQuery) {
			shq.Order(orderFunc)
		})
	}

	// Search by guest info
	if q.Search != "" {
		qb = qb.Where(
			order.Or(
				order.GuestNameContainsFold(q.Search),
				order.GuestEmailContainsFold(q.Search),
				order.GuestPhoneContainsFold(q.Search),
			),
		)
	}

	// Exclude deleted orders
	if !q.IncludeDeleted {
		qb = qb.Where(order.DeletedAtIsNil())
	}

	// Apply ordering
	switch q.Order {
	case query.OrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(order.FieldCreatedAt))
	case query.OrderCreatedAtDesc:
		qb = qb.Order(ent.Desc(order.FieldCreatedAt))
	case query.OrderUpdatedAtAsc:
		qb = qb.Order(ent.Asc(order.FieldUpdatedAt))
	case query.OrderUpdatedAtDesc:
		qb = qb.Order(ent.Desc(order.FieldUpdatedAt))
	case query.OrderTotalAmountAsc:
		qb = qb.Order(ent.Asc(order.FieldTotalAmount))
	case query.OrderTotalAmountDesc:
		qb = qb.Order(ent.Desc(order.FieldTotalAmount))
	default:
		qb = qb.Order(ent.Desc(order.FieldCreatedAt))
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

	items := mapper.FromEntList(ents)

	return &pagination.PageData[domain.Order]{
		Data:  items,
		Total: total,
	}, nil
}

// applyScope ensures tenant-level filtering.
func (r *entImpl) applyScope(ctx *appctx.Context, qb *ent.OrderQuery) *ent.OrderQuery {
	switch ctx.Scope() {
	case appctx.ScopeTenant:
		qb = qb.Where(order.TenantIDEQ(*ctx.TenantID()))
	case appctx.ScopeUser:
		qb = qb.Where(order.UserIDEQ(*ctx.UserID()))
	case appctx.ScopeAdmin:
		// no filtering for admin
	}

	return qb
}
