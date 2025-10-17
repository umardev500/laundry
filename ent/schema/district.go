package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// District holds the schema definition for the District entity.
type District struct {
	ent.Schema
}

// Fields of the District.
func (District) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").MaxLen(6).NotEmpty().Unique(),
		field.String("regency_id").MaxLen(4).NotEmpty(),
		field.String("name").NotEmpty(),
	}
}

// Edges of the District.
func (District) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("regency", Regency.Type).
			Ref("districts").
			Field("regency_id").
			Unique().
			Required(),

		edge.To("villages", Village.Type),
	}
}
