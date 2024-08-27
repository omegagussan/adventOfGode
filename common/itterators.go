package common

func SliceTo(n int) []int {
	var s []int
	for i := 0; i < n; i++ {
		s = append(s, i)
	}
	return s
}

func RemoveFromSlice(s []string, idx int) []string {
	return append(s[:idx], s[idx+1:]...)
}

// Contains Obs! O(N) dont use lightly
func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type Pair[T comparable, V any] struct {
	Key   T
	Value V
}

func ToPairList[T comparable, V any](m map[T]V) []Pair[T, V] {
	pairs := make([]Pair[T, V], 0, len(m))
	for k, v := range m {
		pairs = append(pairs, Pair[T, V]{k, v})
	}
	return pairs
}
func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}

func Filter[T any](ts []T, fn func(T) bool) []T {
	result := make([]T, 0)
	for _, t := range ts {
		if fn(t) {
			result = append(result, t)
		}
	}
	return result
}

// Keys get keys of a map
func Keys[T comparable, V any](m map[T]V) []T {
	keys := make([]T, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
