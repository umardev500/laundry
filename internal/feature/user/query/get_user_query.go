package query

import (
	"fmt"

	"github.com/google/uuid"
)

type GetuserQuery struct {
	ID string `params:"id"`
}

func (q *GetuserQuery) UUID() (uuid.UUID, error) {
	if q.ID == "" {
		return uuid.Nil, fmt.Errorf("id is required")
	}

	uid, err := uuid.Parse(q.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return uid, nil
}
