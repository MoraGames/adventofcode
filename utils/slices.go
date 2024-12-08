package utils

import "fmt"

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

func IsInBounds[T any](slice []T, index int) bool {
	return index >= 0 && index < len(slice)
}

func Equals[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, v := range slice1 {
		if v != slice2[i] {
			return false
		}
	}
	return true
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

func FindFirst[T comparable](slice []T, value T) (int, bool) {
	for i, v := range slice {
		if v == value {
			return i, true
		}
	}
	return -1, false
}

func FindLast[T comparable](slice []T, value T) (int, bool) {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == value {
			return i, true
		}
	}
	return -1, false
}

func FindAll[T comparable](slice []T, value T) []int {
	indexes := make([]int, 0)
	for i, v := range slice {
		if v == value {
			indexes = append(indexes, i)
		}
	}
	if len(indexes) == 0 {
		return nil
	}
	return indexes
}

func Swap[T any](slice []T, index1, index2 int) []T {
	slice[index1], slice[index2] = slice[index2], slice[index1]
	return slice
}

func Copy[T any](slice []T) []T {
	newS := make([]T, len(slice))
	copy(newS, slice)
	return newS
}

func ToString[T any](slice []T, withSpaces, withBrackets bool) string {
	var str string
	if withBrackets {
		str += "["
	}
	for i := 0; i < len(slice); i++ {
		if withSpaces && i != len(slice)-1 {
			str += fmt.Sprintf("%v ", slice[i])
		} else {
			str += fmt.Sprintf("%v", slice[i])
		}
	}
	if withBrackets {
		str += "]"
	}
	str += "\n"
	return str
}
