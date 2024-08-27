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

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
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
