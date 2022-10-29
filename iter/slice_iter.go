package iter

func FromSlice[T any](v []T) Iterator[T] {
	return &sliceIter[T]{
		slice: v,
		index: -1,
	}
}

type sliceIter[T any] struct {
	slice []T
	index int
}

func (it *sliceIter[T]) Next() (T, bool) {
	var e T
	it.index++

	ok := it.index >= 0 && it.index < len(it.slice)
	if ok {
		e = it.slice[it.index]
	}

	return e, ok
}
