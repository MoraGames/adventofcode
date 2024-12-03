package utils

// Replaced by:
// slices.Insert[S ~[]E, E any](s S, i int, v ...E) S
func AddIndex[T any](s []T, index int, value T) []T {
	newS := make([]T, 0)
	newS = append(newS, s[:index]...)
	newS = append(newS, value)
	return append(newS, s[index:]...)
}

func RemoveIndex[T any](s []T, index int) []T {
	newS := make([]T, 0)
	newS = append(newS, s[:index]...)
	return append(newS, s[index+1:]...)
}

func Count[T comparable](s []T, value T) int {
	var count int
	for _, v := range s {
		if v == value {
			count++
		}
	}
	return count
}
