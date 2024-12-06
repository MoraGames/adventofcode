package utils

import "fmt"

func MatrixTtoS[T any](matrix [][]T) [][]string {
	output := make([][]string, len(matrix))
	for ri, row := range matrix {
		output[ri] = make([]string, len(row))
		for _, val := range row {
			output[ri] = append(output[ri], fmt.Sprint(val))
		}
	}
	return output
}

func MatrixStoI(matrix [][]string) [][]int {
	output := make([][]int, len(matrix))
	for ri, row := range matrix {
		output[ri] = make([]int, len(row))
		for _, val := range row {
			output[ri] = append(output[ri], StoI(val))
		}
	}
	return output
}

func MatrixItoS(matrix [][]int) [][]string {
	return MatrixTtoS(matrix)
}

func MatrixStoF(matrix [][]string) [][]float64 {
	output := make([][]float64, len(matrix))
	for ri, row := range matrix {
		output[ri] = make([]float64, len(row))
		for _, val := range row {
			output[ri] = append(output[ri], StoF(val))
		}
	}
	return output
}

func MatrixFtoS(matrix [][]float64) [][]string {
	return MatrixTtoS(matrix)
}
