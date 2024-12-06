package y2023

import (
	"fmt"
	"strconv"

	"github.com/MoraGames/adventofcode/utils"
)

func D1P1() {
	// Read from input1.txt file
	day, part = 1, 1
	lines := utils.SplitLines(ReadFile())

	sumNum := 0
	for l := 0; l < len(lines); l++ {
		strFindedInLine := findNumbers(lines[l])
		numFindedInLine, err := strconv.Atoi(strFindedInLine[0] + strFindedInLine[len(strFindedInLine)-1])
		if err != nil {
			panic(err)
		}
		fmt.Printf("Line %v: strFinded = %v  |  numFinded = %v\n", l, strFindedInLine, numFindedInLine)
		sumNum += numFindedInLine
	}

	PrintSolution(sumNum)
}

func D1P2() {
	// Read from input1.txt file
	day, part = 1, 2
	lines := utils.SplitLines(ReadFile())

	sumNum := 0
	for l := 0; l < len(lines); l++ {
		strFindedInLine := findNumbers2(lines[l])
		numFindedInLine, err := strconv.Atoi(strFindedInLine[0] + strFindedInLine[len(strFindedInLine)-1])
		if err != nil {
			panic(err)
		}
		fmt.Printf("Line %v: strFinded = %v  |  numFinded = %v\n", l, strFindedInLine, numFindedInLine)
		sumNum += numFindedInLine
	}

	PrintSolution(sumNum)
}

func findNumbers(str string) []string {
	var finded []string
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		if str[i] == '0' || str[i] == '1' || str[i] == '2' || str[i] == '3' || str[i] == '4' || str[i] == '5' || str[i] == '6' || str[i] == '7' || str[i] == '8' || str[i] == '9' {
			finded = append(finded, string(str[i]))
		}
	}
	return finded
}

func findNumbers2(str string) []string {
	var finded []string
	strlen := len(str)
	for i := 0; i < strlen; i++ {
		if str[i] == '0' || str[i] == '1' || str[i] == '2' || str[i] == '3' || str[i] == '4' || str[i] == '5' || str[i] == '6' || str[i] == '7' || str[i] == '8' || str[i] == '9' {
			finded = append(finded, string(str[i]))
		} else {
			if i+3 < strlen && str[i] == 'z' && str[i+1] == 'e' && str[i+2] == 'r' && str[i+3] == 'o' {
				finded = append(finded, "0")
				// i += 3
			} else if i+2 < strlen && str[i] == 'o' && str[i+1] == 'n' && str[i+2] == 'e' {
				finded = append(finded, "1")
				// i += 2
			} else if i+2 < strlen && str[i] == 't' && str[i+1] == 'w' && str[i+2] == 'o' {
				finded = append(finded, "2")
				// i += 2
			} else if i+4 < strlen && str[i] == 't' && str[i+1] == 'h' && str[i+2] == 'r' && str[i+3] == 'e' && str[i+4] == 'e' {
				finded = append(finded, "3")
				// i += 4
			} else if i+3 < strlen && str[i] == 'f' && str[i+1] == 'o' && str[i+2] == 'u' && str[i+3] == 'r' {
				finded = append(finded, "4")
				// i += 3
			} else if i+3 < strlen && str[i] == 'f' && str[i+1] == 'i' && str[i+2] == 'v' && str[i+3] == 'e' {
				finded = append(finded, "5")
				// i += 3
			} else if i+2 < strlen && str[i] == 's' && str[i+1] == 'i' && str[i+2] == 'x' {
				finded = append(finded, "6")
				// i += 2
			} else if i+4 < strlen && str[i] == 's' && str[i+1] == 'e' && str[i+2] == 'v' && str[i+3] == 'e' && str[i+4] == 'n' {
				finded = append(finded, "7")
				// i += 4
			} else if i+4 < strlen && str[i] == 'e' && str[i+1] == 'i' && str[i+2] == 'g' && str[i+3] == 'h' && str[i+4] == 't' {
				finded = append(finded, "8")
				// i += 4
			} else if i+3 < strlen && str[i] == 'n' && str[i+1] == 'i' && str[i+2] == 'n' && str[i+3] == 'e' {
				finded = append(finded, "9")
				// i += 3
			}
		}
	}
	return finded
}
