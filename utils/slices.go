package utils

// Replaced by:
// slices.Insert[S ~[]E, E any](sliceS, i int, v ...E) S
func AddIndex[T any](slice []T, index int, value T) []T {
	newS := make([]T, 0)
	newS = append(newS, slice[:index]...)
	newS = append(newS, value)
	return append(newS, slice[index:]...)
}

func RemoveIndex[T any](slice []T, index int) []T {
	newS := make([]T, 0)
	newS = append(newS, slice[:index]...)
	return append(newS, slice[index+1:]...)
}

func Count[T comparable](slice []T, value T) int {
	var count int
	for _, v := range slice {
		if v == value {
			count++
		}
	}
	return count
}
