package immutable

type orderable interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | ~float32 | ~float64 | ~string
}

var _ Comparer[int] = &DefaultComparer[int]{}

type DefaultComparer[T orderable] struct{}

func (c DefaultComparer[T]) Compare(x, y T) int {
	if x == y {
		return 0
	}
	if x < y {
		return -1
	}
	return 1
}
