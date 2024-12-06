package y2024

import (
	"math"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D2P1() {
	day, part = 2, 1
	reports := utils.SplitLines(ReadFile())

	var safeNum int
	for _, report := range reports {
		levels := utils.SliceStoI(strings.Split(report, " "))
		var isAsc, isSafe bool = true, true
		for l := 0; l < len(levels)-1; l++ {
			if l == 0 {
				isAsc = levels[l] < levels[l+1]
				if math.Abs(float64(levels[l]-levels[l+1])) < 1 || math.Abs(float64(levels[l]-levels[l+1])) > 3 {
					isSafe = false
					break
				}
			} else if levels[l] < levels[l+1] != isAsc || math.Abs(float64(levels[l]-levels[l+1])) < 1 || math.Abs(float64(levels[l]-levels[l+1])) > 3 {
				isSafe = false
				break
			}
		}
		if isSafe {
			safeNum++
		}
	}

	PrintSolution(safeNum)
}

func D2P2() {
	day, part = 2, 2
	reports := utils.SplitLines(ReadFile())

	var safeNum int
	for _, report := range reports {
		levels := utils.SliceStoI(strings.Split(report, " "))
		var isAsc bool = true
		var safeLvls int = len(levels)
		for l := 0; l < len(levels)-1; l++ {
			if l == 0 {
				isAsc = levels[l] < levels[l+1]
				if math.Abs(float64(levels[l]-levels[l+1])) < 1 || math.Abs(float64(levels[l]-levels[l+1])) > 3 {
					safeLvls--
				}
			} else if levels[l] < levels[l+1] != isAsc || math.Abs(float64(levels[l]-levels[l+1])) < 1 || math.Abs(float64(levels[l]-levels[l+1])) > 3 {
				safeLvls--
			}
		}
		if safeLvls >= len(levels)-1 {
			safeNum++
		}
	}

	PrintSolution(safeNum)
}
