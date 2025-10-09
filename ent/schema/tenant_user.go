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

// TenantUser holds the schema definition for the TenantUser entity.
type TenantUser struct {
	ent.Schema
}

// Fields of the TenantUser.
func (TenantUser) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("user_id", uuid.UUID{}).Immutable(),
		field.UUID("tenant_id", uuid.UUID{}).Immutable(),
		field.Enum("status").
			Values(string(types.StatusActive), string(types.StatusSuspended), string(types.StatusDeleted)).
			Default("active").Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the TenantUser.
func (TenantUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("tenant_users").
			Field("user_id").
			Unique().
			Required().
			Immutable(),

		edge.From("tenant", Tenant.Type).
			Ref("tenant_users").
			Field("tenant_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Indexes of the TenantUser.
func (TenantUser) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "tenant_id").Unique(),
	}
}
