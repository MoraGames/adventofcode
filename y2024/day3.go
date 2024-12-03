package y2024

import (
	"regexp"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D3P1() {
	instructions := ReadFile(3)

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllStringSubmatch(instructions, -1)

	var result int
	for _, match := range matches {
		temp := utils.SplitCommas(match[0])
		value1 := utils.StoI(strings.Split(temp[0], "(")[1])
		value2 := utils.StoI(strings.Split(temp[1], ")")[0])

		//fmt.Printf("match = %q\n |> values = %v * %v\n", match, value1, value2)

		result += value1 * value2
	}

	utils.PrintSolution("D3P1()", result)
}

func D3P2() {
	instructions := ReadFile(3)

	re := regexp.MustCompile(`(mul\(\d{1,3},\d{1,3}\))|don't\(\)|do\(\)`)
	matches := re.FindAllStringSubmatch(instructions, -1)

	var result int
	var isEnabled bool = true
	for _, match := range matches {
		if match[0] == "don't()" {
			isEnabled = false
			continue
		}
		if match[0] == "do()" {
			isEnabled = true
			continue
		}
		if isEnabled {
			temp := utils.SplitCommas(match[0])
			value1 := utils.StoI(strings.Split(temp[0], "(")[1])
			value2 := utils.StoI(strings.Split(temp[1], ")")[0])

			//fmt.Printf("match = %q\n |> values = %v * %v\n", match, value1, value2)

			result += value1 * value2
		}
	}

	utils.PrintSolution("D3P2()", result)
}
