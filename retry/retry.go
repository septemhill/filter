package retry

import (
	"context"
	"errors"
	"time"
)

const (
	RETRY_BREAK    = true
	RETRY_CONTINUE = false
)

type RetryFunc[T any] func(context.Context) (*T, error)

// RetryCondFunc returns `true` means break, returns `false` means continue retry
type RetryCondFunc[T any] func(*T, error) bool

type Retry[T any] struct {
	retryTimes int
	fn         RetryFunc[T]
	timeout    time.Duration
	condition  RetryCondFunc[T]
}

func (r *Retry[T]) Start(ctx context.Context) (*T, error) {
	var res *T
	var err error
	timeout := r.timeout

	for i := 0; i < r.retryTimes; i++ {
		res, err = r.fn(ctx)

		if r.condition(res, err) {
			break
		}

		if i == (r.retryTimes - 1) {
			return nil, errors.New("retry failed")
		}

		<-time.NewTimer(timeout).C
		timeout = time.Duration(2) * timeout
	}

	return res, err
}

type RetryOption[T any] func(*Retry[T])

func RetryCondition[T any](condition RetryCondFunc[T]) RetryOption[T] {
	return func(r *Retry[T]) {
		r.condition = condition
	}
}

func NewRetry[T any](fn RetryFunc[T], retryTimes int, timeout time.Duration, opts ...RetryOption[T]) *Retry[T] {
	retry := &Retry[T]{
		fn:         fn,
		retryTimes: retryTimes,
		timeout:    timeout,
		condition:  func(_ *T, err error) bool { return (err == nil) },
	}

	for _, opt := range opts {
		opt(retry)
	}

	return retry
}
