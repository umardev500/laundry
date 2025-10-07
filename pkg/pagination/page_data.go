package pagination

type PageData[T any] struct {
	Data  []*T
	Total int
}
