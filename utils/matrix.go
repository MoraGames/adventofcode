package utils

func MatrixAddCoord[T any](matrix [][]T, coord CoordInterface, value T) [][]T {
	newM := make([][]T, 0)
	for ri, row := range matrix {
		if ri != coord.GetRow() {
			newM = append(newM, row)
			continue
		}
		newM = append(newM, AddIndex(row, coord.GetCol(), value))
	}
	return newM
}

func MatrixRemoveCoord[T any](matrix [][]T, coord CoordInterface) [][]T {
	newM := make([][]T, 0)
	for ri, row := range matrix {
		if ri != coord.GetRow() {
			newM = append(newM, row)
			continue
		}
		newM = append(newM, RemoveIndex(row, coord.GetCol()))
	}
	return newM
}

func MatrixIsInbound[T any](matrix [][]T, coords ...CoordInterface) bool {
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

func MatrixFindFirst[T comparable](matrix [][]T, value T) (CoordInterface, bool) {
	for r, row := range matrix {
		for c, v := range row {
			if v == value {
				return MatrixCoord{r, c}, true
			}
		}
	}
	return MatrixCoord{}, false
}

func MatrixFindLast[T comparable](matrix [][]T, value T) (CoordInterface, bool) {
	for r := len(matrix) - 1; r >= 0; r-- {
		for c := len(matrix[r]) - 1; c >= 0; c-- {
			if matrix[r][c] == value {
				return MatrixCoord{r, c}, true
			}
		}
	}
	return MatrixCoord{}, false
}

func MatrixFindAll[T comparable](matrix [][]T, value T) []CoordInterface {
	coords := make([]CoordInterface, 0)
	for r, row := range matrix {
		for c, v := range row {
			if v == value {
				coords = append(coords, MatrixCoord{r, c})
			}
		}
	}
	if len(coords) == 0 {
		return nil
	}
	return coords
}

func MatrixSwap[T any](matrix [][]T, coord1, coord2 CoordInterface) [][]T {
	matrix[coord1.GetRow()][coord1.GetCol()], matrix[coord2.GetRow()][coord2.GetCol()] = matrix[coord2.GetRow()][coord2.GetCol()], matrix[coord1.GetRow()][coord1.GetCol()]
	return matrix
}

func MatrixCopy[T any](matrixSrc [][]T) [][]T {
	matrixDst := make([][]T, len(matrixSrc))
	for r := 0; r < len(matrixSrc); r++ {
		matrixDst[r] = make([]T, len(matrixSrc[r]))
		copy(matrixDst[r], matrixSrc[r])
	}
	return matrixDst
}
