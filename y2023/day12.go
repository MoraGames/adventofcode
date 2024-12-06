package y2023

import (
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

// type SpringsCount struct {
// 	unknown int
// 	working int
// 	broken  int
// }

func D12P1() {
	day, part = 12, 1
	row := utils.SplitLines(ReadTestFile())

	var sumArrangements int
	for r := 0; r < len(row); r++ {
		infos := utils.SplitSpaces(row[r])
		springsScheme := infos[0]
		springsPattern := utils.SliceStoI(utils.Split(infos[1], ","))

		arrangements := countArrangements(springsScheme, springsPattern)
		sumArrangements += arrangements
		// fmt.Printf("Number of arrangements for %q: %d (sum = %d)\n", springsScheme, arrangements, sumArrangements)
	}

	PrintSolution(sumArrangements)
}

func D12P2() {
}

func countArrangements(condition string, damagedGroups []int) int {
	// Base case: If there are no damaged springs, there is only one arrangement.
	if len(damagedGroups) == 0 {
		return 1
	}

	// Find the indices of "?" in the condition
	questionMarkIndices := findQuestionMarkIndices(condition)

	// Initialize the result
	result := 0

	// Iterate over all possible combinations of "?" placements
	for _, combination := range getCombinations(questionMarkIndices, len(damagedGroups)) {
		// Copy the condition to modify
		modifiedCondition := []rune(condition)

		// Place the damaged groups in the correct positions
		for i, index := range combination {
			for j := index; j < index+damagedGroups[i]; j++ {
				modifiedCondition[j] = '#'
			}
		}

		// Count the arrangements considering the modified condition
		result += countArrangementsWithFixedDamaged(modifiedCondition, damagedGroups[1:])
	}

	return result
}

// findQuestionMarkIndices finds the indices of "?" in the given condition.
func findQuestionMarkIndices(condition string) []int {
	var indices []int
	for i, char := range condition {
		if char == '?' {
			indices = append(indices, i)
		}
	}
	return indices
}

// getCombinations generates all combinations of the specified length
// from the given set of elements (indices).
func getCombinations(elements []int, length int) [][]int {
	if length == 0 {
		return [][]int{{}}
	}

	var combinations [][]int
	for i := range elements {
		subElements := append([]int(nil), elements[i+1:]...)
		for _, subCombination := range getCombinations(subElements, length-1) {
			combinations = append(combinations, append([]int{elements[i]}, subCombination...))
		}
	}
	return combinations
}

// countArrangementsWithFixedDamaged counts the number of different arrangements
// of operational and damaged springs based on the given condition with fixed damaged positions.
func countArrangementsWithFixedDamaged(condition []rune, damagedGroups []int) int {
	if len(damagedGroups) == 0 {
		// Base case: If there are no more damaged groups, check if the condition is valid
		if !strings.Contains(string(condition), "##") {
			return 1
		}
		return 0
	}

	// Find the first available position for the damaged group
	startIndex := strings.Index(string(condition), "##")

	// Initialize the result
	result := 0

	// Iterate over all possible positions for the damaged group
	for i := startIndex; i < len(condition); i++ {
		if condition[i] == '.' {
			// Copy the condition to modify
			modifiedCondition := append([]rune(nil), condition...)

			// Place the damaged group in the correct position
			for j := i; j < i+damagedGroups[0]; j++ {
				modifiedCondition[j] = '#'
			}

			// Recursively count the arrangements with the modified condition
			result += countArrangementsWithFixedDamaged(modifiedCondition, damagedGroups[1:])
		}
	}

	return result
}
