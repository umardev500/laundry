package types

// PaymentType is the type of payment
type PaymentType string

const (
	PaymentTypeOrder        PaymentType = "order"
	PaymentTypeSubscription PaymentType = "subscription"
)

// PaymentMethod is the type of payment method
type PaymentMethod string

const (
	PaymentMethodCash     PaymentMethod = "cash"
	PaymentMethodCard     PaymentMethod = "card"
	PaymentMethodTransfer PaymentMethod = "transfer"
)

type PaymentStatus string

const (
	PaymentStatusPending PaymentStatus = "pending"
	PaymentStatusPaid    PaymentStatus = "paid"
	PaymentStatusFailed  PaymentStatus = "failed"
)
