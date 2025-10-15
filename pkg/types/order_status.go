package types

import (
	"slices"
	"strings"
)

type OrderStatus string

const (
	OrderStatusPending          OrderStatus = "PENDING"            // Order placed, awaiting confirmation
	OrderStatusConfirmed        OrderStatus = "CONFIRMED"          // Confirmed by laundry service
	OrderStatusPickedUp         OrderStatus = "PICKED_UP"          // Clothes collected from customer
	OrderStatusInWashing        OrderStatus = "IN_WASHING"         // Currently being washed
	OrderStatusInDrying         OrderStatus = "IN_DRYING"          // Clothes are being dried
	OrderStatusInIroning        OrderStatus = "IN_IRONING"         // Clothes are being ironed/folded
	OrderStatusReadyForDelivery OrderStatus = "READY_FOR_DELIVERY" // Laundry ready, waiting for delivery
	OrderStatusOutForDelivery   OrderStatus = "OUT_FOR_DELIVERY"   // Delivery person is on the way
	OrderStatusDelivered        OrderStatus = "DELIVERED"          // Clothes delivered to customer
	OrderStatusCompleted        OrderStatus = "COMPLETED"          // Order closed and payment settled
	OrderStatusCancelled        OrderStatus = "CANCELLED"          // Order cancelled before pickup
	OrderStatusFailed           OrderStatus = "FAILED"             // Failed due to error (e.g., payment issue)
	OrderStatusRefundRequested  OrderStatus = "REFUND_REQUESTED"   // Refund requested, waiting for approval
	OrderStatusRefunded         OrderStatus = "REFUNDED"           // Refund approved and payment reversed
	OrderStatusPreview          OrderStatus = "PREVIEW"            // Order preview before confirmation
)

var AllowedOrderTransitions = map[OrderStatus][]OrderStatus{
	OrderStatusPreview:          {OrderStatusPending, OrderStatusCancelled},
	OrderStatusPending:          {OrderStatusConfirmed, OrderStatusCancelled, OrderStatusFailed},
	OrderStatusConfirmed:        {OrderStatusPickedUp, OrderStatusCancelled, OrderStatusFailed},
	OrderStatusPickedUp:         {OrderStatusInWashing, OrderStatusFailed},
	OrderStatusInWashing:        {OrderStatusInDrying, OrderStatusFailed},
	OrderStatusInDrying:         {OrderStatusInIroning, OrderStatusFailed},
	OrderStatusInIroning:        {OrderStatusReadyForDelivery, OrderStatusFailed},
	OrderStatusReadyForDelivery: {OrderStatusOutForDelivery, OrderStatusFailed},
	OrderStatusOutForDelivery:   {OrderStatusDelivered, OrderStatusFailed},
	OrderStatusDelivered: {
		OrderStatusCompleted,
		OrderStatusRefundRequested, // allow refund request after delivery
	},
	OrderStatusCompleted: {
		OrderStatusRefundRequested, // allow refund request even after completion
	},
	OrderStatusCancelled: {
		OrderStatusRefundRequested, // allow refund request if payment was captured
	},
	OrderStatusFailed: {
		OrderStatusRefundRequested, // failed payment could still require refund
	},
	OrderStatusRefundRequested: {
		OrderStatusRefunded, // after processing
	},
	// ---- Terminal states ----
	OrderStatusRefunded: {},
}

func (s OrderStatus) CanTransitionTo(next OrderStatus) bool {
	nextNormalize := next.Normalize()
	allowedNext, ok := AllowedOrderTransitions[s]
	if !ok {
		return false
	}
	return slices.Contains(allowedNext, nextNormalize)
}

func (s OrderStatus) AllowedNextStatuses() []OrderStatus {
	return AllowedOrderTransitions[s]
}

func (e OrderStatus) Normalize() OrderStatus {
	return OrderStatus(strings.ToUpper(string(e)))
}

func MapOrderToPaymentStatus(orderStatus OrderStatus, current PaymentStatus) PaymentStatus {
	switch orderStatus {
	case OrderStatusConfirmed, OrderStatusCompleted:
		return PaymentStatusPaid

	case OrderStatusRefundRequested:
		return PaymentStatusRefundRequested

	case OrderStatusRefunded:
		return PaymentStatusRefunded

	case OrderStatusCancelled:
		// // If already paid, trigger refund request instead of direct cancel
		// if current == PaymentStatusPaid {
		// 	return PaymentStatusRefundRequested
		// }
		return PaymentStatusCancelled

	default:
		return current
	}
}
