package utils

type Coord struct {
	ry int
	cx int
}

func IsInbound[T any](matrix [][]T, coords ...Coord) bool {
	for _, coord := range coords {
		if coord.ry < 0 || coord.ry >= len(matrix) || coord.cx < 0 || coord.cx >= len(matrix[coord.ry]) {
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
