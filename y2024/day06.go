package y2024

import (
	"fmt"

	"github.com/MoraGames/adventofcode/utils"
)

var _ utils.CoordInterface = Coord{}

type Coord struct {
	Row int
	Col int
}

func (c Coord) GetRow() int { return c.Row }
func (c Coord) GetCol() int { return c.Col }

func D6P1() {
	guardMap := utils.SplitMatrixLinesCharacters(ReadFile(6))

	currentCoord := Coord{0, 0}
	{
		var startCoordFinded bool
		for r := 0; r < len(guardMap) && !startCoordFinded; r++ {
			for c := 0; c < len(guardMap[r]) && !startCoordFinded; c++ {
				if guardMap[r][c] == "^" {
					currentCoord = Coord{r, c}
					startCoordFinded = true
				}
			}
		}
	}

	walkedCeils := moveGuard(guardMap, currentCoord, "^")

	utils.PrintSolution("D6P1()", walkedCeils)
}

func D6P2() {
	guardMap := utils.SplitMatrixLinesCharacters(ReadFile(6))

	currentCoord := Coord{0, 0}
	{
		var startCoordFinded bool
		for r := 0; r < len(guardMap) && !startCoordFinded; r++ {
			for c := 0; c < len(guardMap[r]) && !startCoordFinded; c++ {
				if guardMap[r][c] == "^" {
					currentCoord = Coord{r, c}
					startCoordFinded = true
				}
			}
		}
	}

	loopsFound := moveGuard2(guardMap, currentCoord, "^")

	utils.PrintSolution("D6P2()", loopsFound)
}

func moveGuard(guardMap [][]string, currentCoord Coord, guardRotation string) int {
	var walkedCeils int = 1

	for {
		var nextCoord Coord
		switch guardRotation {
		case "^":
			nextCoord = Coord{currentCoord.GetRow() - 1, currentCoord.GetCol()}
		case ">":
			nextCoord = Coord{currentCoord.GetRow(), currentCoord.GetCol() + 1}
		case "v":
			nextCoord = Coord{currentCoord.GetRow() + 1, currentCoord.GetCol()}
		case "<":
			nextCoord = Coord{currentCoord.GetRow(), currentCoord.GetCol() - 1}
		}

		if utils.IsInbound(guardMap, nextCoord) {
			if guardMap[nextCoord.GetRow()][nextCoord.GetCol()] == "#" {
				switch guardRotation {
				case "^":
					guardRotation = ">"
				case ">":
					guardRotation = "v"
				case "v":
					guardRotation = "<"
				case "<":
					guardRotation = "^"
				}

				// just for test
				guardMap[currentCoord.GetRow()][currentCoord.GetCol()] = guardRotation

			} else if guardMap[nextCoord.GetRow()][nextCoord.GetCol()] == "." {
				// just for test
				guardMap[currentCoord.GetRow()][currentCoord.GetCol()] = "X"
				guardMap[nextCoord.GetRow()][nextCoord.GetCol()] = guardRotation

				walkedCeils++
				currentCoord = nextCoord
			} else if guardMap[nextCoord.GetRow()][nextCoord.GetCol()] == "X" {
				// just for test
				guardMap[currentCoord.GetRow()][currentCoord.GetCol()] = "X"
				guardMap[nextCoord.GetRow()][nextCoord.GetCol()] = guardRotation

				currentCoord = nextCoord
			}
		} else {
			break
		}

		// just for test
		// utils.PrintMatrix(guardMap)
	}

	// just for test
	// guardMapOut := make([]string, 0)
	// for r := 0; r < len(guardMap); r++ {
	// 	guardMapOut = append(guardMapOut, strings.Join(guardMap[r], ""))
	// }
	// utils.WriteFile("y2024/outputs/day06/guardMap.txt", strings.Join(guardMapOut, "\n"))

	return walkedCeils
}

