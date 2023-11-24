package ops

import "iter"

type SliceOps[Slice ~[]T, T any] interface {
	Filter(func(v T) bool) SliceOps[Slice, T]
	Sort(func(a, b T) bool) SliceOps[Slice, T]
	Map(func(v T) T) SliceOps[Slice, T]
	Result() Slice
}

type IteratorOps[T any] interface {
	iter.Iterator[T]
	Filter(func(v T) bool) IteratorOps[T]
	Map(func(v T) T) IteratorOps[T]
}
