package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Regency holds the schema definition for the Regency entity.
type Regency struct {
	ent.Schema
}

// Fields of the Regency.
func (Regency) Fields() []ent.Field {
	fields := []ent.Field{
		field.String("id").MaxLen(4).NotEmpty().Unique(),
		field.String("province_id").MaxLen(2).NotEmpty(),
		field.String("name").NotEmpty(),
	}
	return fields
}

// Edges of the Regency.
func (Regency) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("province", Province.Type).
			Ref("regencies").
			Field("province_id").
			Unique().
			Required(),

		edge.To("districts", District.Type),
	}
}
