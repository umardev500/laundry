package query

// FindAddressByIDQuery represents options for finding a single address by ID.
type FindAddressByIDQuery struct {
	WithUser       bool `query:"with_user"`
	WithProvince   bool `query:"with_province"`
	WithRegency    bool `query:"with_regency"`
	WithDistrict   bool `query:"with_district"`
	WithVillage    bool `query:"with_village"`
	IncludeDeleted bool `query:"include_deleted"`
}
