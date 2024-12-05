package utils

import "strings"

func SplitLines(input string) []string {
	return strings.Split(input, "\r\n")
}

func SplitDoubleLines(input string) []string {
	return strings.Split(input, "\r\n\r\n")
}

func SplitSpaces(input string) []string {
	return strings.Split(input, " ")
}

func SplitCharacters(input string) []string {
	return strings.Split(input, "")
}

func SplitCommas(input string) []string {
	return strings.Split(input, ",")
}

func SplitColons(input string) []string {
	return strings.Split(input, ":")
}

func SplitSemicolons(input string) []string {
	return strings.Split(input, ";")
}

func SplitPipes(input string) []string {
	return strings.Split(input, "|")
}

func SplitHyphens(input string) []string {
	return strings.Split(input, "-")
}

func SplitMatrixLinesCharacters(input string) [][]string {
	lines := SplitLines(input)
	matrix := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		data := SplitCharacters(lines[i])
		matrix[i] = make([]string, len(data))
		copy(matrix[i], data)
	}
	return matrix
}
