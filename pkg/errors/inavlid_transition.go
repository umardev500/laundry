package errors

import "fmt"

// ErrInvalidStatusTransition represents an invalid order status change
type ErrInvalidStatusTransition struct {
	From string
	To   string
}

func (e *ErrInvalidStatusTransition) Error() string {
	return fmt.Sprintf("invalid status transition from %s to %s", e.From, e.To)
}

// Helper function to create the error
func NewErrInvalidStatusTransition(from, to string) error {
	return &ErrInvalidStatusTransition{
		From: from,
		To:   to,
	}
}
