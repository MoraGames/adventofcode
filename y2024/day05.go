package y2024

import (
	"github.com/MoraGames/adventofcode/utils"
)

func D5P1() {
	inputs := utils.SplitDoubleLines(ReadFile(5))

	rules := utils.SplitLines(inputs[0])
	updates := utils.SplitLines(inputs[1])

	safeUpdatesMidValues := make([]int, 0)
	for _, update := range updates {
		values := utils.SplitCommas(update)
		var safeUpdate bool = true
		for i := 1; i < len(values); i++ {
			if !allNumberRulesValidated(values[i], values[:i], rules) {
				safeUpdate = false
				break
			}
		}
		if safeUpdate {
			safeUpdatesMidValues = append(safeUpdatesMidValues, utils.StoI(values[len(values)/2]))
		}
	}

	sumMidValues := 0
	for _, value := range safeUpdatesMidValues {
		sumMidValues += value
	}

	utils.PrintSolution("D5P1()", sumMidValues)
}

func D5P2() {
	inputs := utils.SplitDoubleLines(ReadFile(5))

	rules := utils.SplitLines(inputs[0])
	updates := utils.SplitLines(inputs[1])

	safeUpdatesMidValues := make([]int, 0)
	for _, update := range updates {
		values := utils.SplitCommas(update)
		if wasUnsafe := checkRulesAndCorrectValues(values, rules); wasUnsafe {
			// fmt.Printf("UPDATED   = %v\n | MidValIndex = %v\n | MidVal = %v\n\n", values, len(values)/2, utils.StoI(values[len(values)/2]))
			safeUpdatesMidValues = append(safeUpdatesMidValues, utils.StoI(values[len(values)/2]))
		}
	}

	sumMidValues := 0
	for _, value := range safeUpdatesMidValues {
		sumMidValues += value
	}

	utils.PrintSolution("D5P2()", sumMidValues)
}

func allNumberRulesValidated(value string, valuesBefore, rules []string) bool {
	for _, rule := range rules {
		ruleValues := utils.SplitPipes(rule)
		if value == ruleValues[0] {
			var ruleValidated bool = true
			for _, otherValue := range valuesBefore {
				if otherValue == ruleValues[1] {
					ruleValidated = false
				}
			}
			if !ruleValidated {
				return false
			}
		}
	}
	return true
}

func checkRulesAndCorrectValues(values, rules []string) bool {
	// fmt.Printf("UNKOWN    = %v\n", values)
	var wasUnsafe bool

	lastResult := make([]string, len(values))
	for !utils.Equals(lastResult, values) {
		copy(lastResult, values)
		for i := 1; i < len(values); i++ {
		recheckAllRules:
			for _, rule := range rules {
				ruleValues := utils.SplitPipes(rule)
				if values[i] == ruleValues[0] {
					var ruleValidated bool = true
					for otherValueIndex, otherValue := range values[:i] {
						if otherValue == ruleValues[1] {
							wasUnsafe = true
							ruleValidated = false
							// fmt.Printf("UNSAFE    = %v\n | caused by %q rule\n", values, rule)
							values = utils.Swap(values, i, otherValueIndex)
							// fmt.Printf("CORRECTED = %v\n", values)
						}
					}
					if !ruleValidated {
						break recheckAllRules
					}
				}
			}
		}
	}

	// fmt.Printf("FINAL     = %v\n", values)
	return wasUnsafe
}
