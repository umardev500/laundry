package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// Permission holds the schema definition for the Permission entity.
type Permission struct {
	ent.Schema
}

// Fields of the Permission.
func (Permission) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("name").NotEmpty().Comment("Name of permission e.g. create_order"),
		field.String("display_name").NotEmpty().Comment("Display name of permission e.g. Create Order"),
		field.String("description").Optional(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Enum("status").
			Values(string(types.StatusActive), string(types.StatusSuspended), string(types.StatusDeleted)).
			Default(string(types.StatusActive)),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Permission.
func (Permission) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("feature", Feature.Type).
			Ref("permissions").
			Unique(),
	}
}
