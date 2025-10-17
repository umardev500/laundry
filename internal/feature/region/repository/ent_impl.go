package repository

import (
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/district"
	"github.com/umardev500/laundry/ent/province"
	"github.com/umardev500/laundry/ent/regency"
	"github.com/umardev500/laundry/ent/village"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/region/domain"
	"github.com/umardev500/laundry/internal/feature/region/mapper"
	"github.com/umardev500/laundry/internal/feature/region/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

type entImpl struct {
	client *entdb.Client
}

// FindVillages implements Repository.
func (e *entImpl) FindVillages(ctx *appctx.Context, districtID string, q *query.ListVillageQuery) (*pagination.PageData[domain.Village], error) {
	conn := e.client.GetConn(ctx)
	qb := conn.Village.
		Query().
		Where(village.DistrictID(districtID))

	if q.Search != "" {
		qb = qb.Where(
			village.NameContainsFold(q.Search),
		)
	}

	switch q.Order {
	case query.VillageOrderNameAsc:
		qb = qb.Order(ent.Asc(village.FieldName))
	case query.VillageOrderNameDesc:
		qb = qb.Order(ent.Desc(village.FieldName))
	}

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

	items := mapper.FromEntVillageList(ents)

	return &pagination.PageData[domain.Village]{
		Data:  items,
		Total: total,
	}, nil
}

// FindDistricts implements Repository.
func (e *entImpl) FindDistricts(ctx *appctx.Context, regencyID string, q *query.ListDistrictQuery) (*pagination.PageData[domain.District], error) {
	conn := e.client.GetConn(ctx)
	qb := conn.District.
		Query().
		Where(district.RegencyID(regencyID))

	if q.Search != "" {
		qb = qb.Where(
			district.NameContainsFold(q.Search),
		)
	}

	switch q.Order {
	case query.DistrictOrderNameAsc:
		qb = qb.Order(ent.Asc(district.FieldName))
	case query.DistrictOrderNameDesc:
		qb = qb.Order(ent.Desc(district.FieldName))
	}

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

	items := mapper.FromEntDistrictList(ents)

	return &pagination.PageData[domain.District]{
		Data:  items,
		Total: total,
	}, nil
}

func NewRepository(client *entdb.Client) Repository {
	return &entImpl{
		client: client,
	}
}

// FindRegencies implements Repository.
func (e *entImpl) FindRegencies(ctx *appctx.Context, provinceID string, q *query.ListRegencyQuery) (*pagination.PageData[domain.Regency], error) {
	conn := e.client.GetConn(ctx)
	qb := conn.Regency.
		Query().
		Where(regency.ProvinceID(provinceID))

	if q.Search != "" {
		qb = qb.Where(
			regency.NameContainsFold(q.Search),
		)
	}

	switch q.Order {
	case query.RegencyOrderNameAsc:
		qb = qb.Order(ent.Asc(regency.FieldName))
	case query.RegencyOrderNameDesc:
		qb = qb.Order(ent.Desc(regency.FieldName))
	}

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

	items := mapper.FromEntRegencyList(ents)

	return &pagination.PageData[domain.Regency]{
		Data:  items,
		Total: total,
	}, nil
}

// FindProvinces implements Repository.
func (e *entImpl) FindProvinces(ctx *appctx.Context, q *query.ListProvinceQuery) (*pagination.PageData[domain.Province], error) {
	conn := e.client.GetConn(ctx)
	qb := conn.Province.
		Query()

	if q.Search != "" {
		qb = qb.Where(
			province.NameContainsFold(q.Search),
		)
	}

	switch q.Order {
	case query.ProviceOrderNameAsc:
		qb = qb.Order(ent.Asc(province.FieldName))
	case query.ProviceOrderNameDesc:
		qb = qb.Order(ent.Desc(province.FieldName))
	}

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

	items := mapper.FromEntProvinceList(ents)

	return &pagination.PageData[domain.Province]{
		Data:  items,
		Total: total,
	}, nil
}
