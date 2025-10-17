package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Province holds the schema definition for the Province entity.
type Province struct {
	ent.Schema
}

// Fields of the Province.
func (Province) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(2).NotEmpty().Unique(),
		field.String("name").NotEmpty(),
	}
}

// Edges of the Province.
func (Province) Edges() []ent.Edge {
	return []ent.Edge{
		// One-to-many: Province has many Regencies
		edge.To("regencies", Regency.Type),

		edge.To("addresses", Addresses.Type),
	}
}
