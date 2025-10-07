package query

import (
	"fmt"

	"github.com/google/uuid"
)

// GetByIDQuery is a reusable query struct for extracting a single "id" parameter from Fiber routes.
type GetByIDQuery struct {
	ID string `params:"id"`
}

func (q *GetByIDQuery) UUID() (uuid.UUID, error) {
	if q.ID == "" {
		return uuid.Nil, fmt.Errorf("id is required")
	}

	uid, err := uuid.Parse(q.ID)
	if err != nil {
		return uuid.Nil, err
	}

	return uid, nil
}