func moveGuard2(guardMap [][]string, currentCoord Coord, guardRotation string) int {
	var startingCoord Coord = currentCoord
	var startingGuardRotation string = guardRotation
	var copyGuardMap [][]string

	var loopsFound int
	for r := 0; r < len(guardMap); r++ {
		fmt.Printf("Row: %d/%d\n", r, len(guardMap))
		for c := 0; c < len(guardMap[r]); c++ {
			fmt.Printf(" | Col: %d/%d\n", c, len(guardMap[r]))
			copyGuardMap = utils.MatrixCopy(guardMap)
			currentCoord = startingCoord
			guardRotation = startingGuardRotation

			if !(startingCoord.GetRow() == r && startingCoord.GetCol() == c) && !(guardMap[r][c] == "#") {
				copyGuardMap[r][c] = "O"

				var walkedCeils, inloopWalkedCeils int = 1, 0
				for {
					var nextCoord Coord
					switch guardRotation {
					case "^":
						nextCoord = Coord{currentCoord.GetRow() - 1, currentCoord.GetCol()}
					case ">":
						nextCoord = Coord{currentCoord.GetRow(), currentCoord.GetCol() + 1}
					case "v":
						nextCoord = Coord{currentCoord.GetRow() + 1, currentCoord.GetCol()}
					case "<":
						nextCoord = Coord{currentCoord.GetRow(), currentCoord.GetCol() - 1}
					}

					if utils.IsInbound(copyGuardMap, nextCoord) {
						if copyGuardMap[nextCoord.GetRow()][nextCoord.GetCol()] == "#" || copyGuardMap[nextCoord.GetRow()][nextCoord.GetCol()] == "O" {
							switch guardRotation {
							case "^":
								guardRotation = ">"
							case ">":
								guardRotation = "v"
							case "v":
								guardRotation = "<"
							case "<":
								guardRotation = "^"
							}

							copyGuardMap[currentCoord.GetRow()][currentCoord.GetCol()] = guardRotation

						} else if copyGuardMap[nextCoord.GetRow()][nextCoord.GetCol()] == "." {
							copyGuardMap[currentCoord.GetRow()][currentCoord.GetCol()] = "X"
							copyGuardMap[nextCoord.GetRow()][nextCoord.GetCol()] = guardRotation

							walkedCeils++
							inloopWalkedCeils = 0
							currentCoord = nextCoord
						} else if copyGuardMap[nextCoord.GetRow()][nextCoord.GetCol()] == "X" {
							copyGuardMap[currentCoord.GetRow()][currentCoord.GetCol()] = "X"
							copyGuardMap[nextCoord.GetRow()][nextCoord.GetCol()] = guardRotation

							inloopWalkedCeils++
							currentCoord = nextCoord
						}

						if inloopWalkedCeils >= walkedCeils {
							fmt.Printf(" !  |> Loop found! (%d/%d)\n", inloopWalkedCeils, walkedCeils)

							loopsFound++
							break
						}
					} else {
						break
					}
					// just for test
					// if (r == 6 && c == 3) || (r == 7 && c == 6) || (r == 7 && c == 7) || (r == 8 && c == 1) || (r == 8 && c == 3) || (r == 9 && c == 7) {
					// 	utils.PrintMatrix(copyGuardMap)
					// 	time.Sleep(500 * time.Millisecond)
					// 	fmt.Println()
					// }
				}

				// just for test
				// NOTE: the maps are wrongly generated (take the map at the r,c before the ones with the loops, so are just confusing)
				// guardMapOut := make([]string, 0)
				// for r := 0; r < len(copyGuardMap); r++ {
				// 	guardMapOut = append(guardMapOut, strings.Join(copyGuardMap[r], ""))
				// }
				// utils.WriteFile(fmt.Sprintf("y2024/outputs/day06/loopGuardMap-%d.txt", loopsFound), strings.Join(guardMapOut, "\n"))
			}
		}
	}

	return loopsFound
}
