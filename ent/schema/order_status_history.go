package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// OrderStatusHistory holds the schema definition for the OrderStatusHistory entity.
type OrderStatusHistory struct {
	ent.Schema
}

// Fields of the OrderStatus.
func (OrderStatusHistory) Fields() []ent.Field {
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
				string(types.OrderStatusRefundRequested),
				string(types.OrderStatusRefunded),
			),
		field.String("notes").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the OrderStatus.
func (OrderStatusHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("order", Order.Type).
			Ref("status_history").
			Field("order_id").
			Unique().
			Immutable().
			Required(),
	}
}
