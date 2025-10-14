package domain

import "errors"

var (
	ErrStatusHistoryNotFound = errors.New("order status history not found")
	ErrStatusHistoryDeleted  = errors.New("order status history deleted")
	ErrUnauthorizedAccess    = errors.New("unauthorized access")
)
