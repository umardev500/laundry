package repository

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/ent"
	"github.com/umardev500/laundry/ent/feature"
	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/feature/rbac/domain"
	"github.com/umardev500/laundry/internal/feature/rbac/mapper"
	"github.com/umardev500/laundry/internal/feature/rbac/query"
	"github.com/umardev500/laundry/internal/infra/database/entdb"
	"github.com/umardev500/laundry/pkg/pagination"
)

type featureEntImpl struct {
	client *entdb.Client
}

func NewFeatureEntRepository(client *entdb.Client) FeatureRepository {
	return &featureEntImpl{client: client}
}

func (e *featureEntImpl) FindByID(ctx *appctx.Context, id uuid.UUID) (*domain.Feature, error) {
	conn := e.client.GetConn(ctx)
	entFeature, err := conn.Feature.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mapper.FromFeatureEntModel(entFeature), nil
}

func (e *featureEntImpl) List(ctx *appctx.Context, q *query.ListFeatureQuery) (*pagination.PageData[domain.Feature], error) {
	q.Normalize()
	conn := e.client.GetConn(ctx)
	queryBuilder := conn.Feature.Query()

	if q.Search != "" {
		queryBuilder = queryBuilder.Where(
			feature.Or(
				feature.NameContainsFold(q.Search),
				feature.DescriptionContainsFold(q.Search),
			),
		)
	}

	if !q.IncludeDeleted {
		queryBuilder = queryBuilder.Where(feature.DeletedAtIsNil())
	}

	switch q.Order {
	case query.OrderNameAsc:
		queryBuilder = queryBuilder.Order(ent.Asc(feature.FieldName))
	case query.OrderNameDesc:
		queryBuilder = queryBuilder.Order(ent.Desc(feature.FieldName))
	case query.OrderCreatedAtAsc:
		queryBuilder = queryBuilder.Order(ent.Asc(feature.FieldCreatedAt))
	case query.OrderCreatedAtDesc:
		queryBuilder = queryBuilder.Order(ent.Desc(feature.FieldCreatedAt))
	default:
		queryBuilder = queryBuilder.Order(ent.Asc(feature.FieldName))
	}

	total, err := queryBuilder.Clone().Count(ctx)
	if err != nil {
		return nil, err
	}

	entFeatures, err := queryBuilder.
		Limit(q.Limit).
		Offset(q.Offset()).
		All(ctx)
	if err != nil {
		return nil, err
	}

	features := mapper.FromFeatureEntModels(entFeatures)

	return &pagination.PageData[domain.Feature]{Data: features, Total: total}, nil
}

func (e *featureEntImpl) Update(ctx *appctx.Context, f *domain.Feature) (*domain.Feature, error) {
	conn := e.client.GetConn(ctx)
	entFeature, err := conn.Feature.
		UpdateOneID(f.ID).
		SetName(f.Name).
		SetDescription(f.Description).
		SetStatus(feature.Status(f.Status)).
		SetNillableDeletedAt(f.DeletedAt).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return mapper.FromFeatureEntModel(entFeature), nil
}
