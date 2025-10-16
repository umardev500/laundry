package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// Plan holds the schema definition for the Plan entity.
type Plan struct {
	ent.Schema
}

// Fields of the Plan.
func (Plan) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),

		field.String("name").NotEmpty().Unique().
			Comment("Name of the plan e.g. 'Basic', 'Premium'"),
		field.String("description").Optional().Nillable(),

		field.Float("price").Default(0),
		field.Enum("billing_interval").
			Values(
				string(types.BillingIntervalMonthly),
				string(types.BillingIntervalYearly),
			).
			Default(string(types.BillingIntervalMonthly)),

		field.JSON("features", map[string]any{}).Optional().
			Comment("Dynamic feature flags / limits per plan, e.g., max_users, storage_mb"),

		field.Bool("active").Default(true),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Plan.
func (Plan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("permissions", Permission.Type).
			Ref("plans"),
	}
}
