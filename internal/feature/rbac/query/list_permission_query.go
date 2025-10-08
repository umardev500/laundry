package query

import "github.com/umardev500/laundry/pkg/pagination"

// --- Order Enum ---
type PermissionOrder string

const (
	PermissionOrderNameAsc         PermissionOrder = "name_asc"
	PermissionOrderNameDesc        PermissionOrder = "name_desc"
	PermissionOrderDisplayNameAsc  PermissionOrder = "display_name_asc"
	PermissionOrderDisplayNameDesc PermissionOrder = "display_name_desc"
	PermissionOrderCreatedAtAsc    PermissionOrder = "created_at_asc"
	PermissionOrderCreatedAtDesc   PermissionOrder = "created_at_desc"
	PermissionOrderUpdatedAtAsc    PermissionOrder = "updated_at_asc"
	PermissionOrderUpdatedAtDesc   PermissionOrder = "updated_at_desc"
)

// --- Query Struct ---
type ListPermissionQuery struct {
	pagination.Query
	FeatureID      string          `query:"feature_id"`
	Search         string          `query:"search"`
	Status         string          `query:"status"`
	Order          PermissionOrder `query:"order"`
	IncludeDeleted bool            `query:"include_deleted"`
}

// --- Normalize ---
func (q *ListPermissionQuery) Normalize() {
	q.Query.Normalize(1, 10)

	if q.Order == "" {
		q.Order = PermissionOrderCreatedAtAsc
	}
}
