package query

import "github.com/umardev500/laundry/pkg/pagination"

// PlanOrder defines sorting options for listing plans.
type PlanOrder string

const (
	PlanOrderNameAsc       PlanOrder = "name_asc"
	PlanOrderNameDesc      PlanOrder = "name_desc"
	PlanOrderPriceAsc      PlanOrder = "price_asc"
	PlanOrderPriceDesc     PlanOrder = "price_desc"
	PlanOrderCreatedAtAsc  PlanOrder = "created_at_asc"
	PlanOrderCreatedAtDesc PlanOrder = "created_at_desc"
	PlanOrderUpdatedAtAsc  PlanOrder = "updated_at_asc"
	PlanOrderUpdatedAtDesc PlanOrder = "updated_at_desc"
)

// ListPlanQuery represents query parameters for listing plans.
type ListPlanQuery struct {
	pagination.Query
	Search         string    `query:"search"`          // search by name or description
	IncludeDeleted bool      `query:"include_deleted"` // include soft-deleted plans
	ActiveOnly     *bool     `query:"active_only"`     // filter by active plans, nil = all
	Order          PlanOrder `query:"order"`           // sorting
}

// Normalize ensures defaults are set.
func (q *ListPlanQuery) Normalize() {
	q.Query.Normalize(1, 10) // default page=1, limit=10
	if q.Order == "" {
		q.Order = PlanOrderCreatedAtDesc
	}
}
