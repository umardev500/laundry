package query

// FindSubscriptionByIDQuery represents query parameters for finding a subscription by ID.
type FindSubscriptionByIDQuery struct {
	IncludeDeleted bool `query:"include_deleted"` // whether to include soft-deleted subscriptions
}
