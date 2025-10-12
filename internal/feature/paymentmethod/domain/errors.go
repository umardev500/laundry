package domain

import "errors"

var (
	ErrPaymentMethodAlreadyExists = errors.New("payment method already exists")
	ErrPaymentMethodNotFound      = errors.New("payment method not found")
	ErrPaymentMethodDeleted       = errors.New("payment method has been deleted")
)
