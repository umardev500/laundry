package query

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/pagination"
)

// AddressOrder defines sorting options for listing addresses.
type AddressOrder string

const (
	AddressOrderCreatedAtAsc  AddressOrder = "created_at_asc"
	AddressOrderCreatedAtDesc AddressOrder = "created_at_desc"
	AddressOrderUpdatedAtAsc  AddressOrder = "updated_at_asc"
	AddressOrderUpdatedAtDesc AddressOrder = "updated_at_desc"
)

// ListAddressQuery represents query parameters for listing addresses.
type ListAddressQuery struct {
	pagination.Query
	UserID         uuid.UUID    `query:"user_id"`         // filter by user
	ProvinceID     string       `query:"province_id"`     // filter by province
	RegencyID      string       `query:"regency_id"`      // filter by regency
	DistrictID     string       `query:"district_id"`     // filter by district
	VillageID      string       `query:"village_id"`      // filter by village
	IncludeDeleted bool         `query:"include_deleted"` // include soft-deleted addresses
	IsPrimary      *bool        `query:"is_primary"`      // filter by primary address flag (nil = all)
	Order          AddressOrder `query:"order"`           // sorting

	// Eager loading
	WithUser     bool `query:"with_user"`
	WithProvince bool `query:"with_province"`
	WithRegency  bool `query:"with_regency"`
	WithDistrict bool `query:"with_district"`
	WithVillage  bool `query:"with_village"`
}

// Normalize ensures defaults are set.
func (q *ListAddressQuery) Normalize() {
	q.Query.Normalize(1, 10) // default page=1, limit=10
	if q.Order == "" {
		q.Order = AddressOrderCreatedAtDesc
	}
}
