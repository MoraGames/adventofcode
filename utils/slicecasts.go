package utils

import "fmt"

func SliceTtoS[T any](slice []T) []string {
	output := make([]string, 0, len(slice))
	for _, v := range slice {
		output = append(output, fmt.Sprint(v))
	}
	return output
}

func SliceStoI(slice []string) []int {
	output := make([]int, 0, len(slice))
	for _, v := range slice {
		output = append(output, StoI(v))
	}
	return output
}

func SliceItoS(slice []int) []string {
	return SliceTtoS(slice)
}

func SliceStoF(slice []string) []float64 {
	output := make([]float64, 0, len(slice))
	for _, v := range slice {
		output = append(output, StoF(v))
	}
	return output
}

func SliceFtoS(slice []float64) []string {
	return SliceTtoS(slice)
}
