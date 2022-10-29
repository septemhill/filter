package future

type Result[T any] interface {
	Value() T
	Error() error
}

type resultImpl[T any] struct {
	val T
	err error
}

func (re *resultImpl[T]) Value() T {
	return re.val
}

func (re *resultImpl[T]) Error() error {
	return re.err
}

func NewResult[T any](val T, err error) Result[T] {
	return &resultImpl[T]{
		val: val,
		err: err,
	}
}
