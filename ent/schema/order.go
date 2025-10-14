package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("tenant_id", uuid.UUID{}).Immutable(),
		field.UUID("user_id", uuid.UUID{}).Immutable().Optional().Nillable().
			Comment("Optional when user is guest"),

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
			).
			Default(string(types.OrderStatusPending)),

		field.Float("total_amount").Default(0.0),
		field.String("notes").Optional().Nillable(),

		// Guest info (only if user_id is null)
		field.String("guest_name").Optional().Nillable(),
		field.String("guest_email").Optional().Nillable(),
		field.String("guest_phone").Optional().Nillable(),
		field.String("guest_address").Optional().Nillable(),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tenant", Tenant.Type).
			Ref("orders").
			Field("tenant_id").
			Immutable().
			Unique().
			Required(),

		edge.From("user", User.Type).
			Ref("orders").
			Field("user_id").
			Immutable().
			Unique(),

		edge.To("items", OrderItem.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),

		edge.To("payment", Payment.Type).
			Unique().
			Annotations(
				entsql.OnDelete(entsql.Restrict),
			),

		edge.To("status_history", OrderStatus.Type).
			Annotations(
				entsql.OnDelete(entsql.Cascade),
			),
	}
}
