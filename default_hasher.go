package immutable

type number interface {
	~int | ~int64 | ~int32 | ~int16 | ~int8 | ~uint | ~uint64 | ~uint32 | ~uint16 | ~uint8 | ~float32 | ~float64
}

type hasherImp[T number] struct{}

func (i hasherImp[T]) hasher() Hasher[T] {
	return &NumberHasher[T]{}
}

type NumberHasher[T number] struct{}

func (h NumberHasher[T]) Hash(key T) uint32 {
	return hashUint64(uint64(key))
}

func (h NumberHasher[T]) Equal(a, b T) bool {
	return a == b
}

var _ Hasher[string] = &stringHasher[string]{}
var _ Hasher[[]byte] = &stringHasher[[]byte]{}

type stringLike interface {
	~string | ~[]byte
}

type stringHasher[T stringLike] struct{}

// Hash returns a hash for value.
func (h *stringHasher[T]) Hash(value T) uint32 {
	var hash uint32
	for i := 0; i < len(value); i++ {
		hash = 31*hash + uint32(value[i])
	}
	return hash
}

// Equal returns true if a is equal to b. Otherwise returns false.
// Panics if a and b are not byte slices.
func (h *stringHasher[T]) Equal(a, b T) bool {
	return string(a) == string(b)
}
