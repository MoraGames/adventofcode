package utils

import "strings"

func Split(input string, sep string) []string {
	return strings.Split(input, sep)
}

func SplitLines(input string) []string {
	return Split(input, "\r\n")
}

func SplitDoubleLines(input string) []string {
	return Split(input, "\r\n\r\n")
}

func SplitSpaces(input string) []string {
	return Split(input, " ")
}

func SplitCharacters(input string) []string {
	return Split(input, "")
}

func SplitCommas(input string) []string {
	return Split(input, ",")
}

func SplitColons(input string) []string {
	return Split(input, ":")
}

func SplitSemicolons(input string) []string {
	return Split(input, ";")
}

func SplitPipes(input string) []string {
	return Split(input, "|")
}

func SplitHyphens(input string) []string {
	return Split(input, "-")
}

func SplitMatrix(input string, rowSep, colSep string) [][]string {
	lines := strings.Split(input, rowSep)
	matrix := make([][]string, len(lines))
	for i := 0; i < len(lines); i++ {
		data := strings.Split(lines[i], colSep)
		matrix[i] = make([]string, len(data))
		copy(matrix[i], data)
	}
	return matrix
}

func SplitMatrixLinesSpaces(input string) [][]string {
	return SplitMatrix(input, "\r\n", " ")
}

func SplitMatrixLinesCharacters(input string) [][]string {
	return SplitMatrix(input, "\r\n", "")
}

func SplitMatrixLinesCommas(input string) [][]string {
	return SplitMatrix(input, "\r\n", ",")
}

func SplitMatrixLinesColons(input string) [][]string {
	return SplitMatrix(input, "\r\n", ":")
}

func SplitMatrixLinesSemicolons(input string) [][]string {
	return SplitMatrix(input, "\r\n", ";")
}

func SplitMatrixLinesPipes(input string) [][]string {
	return SplitMatrix(input, "\r\n", "|")
}

func SplitMatrixLinesHyphens(input string) [][]string {
	return SplitMatrix(input, "\r\n", "-")
}

func SplitMatrixDoubleLinesSpaces(input string) [][]string {
	return SplitMatrix(input, "\r\n\r\n", " ")
}

func SplitMatrixDoubleLinesCharacters(input string) [][]string {
	return SplitMatrix(input, "\r\n\r\n", "")
}

func SplitMatrixDoubleLinesCommas(input string) [][]string {
	return SplitMatrix(input, "\r\n\r\n", ",")
}

func SplitMatrixDoubleLinesColons(input string) [][]string {
	return SplitMatrix(input, "\r\n\r\n", ":")
}

func SplitMatrixDoubleLinesSemicolons(input string) [][]string {
	return SplitMatrix(input, "\r\n\r\n", ";")
}

func SplitMatrixDoubleLinesPipes(input string) [][]string {
	return SplitMatrix(input, "\r\n\r\n", "|")
}

func SplitMatrixDoubleLinesHyphens(input string) [][]string {
	return SplitMatrix(input, "\r\n\r\n", "-")
}
