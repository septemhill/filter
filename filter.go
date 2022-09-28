package filter

import "sort"

type Filter[Slice ~[]T, T any] interface {
	Where(func(v T) bool) Filter[Slice, T]
	Sort(func(a, b T) bool) Filter[Slice, T]
	Result() Slice
}

type filterImpl[Slice ~[]T, T any] struct {
	slice Slice
}

func (l *filterImpl[Slice, T]) Where(fn func(v T) bool) Filter[Slice, T] {
	result := make([]T, 0)

	for _, e := range l.slice {
		if fn(e) {
			result = append(result, e)
		}
	}

	return &filterImpl[Slice, T]{
		slice: result,
	}
}

func (l *filterImpl[Slice, T]) Sort(fn func(a, b T) bool) Filter[Slice, T] {
	n := make([]T, len(l.slice))
	copy(n, l.slice)

	sort.Slice(n, func(i, j int) bool {
		return fn(n[i], n[j])
	})

	return &filterImpl[Slice, T]{
		slice: n,
	}
}

func (l *filterImpl[Slice, T]) Result() Slice {
	return l.slice
}

func From[Slice ~[]T, T any](slice Slice) Filter[Slice, T] {
	return &filterImpl[Slice, T]{
		slice: slice,
	}
}
