package query

import (
	"time"

	"github.com/umardev500/laundry/pkg/pagination"
	"github.com/umardev500/laundry/pkg/types"
)

// SubscriptionOrder defines sorting options for listing subscriptions.
type SubscriptionOrder string

const (
	SubscriptionOrderStartDateAsc  SubscriptionOrder = "start_date_asc"
	SubscriptionOrderStartDateDesc SubscriptionOrder = "start_date_desc"
	SubscriptionOrderEndDateAsc    SubscriptionOrder = "end_date_asc"
	SubscriptionOrderEndDateDesc   SubscriptionOrder = "end_date_desc"
	SubscriptionOrderCreatedAtAsc  SubscriptionOrder = "created_at_asc"
	SubscriptionOrderCreatedAtDesc SubscriptionOrder = "created_at_desc"
	SubscriptionOrderUpdatedAtAsc  SubscriptionOrder = "updated_at_asc"
	SubscriptionOrderUpdatedAtDesc SubscriptionOrder = "updated_at_desc"
)

// ListSubscriptionQuery represents filters and sorting for listing subscriptions.
type ListSubscriptionQuery struct {
	pagination.Query

	TenantID       string                    `query:"tenant_id"`       // filter by tenant
	PlanID         string                    `query:"plan_id"`         // filter by plan
	Status         *types.SubscriptionStatus `query:"status"`          // filter by specific status
	ActiveOnly     *bool                     `query:"active_only"`     // if true, only ACTIVE subscriptions
	IncludeDeleted bool                      `query:"include_deleted"` // include soft-deleted subscriptions
	StartDateFrom  *time.Time                `query:"start_date_from"` // filter by start date range
	StartDateTo    *time.Time                `query:"start_date_to"`   // filter by start date range
	EndDateFrom    *time.Time                `query:"end_date_from"`   // filter by end date range
	EndDateTo      *time.Time                `query:"end_date_to"`     // filter by end date range

	Order SubscriptionOrder `query:"order"` // sorting
}

// Normalize ensures defaults are set for pagination and sorting.
func (q *ListSubscriptionQuery) Normalize() {
	q.Query.Normalize(1, 10) // default: page=1, limit=10

	if q.Order == "" {
		q.Order = SubscriptionOrderCreatedAtDesc
	}
}
