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

func SplitCommas(input string) []string {
	return strings.Split(input, ",")
}

func SplitSemicolons(input string) []string {
	return strings.Split(input, ";")
}

func SplitHyphens(input string) []string {
	return strings.Split(input, "-")
}
