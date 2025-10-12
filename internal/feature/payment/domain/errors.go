package domain

import "errors"

var (
	ErrPaymentDeleted            = errors.New("payment has been deleted")
	ErrPaymentNotFound           = errors.New("payment not found")
	ErrUnauthorizedPaymentAccess = errors.New("unauthorized access to payment")
)
