package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/umardev500/laundry/pkg/types"
)

// SubscriptionEvent holds the schema definition for the SubscriptionEvent entity.
type SubscriptionEvent struct {
	ent.Schema
}

// Fields of the SubscriptionEvent.
func (SubscriptionEvent) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New).Immutable(),

		field.UUID("subscription_id", uuid.UUID{}).Immutable(),
		field.Enum("event_type").
			Values(
				string(types.SubscriptionEventCreated),
				string(types.SubscriptionEventCanceled),
				string(types.SubscriptionEventExpired),
				string(types.SubscriptionEventRenewed),
				string(types.SubscriptionEventReactivated),
				string(types.SubscriptionEventStatusChanged),
				string(types.SubscriptionEventSuspended),
				string(types.SubscriptionEventDeleted),
			),
		field.Enum("old_status").
			Values(
				string(types.SubscriptionStatusActive),
				string(types.SubscriptionStatusCanceled),
				string(types.SubscriptionStatusExpired),
				string(types.SubscriptionStatusSuspended),
				string(types.SubscriptionStatusDeleted),
			).
			Optional().Nillable(),

		field.Enum("new_status").
			Values(
				string(types.SubscriptionStatusActive),
				string(types.SubscriptionStatusCanceled),
				string(types.SubscriptionStatusExpired),
				string(types.SubscriptionStatusSuspended),
				string(types.SubscriptionStatusDeleted),
			).
			Optional().Nillable(),
		field.JSON("data", map[string]any{}).Optional(),

		field.Enum("created_by_type").
			Values(
				string(types.CreatorTypeAdmin),
				string(types.CreatorTypeUser),
				string(types.CreatorTypeSystem),
			).
			Default(string(types.CreatorTypeSystem)),
		field.UUID("created_by", uuid.UUID{}).Optional().Nillable(),
		field.Time("created_at").Default(time.Now).Immutable(),
	}
}

// Edges of the SubscriptionEvent.
func (SubscriptionEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("subscription", Subscription.Type).
			Ref("subscription_events").
			Field("subscription_id").
			Immutable().
			Unique().
			Required(),

		edge.From("user", User.Type).
			Ref("subscription_events").
			Field("created_by").
			Unique(),
	}
}
