package types

// SubscriptionEventType represents the type of a SubscriptionEvent
type SubscriptionEventType string

const (
	SubscriptionEventCreated       SubscriptionEventType = "CREATED"        // New subscription row created
	SubscriptionEventCanceled      SubscriptionEventType = "CANCELED"       // User/admin canceled
	SubscriptionEventExpired       SubscriptionEventType = "EXPIRED"        // Subscription naturally ended
	SubscriptionEventSuspended     SubscriptionEventType = "SUSPENDED"      // Temporarily disabled (e.g., payment failed)
	SubscriptionEventReactivated   SubscriptionEventType = "REACTIVATED"    // Suspended subscription became active again
	SubscriptionEventRenewed       SubscriptionEventType = "RENEWED"        // New subscription row created as renewal
	SubscriptionEventDeleted       SubscriptionEventType = "DELETED"        // Subscription logically deleted
	SubscriptionEventStatusChanged SubscriptionEventType = "STATUS_CHANGED" // Generic catch-all for unusual status changes
)

// SubscriptionEventCreatorType represents the creator of a SubscriptionEvent
type SubscriptionEventCreatorType string

const (
	CreatorTypeUser   SubscriptionEventCreatorType = "USER"
	CreatorTypeAdmin  SubscriptionEventCreatorType = "ADMIN"
	CreatorTypeSystem SubscriptionEventCreatorType = "SYSTEM"
)
