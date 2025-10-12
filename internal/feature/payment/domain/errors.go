package domain

import "errors"

var (
	ErrPaymentDeleted            = errors.New("payment has been deleted")
	ErrPaymentNotFound           = errors.New("payment not found")
	ErrUnauthorizedPaymentAccess = errors.New("unauthorized access to payment")
	ErrInvalidPaymentType        = errors.New("invalid payment type")
	ErrInvalidPaymentStatus      = errors.New("invalid payment status")
	ErrAlreadyPaid               = errors.New("payment is already paid")
	ErrOnlyPendingPayments       = errors.New("only pending payments can be marked as paid")
	ErrInsufficientPayment       = errors.New("insufficient payment")
	ErrInvalidAmount             = errors.New("amount must be greater than zero")
	ErrInvalidReceivedAmount     = errors.New("received amount cannot be negative")
	ErrReceivedAmountLessThanDue = errors.New("received amount cannot be less than the due amount")
	ErrInvalidChangeAmount       = errors.New("invalid change amount")
	ErrPaidAtWithoutPaidStatus   = errors.New("paid_at timestamp provided without 'paid' status")
	ErrPaidStatusWithoutPaidAt   = errors.New("'paid' status requires a paid_at timestamp")
)
