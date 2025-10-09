package pagination

type PageData[T any] struct {
	Data  []*T
	Total int
}

func NewPageData[T any](data []*T, total int) *PageData[T] {
	return &PageData[T]{
		Data:  data,
		Total: total,
	}
}
