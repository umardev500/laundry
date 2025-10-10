package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// ServiceUnit holds the schema definition for the ServiceUnit entity.
type ServiceUnit struct {
	ent.Schema
}

// Fields of the ServiceUnit.
func (ServiceUnit) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("tenant_id", uuid.UUID{}).Immutable(),
		field.String("name").NotEmpty().Unique().Comment("Full name of the unit, e.g. 'Per Piece', 'Per Kilogram'"),
		field.String("symbol").Optional().Comment("Short form like 'pc', 'kg', 'set'"),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the ServiceUnit.
func (ServiceUnit) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("services", Service.Type).
			Annotations(
				entsql.OnDelete(entsql.SetNull),
			),

		edge.From("tenant", Tenant.Type).
			Ref("service_units").
			Field("tenant_id").
			Unique().
			Required().
			Immutable(),
	}
}
