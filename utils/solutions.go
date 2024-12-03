package utils

import "fmt"

func PrintSolution[T any](caption string, solution T) {
	fmt.Printf("%v = %v\n", caption, solution)
}
