package types

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
	OrderStatusPreview          OrderStatus = "PREVIEW"            // Order preview before confirmation
)

var AllowedOrderTransitions = map[OrderStatus][]OrderStatus{
	// Initial states
	OrderStatusPreview: {OrderStatusPending, OrderStatusCancelled},
	OrderStatusPending: {OrderStatusConfirmed, OrderStatusCancelled, OrderStatusFailed},

	// Confirmed order can move to next steps
	OrderStatusConfirmed: {
		OrderStatusPickedUp,
		OrderStatusCancelled,
		OrderStatusFailed,
	},

	OrderStatusPickedUp: {
		OrderStatusInWashing,
		OrderStatusFailed,
	},

	OrderStatusInWashing: {
		OrderStatusInDrying,
		OrderStatusFailed,
	},

	OrderStatusInDrying: {
		OrderStatusInIroning,
		OrderStatusFailed,
	},

	OrderStatusInIroning: {
		OrderStatusReadyForDelivery,
		OrderStatusFailed,
	},

	OrderStatusReadyForDelivery: {
		OrderStatusOutForDelivery,
		OrderStatusFailed,
	},

	OrderStatusOutForDelivery: {
		OrderStatusDelivered,
		OrderStatusFailed,
	},

	OrderStatusDelivered: {
		OrderStatusCompleted,
		OrderStatusFailed,
	},

	// Terminal states
	OrderStatusCompleted: {},
	OrderStatusCancelled: {},
	OrderStatusFailed:    {},
}
