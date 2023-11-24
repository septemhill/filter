package filter

import "sort"

type Filter[Slice ~[]T, T any] struct {
	slice Slice
}

func (l *Filter[Slice, T]) Where(fn func(v T) bool) *Filter[Slice, T] {
	result := make([]T, 0)

	for _, e := range l.slice {
		if fn(e) {
			result = append(result, e)
		}
	}

	return &Filter[Slice, T]{
		slice: result,
	}
}

func (l *Filter[Slice, T]) Take(num int) *Filter[Slice, T] {
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}

	n := min(num, len(l.slice))

	return &Filter[Slice, T]{
		slice: l.slice[:n],
	}
}

func (l *Filter[Slice, T]) Sort(fn func(a, b T) bool) *Filter[Slice, T] {
	n := make([]T, len(l.slice))
	copy(n, l.slice)

	sort.Slice(n, func(i, j int) bool {
		return fn(n[i], n[j])
	})

	return &Filter[Slice, T]{
		slice: n,
	}
}

func (l *Filter[Slice, T]) Result() Slice {
	return l.slice
}

func From[Slice ~[]T, T any](slice Slice) *Filter[Slice, T] {
	return &Filter[Slice, T]{
		slice: slice,
	}
}
