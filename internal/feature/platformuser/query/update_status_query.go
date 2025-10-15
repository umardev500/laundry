package query

import (
	"github.com/google/uuid"
	"github.com/umardev500/laundry/internal/feature/platformuser/domain"
	"github.com/umardev500/laundry/pkg/types"
)

type UpdateStatusQuery struct {
	Status string `params:"status"`
}

func (q *UpdateStatusQuery) ToDomain(uid uuid.UUID) *domain.PlatformUser {
	return &domain.PlatformUser{
		ID:     uid,
		Status: types.PlatformUserStatus(q.Status),
	}
}
