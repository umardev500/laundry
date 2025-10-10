package utils

import "github.com/google/uuid"

// NilIfUUIDZero returns nil if UUID is uuid.Nil
func NilIfUUIDZero(id uuid.UUID) *uuid.UUID {
	return NilIfZero(id, uuid.Nil)
}
