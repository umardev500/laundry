package repository

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/payment"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/payment/domain"
	"github.com/umardev500/laundry/internal/feature/payment/mapper"
	"github.com/umardev500/laundry/internal/feature/payment/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

// EntPaymentRepository implements domain.Payment repository using Ent
type EntPaymentRepository struct {
	client *entdb.Client
}

// NewEntPaymentRepository creates a new repository instance
func NewEntPaymentRepository(client *entdb.Client) Repository {
	return &EntPaymentRepository{
		client: client,
	}
}

// Create inserts a new payment
func (r *EntPaymentRepository) Create(ctx *appctx.Context, p *domain.Payment) (*domain.Payment, error) {
	entPayment, err := r.client.Client.Payment.
		Create().
		SetNillableUserID(p.UserID).
		SetNillableTenantID(p.TenantID).
		SetRefID(p.RefID).
		SetRefType(payment.RefType(p.RefType)).
		SetPaymentMethodID(p.PaymentMethodID).
		SetAmount(p.Amount).
		SetStatus(payment.Status(p.Status)).
		SetNotes(p.Notes).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		SetNillableReceivedAmount(p.ReceivedAmount).
		SetNillableChangeAmount(p.ChangeAmount).
		SetNillablePaidAt(p.PaidAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entPayment), nil
}

// Update modifies an existing payment
func (r *EntPaymentRepository) Update(ctx *appctx.Context, p *domain.Payment) (*domain.Payment, error) {
	entPayment, err := r.client.Client.Payment.
		UpdateOneID(p.ID).
		SetPaymentMethodID(p.PaymentMethodID).
		SetAmount(p.Amount).
		SetStatus(payment.Status(p.Status)).
		SetNotes(p.Notes).
		SetUpdatedAt(time.Now()).
		SetNillableReceivedAmount(p.ReceivedAmount).
		SetNillableChangeAmount(p.ChangeAmount).
		SetNillablePaidAt(p.PaidAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entPayment), nil
}

// FindById returns a payment by its ID
func (r *EntPaymentRepository) FindById(ctx *appctx.Context, id uuid.UUID) (*domain.Payment, error) {
	entPayment, err := r.client.Client.Payment.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return mapper.FromEnt(entPayment), nil
}

// Delete performs a hard delete
func (r *EntPaymentRepository) Delete(ctx *appctx.Context, id uuid.UUID) error {
	err := r.client.Client.Payment.
		DeleteOneID(id).
		Exec(ctx)
	return err
}

// List retrieves paginated payments with filtering, ordering, and tenant scoping.
func (r *EntPaymentRepository) List(ctx *appctx.Context, q *query.ListPaymentQuery) (*pagination.PageData[domain.Payment], error) {
	q.Normalize()

	conn := r.client.GetConn(ctx)
	qb := conn.Payment.Query()
	qb = r.applyScope(ctx, qb)

	// Search by notes
	if q.Search != "" {
		qb = qb.Where(payment.NotesContainsFold(q.Search))
	}

	// Filter by status
	if q.Status != "" {
		qb = qb.Where(payment.StatusEQ(payment.Status(q.Status)))
	}

	// Filter by ref type
	if q.RefType != "" {
		qb = qb.Where(payment.RefTypeEQ(payment.RefType(q.RefType)))
	}

	// Filter deleted
	if !q.IncludeDeleted {
		qb = qb.Where(payment.DeletedAtIsNil())
	}

	// Ordering
	switch q.Order {
	case query.PaymentOrderCreatedAtAsc:
		qb = qb.Order(ent.Asc(payment.FieldCreatedAt))
	case query.PaymentOrderCreatedAtDesc:
		qb = qb.Order(ent.Desc(payment.FieldCreatedAt))
	case query.PaymentOrderAmountAsc:
		qb = qb.Order(ent.Asc(payment.FieldAmount))
	case query.PaymentOrderAmountDesc:
		qb = qb.Order(ent.Desc(payment.FieldAmount))
	default:
		qb = qb.Order(ent.Desc(payment.FieldCreatedAt))
	}

	// Total count for pagination
	total, err := qb.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	// Fetch paginated data
	ents, err := qb.
		Limit(q.Limit).
		Offset(q.Offset()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	items := mapper.FromEntList(ents)

	return &pagination.PageData[domain.Payment]{
		Data:  items,
		Total: total,
	}, nil
}

// -------------------------
// Helpers
// -------------------------

// applyScope ensures tenant-level filtering.
func (r *EntPaymentRepository) applyScope(ctx *appctx.Context, qb *ent.PaymentQuery) *ent.PaymentQuery {
	switch ctx.Scope() {
	case appctx.ScopeTenant:
		qb = qb.Where(payment.TenantIDEQ(*ctx.TenantID()))
	case appctx.ScopeUser:
		fmt.Println("scope user")
		qb = qb.Where(payment.UserIDEQ(*ctx.UserID()))
	case appctx.ScopeAdmin:
		// no filtering for admin
	}

	return qb
}
