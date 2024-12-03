package utils

import "strconv"

func StoI(input string) int {
	num, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return num
}

func ItoS(input int) string {
	return strconv.Itoa(input)
}

func StoF(input string) float64 {
	num, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic(err)
	}
	return num
}

func FtoS(input float64) string {
	return strconv.FormatFloat(input, 'f', -1, 64)
}
