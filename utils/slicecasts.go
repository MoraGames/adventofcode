package utils

import "fmt"

func SliceTtoS[T any](s []T) []string {
	var output []string
	for _, v := range s {
		output = append(output, fmt.Sprint(v))
	}
	return output
}

func SliceStoI(input []string) []int {
	var output []int
	for _, v := range input {
		output = append(output, StoI(v))
	}
	return output
}

func SliceItoS(input []int) []string {
	return SliceTtoS(input)
}

func SliceStoF(input []string) []float64 {
	var output []float64
	for _, v := range input {
		output = append(output, StoF(v))
	}
	return output
}

func SliceFtoS(input []float64) []string {
	return SliceTtoS(input)
}
