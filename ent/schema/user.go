package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("email").Unique().NotEmpty(),
		field.String("password").Sensitive().NotEmpty(),
		field.Enum("status").Values("active", "suspended", "deleted").Default("active").Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("platform_users", PlatformUser.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),

		edge.To("tenant_users", TenantUser.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),

		edge.To("orders", Order.Type).
			Annotations(
				entsql.OnDelete(entsql.SetNull),
			),
	}
}
