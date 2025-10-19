package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/addresses"
	"github.com/umardev500/laundry/ent/user"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/address/domain"
	"github.com/umardev500/laundry/internal/feature/address/mapper"
	"github.com/umardev500/laundry/internal/feature/address/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

// Ensure implementation satisfies the interface.
var _ Repository = (*entAddressRepository)(nil)

type entAddressRepository struct {
	client *entdb.Client
}

// NewEntRepository creates a new Ent-based address repository.
func NewEntRepository(client *entdb.Client) Repository {
	return &entAddressRepository{client: client}
}

// Create inserts a new address.
func (r *entAddressRepository) Create(ctx *appctx.Context, a *domain.Address) (*domain.Address, error) {
	conn := r.client.GetConn(ctx)
	model, err := conn.Addresses.Create().
		SetID(uuid.New()).
		SetUserID(a.User.ID).
		SetProvinceID(a.Province.ID).
		SetRegencyID(a.Regency.ID).
		SetDistrictID(a.District.ID).
		SetVillageID(a.Village.ID).
		SetNillableStreet(a.Street).
		SetIsPrimary(a.IsPrimary).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEnt(model), nil
}

// FindByID retrieves an address by its ID with optional preloads.
func (r *entAddressRepository) FindByID(ctx *appctx.Context, id uuid.UUID, q *query.FindAddressByIDQuery) (*domain.Address, error) {
	conn := r.client.GetConn(ctx)
	builder := conn.Addresses.Query().Where(addresses.IDEQ(id))

	if q == nil {
		q = &query.FindAddressByIDQuery{}
	}

	if q.WithUser {
		builder.WithUser()
	}
	if q.WithProvince {
		builder.WithProvince()
	}
	if q.WithRegency {
		builder.WithRegency()
	}
	if q.WithDistrict {
		builder.WithDistrict()
	}
	if q.WithVillage {
		builder.WithVillage()
	}

	model, err := builder.Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEnt(model), nil
}

// FindPrimaryByUserID retrieves the primary address for a given user.
func (r *entAddressRepository) FindPrimaryByUserID(ctx *appctx.Context, userID uuid.UUID) (*domain.Address, error) {
	conn := r.client.GetConn(ctx)
	model, err := conn.Addresses.
		Query().
		Where(
			addresses.HasUserWith(user.IDEQ(userID)),
			addresses.IsPrimaryEQ(true),
			addresses.DeletedAtIsNil(),
		).
		Only(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEnt(model), nil
}

// UnsetPrimary marks all addresses of a user as non-primary.
func (r *entAddressRepository) UnsetPrimary(ctx *appctx.Context, userID uuid.UUID) error {
	conn := r.client.GetConn(ctx)
	_, err := conn.Addresses.
		Update().
		Where(
			addresses.HasUserWith(user.IDEQ(userID)),
			addresses.IsPrimary(true),
		).
		SetIsPrimary(false).
		Save(ctx)
	return err
}

// Update modifies an existing address record.
func (r *entAddressRepository) Update(ctx *appctx.Context, a *domain.Address) (*domain.Address, error) {
	conn := r.client.GetConn(ctx)

	model, err := conn.Addresses.UpdateOneID(a.ID).
		SetProvinceID(a.Province.ID).
		SetRegencyID(a.Regency.ID).
		SetDistrictID(a.District.ID).
		SetVillageID(a.Village.ID).
		SetNillableStreet(a.Street).
		SetIsPrimary(a.IsPrimary).
		SetNillableDeletedAt(a.DeletedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromEnt(model), nil
}

// Delete hard-deletes an address by ID.
func (r *entAddressRepository) Delete(ctx *appctx.Context, id uuid.UUID) error {
	conn := r.client.GetConn(ctx)

	err := conn.Addresses.DeleteOneID(id).
		Exec(ctx)
	return err
}

// List returns a paginated list of addresses based on filters.
func (r *entAddressRepository) List(ctx *appctx.Context, q *query.ListAddressQuery) (*pagination.PageData[domain.Address], error) {
	q.Normalize()

	conn := r.client.GetConn(ctx)

	builder := conn.Addresses.Query()

	if q.WithUser {
		builder.WithUser()
	}
	if q.WithProvince {
		builder.WithProvince()
	}
	if q.WithRegency {
		builder.WithRegency()
	}
	if q.WithDistrict {
		builder.WithDistrict()
	}
	if q.WithVillage {
		builder.WithVillage()
	}

	// Filters
	if q.UserID != uuid.Nil {
		builder = builder.Where(addresses.HasUserWith(user.IDEQ(q.UserID)))
	}
	if q.ProvinceID != "" {
		builder = builder.Where(addresses.ProvinceIDEQ(q.ProvinceID))
	}
	if q.RegencyID != "" {
		builder = builder.Where(addresses.RegencyIDEQ(q.RegencyID))
	}
	if q.DistrictID != "" {
		builder = builder.Where(addresses.DistrictIDEQ(q.DistrictID))
	}
	if q.VillageID != "" {
		builder = builder.Where(addresses.VillageIDEQ(q.VillageID))
	}
	if q.IsPrimary != nil {
		builder = builder.Where(addresses.IsPrimaryEQ(*q.IsPrimary))
	}
	if !q.IncludeDeleted {
		builder = builder.Where(addresses.DeletedAtIsNil())
	}

	// Sorting
	switch q.Order {
	case query.AddressOrderCreatedAtAsc:
		builder = builder.Order(ent.Asc(addresses.FieldCreatedAt))
	case query.AddressOrderCreatedAtDesc:
		builder = builder.Order(ent.Desc(addresses.FieldCreatedAt))
	case query.AddressOrderUpdatedAtAsc:
		builder = builder.Order(ent.Asc(addresses.FieldUpdatedAt))
	case query.AddressOrderUpdatedAtDesc:
		builder = builder.Order(ent.Desc(addresses.FieldUpdatedAt))
	default:
		builder = builder.Order(ent.Desc(addresses.FieldCreatedAt))
	}

	// Pagination
	total, err := builder.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	models, err := builder.
		Offset((q.Page - 1) * q.Limit).
		Limit(q.Limit).
		All(ctx)
	if err != nil {
		return nil, err
	}

	domainList := mapper.FromEntList(models)

	return pagination.NewPageData(domainList, total), nil
}
