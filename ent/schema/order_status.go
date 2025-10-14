package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// OrderStatus holds the schema definition for the OrderStatus entity.
type OrderStatus struct {
	ent.Schema
}

// Fields of the OrderStatus.
func (OrderStatus) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("order_id", uuid.UUID{}).Immutable().Comment("Reference to the order"),
		field.Enum("status").
			Values(
				string(types.OrderStatusPending),
				string(types.OrderStatusConfirmed),
				string(types.OrderStatusPickedUp),
				string(types.OrderStatusInWashing),
				string(types.OrderStatusInDrying),
				string(types.OrderStatusInIroning),
				string(types.OrderStatusReadyForDelivery),
				string(types.OrderStatusOutForDelivery),
				string(types.OrderStatusDelivered),
				string(types.OrderStatusCompleted),
				string(types.OrderStatusCancelled),
				string(types.OrderStatusFailed),
				string(types.OrderStatusPreview),
			),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.String("notes").Optional().Nillable(),
	}
}

// Edges of the OrderStatus.
func (OrderStatus) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("status_history").
			Field("order_id").
			Unique().
			Immutable().
			Required(),
	}
}
