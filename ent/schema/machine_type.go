package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

type MachineType struct{ ent.Schema }

func (MachineType) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("tenant_id", uuid.UUID{}).Immutable(),
		field.String("name").NotEmpty().Unique(),
		field.String("description").Optional(),
		field.Int("capacity").Optional(),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

func (MachineType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("machines", Machine.Type).
			Annotations(
				entsql.OnDelete(entsql.SetNull),
			),

		edge.From("tenant", Tenant.Type).
			Ref("machine_types").
			Field("tenant_id").
			Required().
			Unique().
			Immutable(),
	}
}
