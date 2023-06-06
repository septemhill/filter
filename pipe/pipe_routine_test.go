package pipe

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testAppendString(str string) (string, error) {
	return str + "123", nil
}

func testToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func testMultiplyPi(i int) (float64, error) {
	return float64(i) * 3.1415926, nil
}

func TestPipeRoutine2(t *testing.T) {
	tests := []string{
		"3",
		"4",
		"5",
		"b",
		"8",
		"n",
	}

	t.Run("", func(t *testing.T) {
		got := PipeRoutine2(tests, testAppendString, testToInt)
		assert.Equal(t, []int{
			3123,
			4123,
			5123,
			8123,
		}, got)
	})
}

func TestPipeRoutine3(t *testing.T) {
	tests := []string{
		"3",
		"4",
		"5",
		"b",
		"2",
		"n",
	}

	t.Run("", func(t *testing.T) {
		got := PipeRoutine3(tests, testAppendString, testToInt, testMultiplyPi)
		assert.Equal(t, []float64{
			9811.1936898,
			12952.7862898,
			16094.3788898,
			6669.6010898,
		}, got)
	})
}
