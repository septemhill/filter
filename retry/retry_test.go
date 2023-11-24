package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func toPtr[T any](v T) *T {
	return &v
}

func Test_Retry(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx       context.Context
		times     int
		duration  time.Duration
		fn        RetryFunc[int]
		condition RetryCondFunc[int]
	}

	tests := []struct {
		name    string
		args    args
		want    *int
		wantErr bool
	}{
		{
			name: "retry-failed",
			args: args{
				ctx:      ctx,
				times:    3,
				duration: time.Millisecond * 20,
				fn: func(ctx context.Context) (*int, error) {
					return nil, errors.New("error")
				},
				condition: func(i *int, err error) bool { return err == nil },
			},
			wantErr: true,
		},
		{
			name: "condition-failed",
			args: args{
				ctx:      ctx,
				times:    3,
				duration: time.Millisecond * 20,
				fn: func(ctx context.Context) (*int, error) {
					return toPtr(123), nil
				},
				condition: func(i *int, err error) bool { return (*i != 123) },
			},
			wantErr: true,
		},
		{
			name: "happy-path",
			args: args{
				ctx:      ctx,
				times:    3,
				duration: time.Millisecond * 20,
				fn: func(ctx context.Context) (*int, error) {
					return toPtr(123), nil
				},
				condition: func(i *int, err error) bool { return err == nil },
			},
			wantErr: false,
			want:    toPtr(123),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewRetry(
				tt.args.fn,
				tt.args.times,
				tt.args.duration,
				RetryCondition(tt.args.condition),
			).Start(tt.args.ctx)

			assert.Equal(t, tt.wantErr, (err != nil))
			assert.Equal(t, tt.want, got)
		})
	}
}
