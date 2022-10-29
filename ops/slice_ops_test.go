package ops

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceOpsFilter(t *testing.T) {
	t.Run("integer-slice", func(t *testing.T) {
		args := []int{1, 3, 5, 7, 9}
		want := []int{1, 3, 5}

		got := From(args).
			Filter(func(v int) bool { return v < 7 }).
			Result()

		assert.Equal(t, want, got)
	})

	t.Run("struct-slice", func(t *testing.T) {
		type Base struct {
			Age  int
			Name string
		}

		args := []Base{
			{Age: 32, Name: "David"},
			{Age: 26, Name: "Lcuy"},
			{Age: 63, Name: "Kiwi"},
			{Age: 78, Name: "Rebecca"},
			{Age: 56, Name: "Falcon"},
		}

		want := []Base{
			{Age: 32, Name: "David"},
			{Age: 26, Name: "Lcuy"},
		}

		got := From(args).
			Filter(func(v Base) bool { return len(v.Name) <= 5 }).
			Filter(func(v Base) bool { return v.Age < 40 }).
			Result()

		assert.Equal(t, want, got)
	})
}

func TestSliceOpsSort(t *testing.T) {
	t.Run("integer-slice", func(t *testing.T) {
		args := []int{1, 3, 5, 7, 9}
		want := []int{9, 7, 5, 3, 1}

		got := From(args).
			Sort(func(a, b int) bool { return a > b }).
			Result()

		assert.Equal(t, want, got)
	})

	t.Run("struct-slice", func(t *testing.T) {
		type Base struct {
			Age  int
			Name string
		}

		args := []Base{
			{Age: 32, Name: "David"},
			{Age: 26, Name: "Lcuy"},
			{Age: 63, Name: "Kiwi"},
			{Age: 78, Name: "Rebecca"},
			{Age: 56, Name: "Aaron"},
		}

		want := []Base{
			{Age: 78, Name: "Rebecca"},
			{Age: 26, Name: "Lcuy"},
			{Age: 63, Name: "Kiwi"},
			{Age: 32, Name: "David"},
			{Age: 56, Name: "Aaron"},
		}

		got := From(args).
			Sort(func(a, b Base) bool { return a.Age > b.Age }).
			Sort(func(a, b Base) bool { return a.Name > b.Name }).
			Result()

		assert.Equal(t, want, got)
	})
}

func TestSliceOpsMap(t *testing.T) {
	t.Run("integer-slice", func(t *testing.T) {
		args := []int{1, 3, 5, 7, 9}
		want := []int{7, 13, 19, 25, 31}

		got := From(args).
			Map(func(v int) int { return v*3 + 4 }).
			Result()

		assert.Equal(t, want, got)
	})

	t.Run("struct-slice", func(t *testing.T) {
		type Base struct {
			Age    int
			Name   string
			Expect string
		}

		args := []Base{
			{Age: 32, Name: "David"},
			{Age: 26, Name: "Lcuy"},
			{Age: 63, Name: "Kiwi"},
			{Age: 78, Name: "Rebecca"},
			{Age: 56, Name: "Aaron"},
		}

		want := []Base{
			{Expect: "David-32"},
			{Expect: "Lcuy-26"},
			{Expect: "Kiwi-63"},
			{Expect: "Rebecca-78"},
			{Expect: "Aaron-56"},
		}

		got := From(args).
			Map(func(v Base) Base {
				return Base{
					Expect: fmt.Sprintf("%s-%d", v.Name, v.Age),
				}
			}).
			Result()

		assert.Equal(t, want, got)
	})
}

func TestSliceOpsFilterSortMix(t *testing.T) {
	t.Run("integer-slice", func(t *testing.T) {
		args := []int{1, 3, 5, 7, 9}
		want := []int{5, 3, 1}

		got := From(args).
			Filter(func(v int) bool { return v < 7 }).
			Sort(func(a, b int) bool { return a > b }).
			Result()

		assert.Equal(t, want, got)
	})

	t.Run("struct-slice", func(t *testing.T) {
		type Base struct {
			Age  int
			Name string
		}

		args := []Base{
			{Age: 32, Name: "David"},
			{Age: 26, Name: "Lcuy"},
			{Age: 63, Name: "Kiwi"},
			{Age: 78, Name: "Rebecca"},
			{Age: 56, Name: "Falcon"},
		}

		want := []Base{
			{Age: 63, Name: "Kiwi"},
			{Age: 32, Name: "David"},
			{Age: 26, Name: "Lcuy"},
		}

		got := From(args).
			Filter(func(v Base) bool { return len(v.Name) <= 5 }).
			Sort(func(a, b Base) bool { return a.Age > b.Age }).
			Result()

		assert.Equal(t, want, got)
	})
}
