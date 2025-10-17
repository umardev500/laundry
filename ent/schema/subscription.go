package schema

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// Subscription holds the schema definition for the Subscription entity.
type Subscription struct {
	ent.Schema
}

// Fields of the Subscription.
func (Subscription) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),

		field.UUID("tenant_id", uuid.UUID{}).Immutable(),
		field.UUID("plan_id", uuid.UUID{}).Immutable(),

		field.Enum("status").
			Values(
				string(types.SubscriptionStatusActive),
				string(types.SubscriptionStatusExpired),
				string(types.SubscriptionStatusCanceled),
				string(types.SubscriptionStatusSuspended),
				string(types.SubscriptionStatusDeleted),
			).
			Default(string(types.SubscriptionStatusActive)),
		field.Time("start_date").Optional().Nillable().
			Comment("Date when the subscription starts"),
		field.Time("end_date").Optional().Nillable().
			Comment("Date when the subscription ends"),

		field.Time("created_at").Default(time.Now).Immutable(),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
		field.Time("deleted_at").Optional().Nillable(),
	}
}

// Edges of the Subscription.
func (Subscription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("subscription_events", SubscriptionEvent.Type).Annotations(
			entsql.OnDelete(entsql.Cascade),
		),

		edge.From("tenant", Tenant.Type).
			Ref("subscriptions").
			Field("tenant_id").
			Required().
			Unique().
			Immutable(),
	}
}

// Indexes of the Subscription.
func (Subscription) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tenant_id").
			Unique().
			Annotations(
				entsql.IndexWhere(fmt.Sprintf("status = '%s'", types.SubscriptionStatusActive)),
			),
	}
}
