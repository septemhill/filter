package ops

import "sort"

type SliceOpsImpl[Slice ~[]T, T any] struct {
	slice Slice
}

func (l *SliceOpsImpl[Slice, T]) Filter(fn func(v T) bool) SliceOps[Slice, T] {
	result := make([]T, 0)

	for _, e := range l.slice {
		if fn(e) {
			result = append(result, e)
		}
	}

	return &SliceOpsImpl[Slice, T]{
		slice: result,
	}
}

func (l *SliceOpsImpl[Slice, T]) Sort(fn func(a, b T) bool) SliceOps[Slice, T] {
	n := make([]T, len(l.slice))
	copy(n, l.slice)

	sort.Slice(n, func(i, j int) bool {
		return fn(n[i], n[j])
	})

	return &SliceOpsImpl[Slice, T]{
		slice: n,
	}
}

func (l *SliceOpsImpl[Slice, T]) Map(fn func(v T) T) SliceOps[Slice, T] {
	result := make([]T, len(l.slice))

	for i, e := range l.slice {
		result[i] = fn(e)
	}

	return &SliceOpsImpl[Slice, T]{
		slice: result,
	}
}

func (l *SliceOpsImpl[Slice, T]) Result() Slice {
	return l.slice
}

func From[Slice ~[]T, T any](slice Slice) SliceOps[Slice, T] {
	return &SliceOpsImpl[Slice, T]{
		slice: slice,
	}
}
