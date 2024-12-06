package y2023

import (
	"github.com/MoraGames/adventofcode/utils"
)

func D9P1() {
	day, part = 9, 1
	oasis := utils.SplitLines(ReadFile())

	var sum int
	for o := 0; o < len(oasis); o++ {
		history := utils.SplitSpaces(oasis[o])

		// fmt.Printf("History:\n  %v\n\n", history)

		var allZero []bool
		var gaps [][]int

		// fmt.Printf("Gaps: (len = %v)\n  %v\n\n", len(gaps), gaps)

		allZero = append(allZero, true)
		gaps = append(gaps, []int{})

		for h := 0; h < len(history)-1; h++ {
			gap := utils.StoI(history[h+1]) - utils.StoI(history[h])
			gaps[0] = append(gaps[0], gap)
			if gap != 0 {
				allZero[len(allZero)-1] = false
			}
		}

		for s := 1; ; s++ {
			if allZero[s-1] {
				break
			}

			allZero = append(allZero, true)
			gaps = append(gaps, []int{})

			for g := 0; g < len(gaps[s-1])-1; g++ {
				gap := gaps[s-1][g+1] - gaps[s-1][g]
				gaps[s] = append(gaps[s], gap)
				if gap != 0 {
					allZero[len(allZero)-1] = false
				}
			}
		}

		// fmt.Printf("All Zero: (len = %v)\n  %v\n\n", len(allZero), allZero)
		// fmt.Printf("Gaps: (len = %v)\n  %v\n\n", len(gaps), gaps)

		// fmt.Printf("Gaps:\n")
		// for g := 0; g < len(gaps); g++ {
		// 	fmt.Printf("  %v  --> allZero = %v\n", gaps[g], allZero[g])
		// }
		// fmt.Printf("\n")

		var newValue []int
		newValue = append(newValue, 0)
		if len(gaps) > 1 {
			for g := len(gaps) - 2; g >= 0; g-- {
				newValue = append(newValue, gaps[g][len(gaps[g])-1]+newValue[len(newValue)-1])
			}
		}
		newValue = append(newValue, utils.StoI(history[len(history)-1])+newValue[len(newValue)-1])

		// fmt.Printf("New Value:\n  %v\n\n", newValue)

		sum += newValue[len(newValue)-1]

		// fmt.Printf("Sum:\n  %v\n\n", sum)
	}

	// Logger.WithFields(logrus.Fields{
	// 	"sum": sum,
	// }).Info("Sum of all new numbers")
	PrintSolution(sum)
}

func D9P2() {
	day, part = 9, 2
	oasis := utils.SplitLines(ReadTestFile())

	var sum int
	for o := 0; o < len(oasis); o++ {
		history := utils.SplitSpaces(oasis[o])

		// fmt.Printf("History:\n  %v\n\n", history)

		var allZero []bool
		var gaps [][]int

		// fmt.Printf("Gaps: (len = %v)\n  %v\n\n", len(gaps), gaps)

		allZero = append(allZero, true)
		gaps = append(gaps, []int{})

		for h := 0; h < len(history)-1; h++ {
			gap := utils.StoI(history[h+1]) - utils.StoI(history[h])
			gaps[0] = append(gaps[0], gap)
			if gap != 0 {
				allZero[len(allZero)-1] = false
			}
		}

		for s := 1; ; s++ {
			if allZero[s-1] {
				break
			}

			allZero = append(allZero, true)
			gaps = append(gaps, []int{})

			for g := 0; g < len(gaps[s-1])-1; g++ {
				gap := gaps[s-1][g+1] - gaps[s-1][g]
				gaps[s] = append(gaps[s], gap)
				if gap != 0 {
					allZero[len(allZero)-1] = false
				}
			}
		}

		// fmt.Printf("All Zero: (len = %v)\n  %v\n\n", len(allZero), allZero)
		// fmt.Printf("Gaps: (len = %v)\n  %v\n\n", len(gaps), gaps)

		// fmt.Printf("Gaps:\n")
		// for g := 0; g < len(gaps); g++ {
		// 	fmt.Printf("  %v  --> allZero = %v\n", gaps[g], allZero[g])
		// }
		// fmt.Printf("\n")

		var newValue []int
		newValue = append(newValue, 0)
		if len(gaps) > 1 {
			for g := len(gaps) - 2; g >= 0; g-- {
				newValue = append(newValue, gaps[g][0]-newValue[len(newValue)-1])
			}
		}
		newValue = append(newValue, utils.StoI(history[0])-newValue[len(newValue)-1])

		// fmt.Printf("New Value:\n  %v\n\n", newValue)

		sum += newValue[len(newValue)-1]

		//fmt.Printf("Sum:\n  %v\n\n", sum)
	}

	// Logger.WithFields(logrus.Fields{
	// 	"sum": sum,
	// }).Info("Sum of all new numbers")

	PrintSolution(sum)
}
