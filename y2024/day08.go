package y2024

import (
	"slices"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D8P1() {
	day, part = 8, 1
	antennasMap := utils.SplitMatrixLinesCharacters(ReadFile())

	var signals = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var antennasCoords map[string][]utils.MatrixCoord = make(map[string][]utils.MatrixCoord)
	for r := 0; r < len(antennasMap); r++ {
		for c := 0; c < len(antennasMap[r]); c++ {
			if slices.Contains(signals, antennasMap[r][c]) {
				antennasCoords[antennasMap[r][c]] = append(antennasCoords[antennasMap[r][c]], utils.MatrixCoord{Row: r, Col: c})
			}
		}
	}

	var antinodesCoords map[utils.MatrixCoord]bool = make(map[utils.MatrixCoord]bool)
	for _, coords := range antennasCoords {
		for c := 0; c < len(coords); c++ {
			for lc := 0; lc < len(coords); lc++ {
				if c != lc {
					locationDifference := utils.MatrixCoord{Row: coords[c].Row - coords[lc].Row, Col: coords[c].Col - coords[lc].Col}
					antinodesCoord := utils.MatrixCoord{Row: coords[c].Row + locationDifference.Row, Col: coords[c].Col + locationDifference.Col}
					if utils.MatrixIsInbound(antennasMap, antinodesCoord) {
						antinodesCoords[antinodesCoord] = true
					}
				}
			}
		}
	}

	PrintSolution(len(antinodesCoords))
}

func D8P2() {
	day, part = 8, 2
	antennasMap := utils.SplitMatrixLinesCharacters(ReadFile())

	var signals = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var antennasCoords map[string][]utils.MatrixCoord = make(map[string][]utils.MatrixCoord)
	for r := 0; r < len(antennasMap); r++ {
		for c := 0; c < len(antennasMap[r]); c++ {
			if slices.Contains(signals, antennasMap[r][c]) {
				antennasCoords[antennasMap[r][c]] = append(antennasCoords[antennasMap[r][c]], utils.MatrixCoord{Row: r, Col: c})
			}
		}
	}

	var antinodesCoords map[utils.MatrixCoord]bool = make(map[utils.MatrixCoord]bool)
	for _, coords := range antennasCoords {
		if len(coords) >= 2 {
			for c := 0; c < len(coords); c++ {
				antinodesCoords[coords[c]] = true
			}
		}
		for c := 0; c < len(coords); c++ {
			for lc := 0; lc < len(coords); lc++ {
				if c != lc {
					locationDifference := utils.MatrixCoord{Row: coords[c].Row - coords[lc].Row, Col: coords[c].Col - coords[lc].Col}

					lastAntinodesCoord := utils.MatrixCoord{Row: coords[c].Row + locationDifference.Row, Col: coords[c].Col + locationDifference.Col}
					for exit := false; !exit; {
						if utils.MatrixIsInbound(antennasMap, lastAntinodesCoord) {
							antinodesCoords[lastAntinodesCoord] = true
							lastAntinodesCoord = utils.MatrixCoord{Row: lastAntinodesCoord.Row + locationDifference.Row, Col: lastAntinodesCoord.Col + locationDifference.Col}
						} else {
							exit = true
						}
					}
				}
			}
		}
	}

	// Only for visualization and debug purposes
	antinodesMap := utils.MatrixCopy(antennasMap)
	for coord := range antinodesCoords {
		antinodesMap[coord.Row][coord.Col] = "#"
	}
	var antinodesMapStr string
	for r := 0; r < len(antinodesMap); r++ {
		antinodesMapStr += strings.Join(antinodesMap[r], "") + "\n"
	}
	WriteOutput(antinodesMapStr, "antinodesMap.txt")

	PrintSolution(len(antinodesCoords))
}
