package y2024

import (
	"github.com/MoraGames/adventofcode/utils"
)

var (
	topReachedByTrailHeads    = make(map[utils.MatrixCoord][]utils.MatrixCoord)
	useTopReachedByTrailHeads = true
)

func D10P1() {
	day, part = 10, 1
	useTopReachedByTrailHeads = true
	topoMap := utils.MatrixStoI(utils.SplitMatrixLinesCharacters(ReadFile()))

	trailHeads := make([]utils.MatrixCoord, 0)
	for r := 0; r < len(topoMap); r++ {
		for c := 0; c < len(topoMap[r]); c++ {
			if topoMap[r][c] == 0 {
				trailHeads = append(trailHeads, utils.MatrixCoord{Row: r, Col: c})
			}
		}
	}

	var sum int
	for th := 0; th < len(trailHeads); th++ {
		sum += makeMove(topoMap, trailHeads[th], trailHeads[th])
	}

	PrintSolution(sum)
}

func D10P2() {
	day, part = 10, 2
	useTopReachedByTrailHeads = false
	topoMap := utils.MatrixStoI(utils.SplitMatrixLinesCharacters(ReadFile()))

	trailHeads := make([]utils.MatrixCoord, 0)
	for r := 0; r < len(topoMap); r++ {
		for c := 0; c < len(topoMap[r]); c++ {
			if topoMap[r][c] == 0 {
				trailHeads = append(trailHeads, utils.MatrixCoord{Row: r, Col: c})
			}
		}
	}

	var sum int
	for th := 0; th < len(trailHeads); th++ {
		sum += makeMove(topoMap, trailHeads[th], trailHeads[th])
	}

	PrintSolution(sum)
}

func makeMove(topoMap [][]int, currentCoord utils.MatrixCoord, startingTrailHead utils.MatrixCoord) int {
	// fmt.Printf("Current Coord = %v | Current Value = %v\n", currentCoord, topoMap[currentCoord.Row][currentCoord.Col])
	if utils.MatrixIsInbound(topoMap, currentCoord) && topoMap[currentCoord.Row][currentCoord.Col] == 9 {
		var resp int
		if useTopReachedByTrailHeads {
			if utils.Count(topReachedByTrailHeads[startingTrailHead], currentCoord) == 0 {
				topReachedByTrailHeads[startingTrailHead] = append(topReachedByTrailHeads[startingTrailHead], currentCoord)
				resp = 1
			}
		} else {
			resp = 1
		}
		return resp
	}

	var availableDirections = make([]string, 0)
	if utils.MatrixIsInbound(topoMap, utils.MatrixCoord{Row: currentCoord.Row - 1, Col: currentCoord.Col}) && topoMap[currentCoord.Row-1][currentCoord.Col] == (topoMap[currentCoord.Row][currentCoord.Col]+1) {
		availableDirections = append(availableDirections, "U")
	}
	if utils.MatrixIsInbound(topoMap, utils.MatrixCoord{Row: currentCoord.Row, Col: currentCoord.Col + 1}) && topoMap[currentCoord.Row][currentCoord.Col+1] == (topoMap[currentCoord.Row][currentCoord.Col]+1) {
		availableDirections = append(availableDirections, "R")
	}
	if utils.MatrixIsInbound(topoMap, utils.MatrixCoord{Row: currentCoord.Row + 1, Col: currentCoord.Col}) && topoMap[currentCoord.Row+1][currentCoord.Col] == (topoMap[currentCoord.Row][currentCoord.Col]+1) {
		availableDirections = append(availableDirections, "D")
	}
	if utils.MatrixIsInbound(topoMap, utils.MatrixCoord{Row: currentCoord.Row, Col: currentCoord.Col - 1}) && topoMap[currentCoord.Row][currentCoord.Col-1] == (topoMap[currentCoord.Row][currentCoord.Col]+1) {
		availableDirections = append(availableDirections, "L")
	}

	// fmt.Printf("Available Directions = %v\n", availableDirections)

	var resp int
	for _, dir := range availableDirections {
		nextCoord := utils.MatrixCoord{Row: currentCoord.Row, Col: currentCoord.Col}
		switch dir {
		case "U":
			nextCoord.Row--
		case "R":
			nextCoord.Col++
		case "D":
			nextCoord.Row++
		case "L":
			nextCoord.Col--
		}
		resp += makeMove(topoMap, nextCoord, startingTrailHead)
	}
	return resp
}
