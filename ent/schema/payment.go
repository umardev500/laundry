package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),
		field.UUID("ref_id", uuid.UUID{}).Immutable(),
		field.Enum("ref_type").
			Values(string(types.PaymentTypeOrder), string(types.PaymentTypeSubscription)),
		field.UUID("payment_method_id", uuid.UUID{}).Optional(),
		field.Float("amount").Default(0.0),
		field.String("notes").Optional(),
		field.Enum("status").
			Values(
				string(types.PaymentStatusPending),
				string(types.PaymentStatusPaid),
				string(types.PaymentStatusFailed),
			).
			Default(string(types.PaymentStatusPending)),

		field.Time("paid_at").Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("payment_method", PaymentMethod.Type).
			Ref("payments").
			Field("payment_method_id").
			Unique(),

		edge.From("order", Order.Type).
			Ref("payments").
			Unique(),
	}
}
