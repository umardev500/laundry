package domain

import "fmt"

var (
	// Subscription errors
	ErrSubscriptionNotFound           = fmt.Errorf("subscription not found")
	ErrSubscriptionDeleted            = fmt.Errorf("subscription has been deleted")
	ErrUnauthorizedSubscriptionAccess = fmt.Errorf("unauthorized access to subscription")
	ErrSubscriptionNotDeleted         = fmt.Errorf("subscription is not deleted")
)
