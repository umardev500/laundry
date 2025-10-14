package types

import (
	"slices"
)

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
	PaymentStatusPending         PaymentStatus = "pending"          // waiting for payment
	PaymentStatusProcessing      PaymentStatus = "processing"       // payment is being processed
	PaymentStatusPaid            PaymentStatus = "paid"             // payment completed successfully
	PaymentStatusFailed          PaymentStatus = "failed"           // payment failed
	PaymentStatusCancelled       PaymentStatus = "cancelled"        // payment cancelled before completion
	PaymentStatusRefundRequested PaymentStatus = "refund_requested" // refund requested, awaiting processing
	PaymentStatusRefunded        PaymentStatus = "refunded"         // refund completed
)

// AllowedPaymentTransitions defines valid payment state changes.
var AllowedPaymentTransitions = map[PaymentStatus][]PaymentStatus{
	// ---- Payment creation & processing ----
	PaymentStatusPending: {
		PaymentStatusProcessing, // start processing payment
		PaymentStatusCancelled,  // cancelled before processing
		PaymentStatusFailed,     // failed before completion
	},

	PaymentStatusProcessing: {
		PaymentStatusPaid,   // successfully processed
		PaymentStatusFailed, // failed during processing
	},

	// ---- After payment is successful ----
	PaymentStatusPaid: {
		PaymentStatusRefundRequested, // refund initiated by user/admin
	},

	// ---- Refund lifecycle ----
	PaymentStatusRefundRequested: {
		PaymentStatusRefunded, // refund completed
		PaymentStatusFailed,   // refund failed
	},

	// ---- Allow retry or reattempts ----
	PaymentStatusFailed: {
		PaymentStatusPending, // can retry payment
	},

	// ---- Terminal states ----
	PaymentStatusCancelled: {}, // cannot change after cancelled
	PaymentStatusRefunded:  {}, // refund done — closed
}

// CanTransition checks if a payment can move from current → next.
func (s PaymentStatus) CanTransitionTo(next PaymentStatus) bool {
	allowedNext, ok := AllowedPaymentTransitions[s]
	if !ok {
		return false
	}
	return slices.Contains(allowedNext, next)
}

func (s PaymentStatus) AllowedNextStatuses() []PaymentStatus {
	return AllowedPaymentTransitions[s]
}

// MapPaymentToOrderStatus maps payment status → appropriate order status
func MapPaymentToOrderStatus(paymentStatus PaymentStatus, current OrderStatus) OrderStatus {
	switch paymentStatus {
	case PaymentStatusPending, PaymentStatusProcessing:
		// Order placed but not yet paid
		if current == OrderStatusPreview {
			return OrderStatusPending
		}
		return current

	case PaymentStatusPaid:
		// Once paid, confirm the order (unless already in progress)
		if current == OrderStatusPending || current == OrderStatusPreview {
			return OrderStatusConfirmed
		}
		return current

	case PaymentStatusRefundRequested:
		// Customer or admin requested refund
		return OrderStatusRefundRequested

	case PaymentStatusRefunded:
		// Refund completed
		return OrderStatusRefunded

	case PaymentStatusFailed:
		// Payment failed → order failed
		return OrderStatusFailed

	case PaymentStatusCancelled:
		// Cancelled payment → cancelled order
		return OrderStatusCancelled

	default:
		return current
	}
}
