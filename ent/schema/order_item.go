package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// OrderItem holds the schema definition for the OrderItem entity.
type OrderItem struct {
	ent.Schema
}

// Fields of the OrderItem.
func (OrderItem) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("order_id", uuid.UUID{}).Immutable(),
		field.UUID("service_id", uuid.UUID{}).Immutable(),
		field.Float("quantity").Default(1),
		field.Float("price").Default(0.0),
		field.Float("subtotal").Default(0.0),
		field.Float("total_amount").Default(0.0),
	}
}

// Edges of the OrderItem.
func (OrderItem) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("items").
			Field("order_id").
			Unique().
			Immutable().
			Required(),

		edge.From("service", Service.Type).
			Ref("items").
			Field("service_id").
			Immutable().
			Unique().
			Required(),
	}
}
