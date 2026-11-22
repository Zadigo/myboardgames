package utils

import (
	"math/rand"
)

type TypeNumbers interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

func Sum[T TypeNumbers](values ...T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

func SumFunc[T any, R TypeNumbers](f func(k T) R, values ...T) R {
	var sum R
	for _, v := range values {
		sum += f(v)
	}
	return sum
}

// Average calculates the average of the given values. Returns 0 if empty.
//
// For example:
//
//	Average(1, 2, 3) // returns 2
func Average[T TypeNumbers](values ...T) float64 {
	if len(values) == 0 {
		return 0
	}
	return float64(Sum(values...)) / float64(len(values))
}

// AverageFunc calculates the average using a function to get values
//
// For example:
//
//	values := []TestValue{{1}, {2}}
//	utils.AverageFunc(func(v TestValue) int { return v.Value }, values...)
//
// It can also work with structs:
//
//	values := []TestValue{{1}, {2}}
//	utils.AverageFunc(func(v TestValue) int { return v.Value }, values...)
func AverageFunc[T any, R TypeNumbers](f func(k T) R, values ...T) float64 {
	if len(values) == 0 {
		return 0
	}
	return float64(SumFunc(f, values...)) / float64(len(values))
}

// Pick picks a random value from the list
//
//	values := []int{1, 2, 3}
//	utils.Pick(values...)
func Pick[T any](values ...T) T {
	return values[rand.Intn(len(values))]
}

// PickFunc picks a random value and applies f to it
//
// For example:
//
//	values := []int{1, 2, 3}
//	utils.PickFunc(func(v int) int { return v + 1 }, values...)
func PickFunc[T any](f func(k T) T, values ...T) T {
	value := values[rand.Intn(len(values))]
	return f(value)
}

// ApplyK applies k to each value using f and returns the results
//
// For example:
//
//	values := []int{1, 2, 3}
//	utils.ApplyK(func(v, k int) int { return v + k }, 1, values...)
func ApplyK[T any, K TypeNumbers](f func(v T, k K) T, k K, values ...T) []T {
	result := make([]T, len(values))
	for i, v := range values {
		result[i] = f(v, k)
	}
	return result
}

func Shuffle[T any](values []T) {
	rand.Shuffle(len(values), func(i, j int) {
		values[i], values[j] = values[j], values[i]
	})
}
