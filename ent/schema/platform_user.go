package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// PlatformUser holds the schema definition for the PlatformUser entity.
type PlatformUser struct {
	ent.Schema
}

// Fields of the PlatformUser.
func (PlatformUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("user_id", uuid.UUID{}).Immutable().Unique(),
		field.Enum("status").
			Values(string(types.PlatformUserStatusActive), string(types.PlatformUserStatusSuspended), string(types.PlatformUserStatusDeleted)).
			Default(string(types.PlatformUserStatusActive)).Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the PlatformUser.
func (PlatformUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("platform_users").
			Field("user_id").
			Unique().
			Required().
			Immutable(),

		edge.From("role", Role.Type).
			Ref("platform_users").
			Required().
			Immutable(),
	}
}

func (PlatformUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id").Unique(),
	}
}
