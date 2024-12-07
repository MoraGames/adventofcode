package y2024

import (
	"math"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D7P1() {
	day, part = 7, 1
	calibrations := utils.SplitLines(ReadFile())

	var sumTestValues int
	for _, calibration := range calibrations {
		data := strings.Split(calibration, ": ")
		testResult := utils.StoI(data[0])

		if checkCalibrationValidity(testResult, strings.Split(data[1], " "), 1) {
			sumTestValues += testResult
		}
	}

	PrintSolution(sumTestValues)
}

func D7P2() {
	day, part = 7, 2
	calibrations := utils.SplitLines(ReadFile())

	var sumTestValues int
	for _, calibration := range calibrations {
		data := strings.Split(calibration, ": ")
		testResult := utils.StoI(data[0])

		if checkCalibrationValidity(testResult, strings.Split(data[1], " "), 2) {
			sumTestValues += testResult
		}
	}

	PrintSolution(sumTestValues)
}

func checkCalibrationValidity(testResult int, operators []string, generationAlgorithmArg int) bool {
	var operationsList []string
	if generationAlgorithmArg == 1 {
		operationsList = generateOperations(operators)
	} else if generationAlgorithmArg == 2 {
		operationsList = generateOperations_2(operators)
	}

	if len(operationsList) == 0 {
		return testResult == utils.StoI(operators[0])
	}

	for _, operations := range operationsList {
		var result int = utils.StoI(operators[0])
		for operationIndex, operation := range utils.SplitCharacters(operations) {
			switch operation {
			case "+":
				result += utils.StoI(operators[1+operationIndex])
			case "*":
				result *= utils.StoI(operators[1+operationIndex])
			case "|":
				result = utils.StoI(utils.ItoS(result) + operators[1+operationIndex])
			}
		}
		if result == testResult {
			return true
		}
	}
	return false
}

func generateOperations(operators []string) []string {
	if len(operators) == 1 {
		return make([]string, 0)
	}

	numCombinations := int(math.Pow(2, float64(len(operators)-1)))
	operations := make([]string, 0, numCombinations)

	for combination := 0; combination < numCombinations; combination++ {
		currCombination := combination
		operatorList := ""
		for op := 0; op < len(operators)-1; op++ {
			operator := "+"
			if currCombination%2 == 1 {
				operator = "*"
			}
			operatorList += operator
			currCombination /= 2
		}
		operations = append(operations, operatorList)
	}

	return operations
}

func generateOperations_2(operators []string) []string {
	if len(operators) == 1 {
		return make([]string, 0)
	}

	numCombinations := int(math.Pow(3, float64(len(operators)-1)))
	operations := make([]string, 0, numCombinations)

	for combination := 0; combination < numCombinations; combination++ {
		currCombination := combination
		operatorList := ""
		for op := 0; op < len(operators)-1; op++ {
			operator := "+"
			if currCombination%3 == 1 {
				operator = "*"
			} else if currCombination%3 == 2 {
				operator = "|"
			}
			operatorList += operator
			currCombination /= 3
		}
		operations = append(operations, operatorList)
	}

	return operations
}
