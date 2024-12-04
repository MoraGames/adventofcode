package utils

type MatrixCoord interface {
	GetCol() int
	GetRow() int
}

func IsInbound[T any](matrix [][]T, coords ...MatrixCoord) bool {
	for _, coord := range coords {
		if coord.GetRow() < 0 || coord.GetRow() >= len(matrix) || coord.GetCol() < 0 || coord.GetCol() >= len(matrix[coord.GetRow()]) {
			return false
		}
	}
	return true
}

func MatrixCount[T comparable](matrix [][]T, value T) int {
	var count int
	for _, r := range matrix {
		for _, v := range r {
			if v == value {
				count++
			}
		}
	}
	return count
}
