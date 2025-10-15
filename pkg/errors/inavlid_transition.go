package errors

import (
	"errors"
	"fmt"
)

// ErrInvalidStatusTransition represents an invalid order status change.
type ErrInvalidStatusTransition[T any] struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Allowed []T    `json:"allowed"`
}

func (e *ErrInvalidStatusTransition[T]) Error() string {
	if len(e.Allowed) > 0 {
		return fmt.Sprintf(
			"invalid status transition from %s → %s",
			e.From, e.To,
		)
	}
	return fmt.Sprintf("invalid status transition from %s → %s", e.From, e.To)
}

// NewErrInvalidStatusTransition creates a detailed invalid transition error.
func NewErrInvalidStatusTransition[T any](from, to string, allowed []T) error {
	return &ErrInvalidStatusTransition[T]{
		From:    from,
		To:      to,
		Allowed: allowed,
	}
}

// IsInvalidTransition checks if the given error is an ErrInvalidStatusTransition for any type T.
// Returns true if so.
func IsInvalidTransition[T any](err error) (*ErrInvalidStatusTransition[T], bool) {
	var e *ErrInvalidStatusTransition[T]
	if errors.As(err, &e) {
		return e, true
	}
	return nil, false
}

func IsInvalidTransitionErr[T any](err error) bool {
	_, ok := IsInvalidTransition[T](err)
	return ok
}
