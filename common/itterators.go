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
