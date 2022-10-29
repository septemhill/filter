package future

type FutureFn[T any] func() (T, error)

type Future[T any, U Result[T]] struct {
	dat chan U
}

func (f *Future[T, U]) Await() U {
	return <-f.dat
}

func NewFuture[T any, U Result[T]](fn FutureFn[T]) *Future[T, U] {
	fut := &Future[T, U]{
		dat: make(chan U, 1),
	}

	go func() {
		val, err := fn()
		result := NewResult(val, err).(U)
		fut.dat <- result
	}()

	return fut
}
