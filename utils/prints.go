package utils

import "fmt"

func PrintSolution[T any](caption string, solution T) {
	fmt.Printf("%v = %v\n", caption, solution)
}

func PrintSlice[T any](slice []T) {
	fmt.Printf("[")
	for i := 0; i < len(slice); i++ {
		if i != len(slice)-1 {
			fmt.Printf("%v ", slice[i])
		} else {
			fmt.Printf("%v", slice[i])
		}
	}
	fmt.Printf("]\n")
}

func PrintMatrix[T any](matrix [][]T) {
	fmt.Printf("[\n")
	for r := 0; r < len(matrix); r++ {
		fmt.Printf("[")
		for c := 0; c < len(matrix[r]); c++ {
			if c != len(matrix[r])-1 {
				fmt.Printf("%v ", matrix[r][c])
			} else {
				fmt.Printf("%v", matrix[r][c])
			}
		}
		fmt.Printf("]\n")
	}
	fmt.Printf("]\n")
}
