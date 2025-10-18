package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Addresses holds the schema definition for the Addresses entity.
type Addresses struct {
	ent.Schema
}

// Fields of the Addresses.
func (Addresses) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.String("province_id").MaxLen(2).NotEmpty(),
		field.String("regency_id").MaxLen(4).NotEmpty(),
		field.String("district_id").MaxLen(6).NotEmpty(),
		field.String("village_id").MaxLen(10).NotEmpty(),
		field.String("street").Optional().Nillable(),

		field.Bool("is_primary").Default(false),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Addresses.
func (Addresses) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("addresses").
			Required(),

		edge.From("province", Province.Type).
			Ref("addresses").
			Field("province_id").
			Unique().
			Required(),

		edge.From("regency", Regency.Type).
			Ref("addresses").
			Field("regency_id").
			Unique().
			Required(),

		edge.From("district", District.Type).
			Ref("addresses").
			Field("district_id").
			Unique().
			Required(),

		edge.From("village", Village.Type).
			Ref("addresses").
			Field("village_id").
			Unique().
			Required(),
	}
}
