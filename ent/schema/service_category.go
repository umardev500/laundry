package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ServiceCategory holds the schema definition for the ServiceCategory entity.
type ServiceCategory struct {
	ent.Schema
}

// Fields of the ServiceCategory.
func (ServiceCategory) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("tenant_id", uuid.UUID{}).Immutable(),

		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the ServiceCategory.
func (ServiceCategory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("services", Service.Type).
			Annotations(
				entsql.OnDelete(entsql.SetNull),
			),
		edge.From("tenant", Tenant.Type).
			Ref("service_categories").
			Field("tenant_id").
			Unique().
			Required().
			Immutable(),
	}
}
