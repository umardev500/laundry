package query

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/user/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type UpdateStatusQuery struct {
	Status string `params:"status"`
}

func (q *UpdateStatusQuery) ToDomain(uid uuid.UUID) *domain.User {
	return &domain.User{
		ID:     uid,
		Status: types.UserStatus(q.Status),
	}
}
