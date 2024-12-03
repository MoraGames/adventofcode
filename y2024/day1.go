package y2024

import (
	"math"
	"sort"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D1P1() {
	lines := utils.SplitLines(ReadFile(1))

	list1, list2 := make([]int, 0, len(lines)), make([]int, 0, len(lines))
	for i := 0; i < len(lines); i++ {
		values := strings.Split(lines[i], "   ")
		list1 = append(list1, utils.StoI(values[0]))
		list2 = append(list2, utils.StoI(values[1]))
	}

	sort.Ints(list1)
	sort.Ints(list2)

	var distanceSum int
	for i := 0; i < len(list1); i++ {
		distanceSum += int(math.Abs(float64(list1[i] - list2[i])))
	}

	utils.PrintSolution("D1P1()", distanceSum)
}

func D1P2() {
	lines := utils.SplitLines(ReadFile(1))

	list1, list2 := make([]int, 0, len(lines)), make([]int, 0, len(lines))
	for i := 0; i < len(lines); i++ {
		values := strings.Split(lines[i], "   ")
		list1 = append(list1, utils.StoI(values[0]))
		list2 = append(list2, utils.StoI(values[1]))
	}

	var similarityScore int
	for i := 0; i < len(list1); i++ {
		similarityScore += list1[i] * utils.Count(list2, list1[i])
	}

	utils.PrintSolution("D1P2()", similarityScore)
}
