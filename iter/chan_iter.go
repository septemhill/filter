package iter

type channelIter[T any] struct {
	ch <-chan T
}

func FromChannel[T any](ch <-chan T) *channelIter[T] {
	return &channelIter[T]{
		ch: ch,
	}
}

func (it *channelIter[T]) Next() (T, bool) {
	e, ok := <-it.ch
	return e, ok
}
