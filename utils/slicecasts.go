package utils

import "fmt"

func SliceTtoS[T any](s []T) []string {
	output := make([]string, 0, len(s))
	for _, v := range s {
		output = append(output, fmt.Sprint(v))
	}
	return output
}

func SliceStoI(input []string) []int {
	output := make([]int, 0, len(input))
	for _, v := range input {
		output = append(output, StoI(v))
	}
	return output
}

func SliceItoS(input []int) []string {
	return SliceTtoS(input)
}

func SliceStoF(input []string) []float64 {
	output := make([]float64, 0, len(input))
	for _, v := range input {
		output = append(output, StoF(v))
	}
	return output
}

func SliceFtoS(input []float64) []string {
	return SliceTtoS(input)
}
