package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Service holds the schema definition for the Service entity.
type Service struct {
	ent.Schema
}

// Fields of the Service.
func (Service) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("tenant_id", uuid.UUID{}).Immutable(),
		field.UUID("service_unit_id", uuid.UUID{}).Optional().Nillable(),
		field.UUID("service_category_id", uuid.UUID{}).Optional().Nillable(),

		field.String("name").NotEmpty().Unique(),
		field.Float("base_price").Default(0.0).
			Comment("Price of the service"),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Service.
func (Service) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("service_category", ServiceCategory.Type).
			Ref("services").
			Field("service_category_id").
			Unique(),

		edge.From("tenant", Tenant.Type).
			Ref("services").
			Field("tenant_id").
			Unique().
			Required().
			Immutable(),

		edge.From("unit", ServiceUnit.Type).
			Ref("services").
			Field("service_unit_id").
			Unique(),

		edge.To("items", OrderItem.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
	}
}
