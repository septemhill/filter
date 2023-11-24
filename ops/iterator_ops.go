package ops

import "iter"

type iterOpsImpl[T any] struct {
	next func() (T, bool)
}

func (it *iterOpsImpl[T]) Filter(fn func(v T) bool) IteratorOps[T] {
	return &iterOpsImpl[T]{
		next: func() (T, bool) {
			for e, ok := it.next(); ok; e, ok = it.next() {
				if fn(e) {
					return e, ok
				}
			}

			var e T
			return e, false
		},
	}
}

func (it *iterOpsImpl[T]) Map(fn func(v T) T) IteratorOps[T] {
	return &iterOpsImpl[T]{
		next: func() (T, bool) {
			var e T
			ok := false

			e, ok = it.next()
			if ok {
				return fn(e), ok
			}

			return e, ok
		},
	}
}

func (it *iterOpsImpl[T]) Next() (T, bool) {
	return it.next()
}

func FromIter[T any](v iter.Iterator[T]) *iterOpsImpl[T] {
	return &iterOpsImpl[T]{
		next: func() (T, bool) {
			return v.Next()
		},
	}
}
