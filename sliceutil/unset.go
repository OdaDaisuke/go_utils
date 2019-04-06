package sliceutil

func unset(a []int, idx int) []int {
	if idx >= len(a) {
		return a
	}
	return append(a[:idx], a[idx + 1:]...)
}