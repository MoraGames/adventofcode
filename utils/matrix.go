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

func MatrixEquals[T comparable](matrix1, matrix2 [][]T) bool {
	if len(matrix1) != len(matrix2) {
		return false
	}
	for i, r := range matrix1 {
		if len(r) != len(matrix2[i]) {
			return false
		}
		for j, v := range r {
			if v != matrix2[i][j] {
				return false
			}
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

func MatrixSwap[T any](matrix [][]T, coord1, coord2 MatrixCoord) [][]T {
	matrix[coord1.GetRow()][coord1.GetCol()], matrix[coord2.GetRow()][coord2.GetCol()] = matrix[coord2.GetRow()][coord2.GetCol()], matrix[coord1.GetRow()][coord1.GetCol()]
	return matrix
}
