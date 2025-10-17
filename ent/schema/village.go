package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Village holds the schema definition for the Village entity.
type Village struct {
	ent.Schema
}

// Fields of the Village.
func (Village) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(10).NotEmpty().Unique(),
		field.String("district_id").MaxLen(6).NotEmpty().Unique(),
		field.String("name").NotEmpty(),
	}
}

// Edges of the Village.
func (Village) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("district", District.Type).
			Ref("villages").
			Field("district_id").
			Unique().
			Required(),

		edge.To("addresses", Addresses.Type),
	}
}
