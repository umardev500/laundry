package pagination

type Query struct {
	Page  int `query:"page" json:"page"`
	Limit int `query:"limit" json:"limit"`
}

func (q *Query) Normalize(defaultPage, defaultLimit int) {
	if q.Page <= 0 {
		q.Page = defaultPage
	}
	if q.Limit <= 0 {
		q.Limit = defaultLimit
	}
}

func (q *Query) Offset() int {
	return (q.Page - 1) * q.Limit
}
