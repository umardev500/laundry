package httpx

import "math"

func NewPagination(page, limit, total int) *Pagination {
	totalPages := int(math.Ceil(float64(total) / float64(limit)))

	return &Pagination{
		Page:       page,
		PageSize:   limit,
		TotalItems: total,
		TotalPages: totalPages,
		HasNext:    page < totalPages,
		HasPrev:    page > 1,
	}
}
