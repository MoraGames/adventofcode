package y2023

import (
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

// type Race struct {
// 	time     int
// 	distance int
// }

func D6P1() {
	day, part = 6, 1
	lines := utils.SplitLines(ReadFile())

	times := utils.SplitSpaces(utils.Split(lines[0], ": ")[1])
	distances := utils.SplitSpaces(utils.Split(lines[1], ": ")[1])

	var allWinPossibles []int
	for r := 0; r < len(times); r++ { // for each race
		winPossibles := 0
		for ht := 0; ht <= utils.StoI(times[r]); ht++ {
			speed := ht
			rt := utils.StoI(times[r]) - ht
			yourDistance := speed * rt
			if yourDistance > utils.StoI(distances[r]) {
				winPossibles++
			}
		}
		allWinPossibles = append(allWinPossibles, winPossibles)
	}

	var marginOfError int
	for i := 0; i < len(allWinPossibles); i++ {
		if i == 0 {
			marginOfError = allWinPossibles[0]
		} else {
			marginOfError *= allWinPossibles[i]
		}
	}

	// Logger.WithFields(logrus.Fields{
	// 	"magOfErr": marginOfError,
	// }).Info("Margin of error finded")
	PrintSolution(marginOfError)
}

func D6P2() {
	day, part = 6, 2
	lines := utils.SplitLines(ReadFile())

	recordTime := utils.StoI(strings.ReplaceAll(utils.Split(lines[0], ": ")[1], " ", ""))
	recordDistance := utils.StoI(strings.ReplaceAll(utils.Split(lines[1], ": ")[1], " ", ""))

	// fmt.Printf("t = %v\nd= %v\n", recordTime, recordDistance)

	winPossibles := 0
	for holdTime := 0; holdTime <= recordTime; holdTime++ {
		raceSpeed := holdTime
		raceTime := recordTime - holdTime
		raceDistance := raceSpeed * raceTime
		if raceDistance > recordDistance {
			winPossibles++
		}
	}

	// Logger.WithFields(logrus.Fields{
	// 	"winPos": winPossibles,
	// }).Info("All possibilities")
	PrintSolution(winPossibles)
}
