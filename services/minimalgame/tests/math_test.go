package tests

import (
	"myboardgames/minimalgame/internal/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestValue struct {
	Value int
}

func getInts() []int {
	return []int{1, 2, 3}
}

func getStructs() []TestValue {
	return []TestValue{{1}, {2}}
}

func TestSum(t *testing.T) {
	t.Run("should sum integer values", func(t *testing.T) {
		t.Parallel()

		values := getInts()
		r := utils.Sum(values...)
		assert.Equal(t, 6, r)
	})

	t.Run("should sum integer values from struct", func(t *testing.T) {
		t.Parallel()

		values := getStructs()
		r := utils.SumFunc(func(v TestValue) int { return v.Value }, values...)
		assert.Equal(t, 3, r)
	})
}

func TestAverage(t *testing.T) {
	t.Run("should average integer values", func(t *testing.T) {
		t.Parallel()

		values := getInts()
		r := utils.Average(values...)
		assert.Equal(t, float64(2), r)
	})

	t.Run("should average integer values from struct", func(t *testing.T) {
		t.Parallel()

		values := getStructs()
		r := utils.AverageFunc(func(v TestValue) int { return v.Value }, values...)
		assert.Equal(t, float64(1.5), r)
	})
}

func TestPick(t *testing.T) {
	t.Run("should pick random value", func(t *testing.T) {
		t.Parallel()

		values := getInts()
		r := utils.Pick(values...)
		assert.NotNil(t, r)
		assert.Contains(t, values, r)
	})

	t.Run("should pick random struct", func(t *testing.T) {
		t.Parallel()

		values := getStructs()
		r := utils.Pick(values...)
		assert.NotNil(t, r)
		assert.Contains(t, values, r)
	})

	t.Run("should pick random value from pickFunc", func(t *testing.T) {
		values := getInts()
		r := utils.PickFunc(func(v int) int { return v }, values...)
		assert.Contains(t, values, r)
	})
}

func TestApplyK(t *testing.T) {
	t.Run("should apply k to each value in values", func(t *testing.T) {
		t.Parallel()

		values := getInts()
		r := utils.ApplyK(func(v, k int) int { return v + k }, 1, values...)
		assert.Equal(t, []int{2, 3, 4}, r)
	})

	t.Run("should apply k to struct", func(t *testing.T) {
		t.Parallel()

		values := getStructs()
		r := utils.ApplyK(func(v TestValue, k int) TestValue { v.Value += k; return v }, 1, values...)
		assert.Equal(t, []TestValue{{2}, {3}}, r)
	})
}

func TestShuffle(t *testing.T) {
	t.Run("should shuffle values", func(t *testing.T) {
		t.Parallel()

		values := getInts()
		utils.Shuffle(values)
		assert.ElementsMatch(t, []int{1, 2, 3}, values)
	})
}
