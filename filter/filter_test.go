package filter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterWhere(t *testing.T) {
	t.Run("integer-slice", func(t *testing.T) {
		args := []int{1, 3, 5, 7, 9}
		want := []int{1, 3, 5}

		got := From(args).
			Where(func(v int) bool { return v < 7 }).
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
			Where(func(v Base) bool { return len(v.Name) <= 5 }).
			Where(func(v Base) bool { return v.Age < 40 }).
			Result()

		assert.Equal(t, want, got)
	})
}

func TestFilterTake(t *testing.T) {
	t.Run("integer-slice-after-take", func(t *testing.T) {
		args := []int{1, 3, 5, 7, 9}
		want := []int{1, 3, 5}

		got := From(args).
			Take(3).
			Result()

		assert.Equal(t, want, got)
	})

	t.Run("integer-slice-after-take-over-num", func(t *testing.T) {
		args := []int{1, 3, 5, 7, 9}
		want := []int{1, 3, 5, 7, 9}

		got := From(args).
			Take(100).
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
			Take(2).
			Result()

		assert.Equal(t, want, got)
	})

}

func TestFilterOrderBy(t *testing.T) {
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

func TestFilterWhereSortMix(t *testing.T) {
	t.Run("integer-slice", func(t *testing.T) {
		args := []int{1, 3, 5, 7, 9}
		want := []int{5, 3, 1}

		got := From(args).
			Where(func(v int) bool { return v < 7 }).
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
			Where(func(v Base) bool { return len(v.Name) <= 5 }).
			Sort(func(a, b Base) bool { return a.Age > b.Age }).
			Result()

		assert.Equal(t, want, got)
	})
}
