package y2024

import (
	"slices"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D15P1() {
	day, part = 15, 1
	input := utils.SplitDoubleLines(ReadFile())

	warehouse := utils.SplitMatrixLinesCharacters(input[0])
	robotMoves := utils.SplitCharacters(strings.ReplaceAll(input[1], "\r\n", ""))

	var exit bool
	var robotCoords utils.MatrixCoord
	for r := 0; r < len(warehouse) && !exit; r++ {
		for c := 0; c < len(warehouse[r]) && !exit; c++ {
			if warehouse[r][c] == "@" {
				robotCoords = utils.MatrixCoord{Row: r, Col: c}
				exit = true
			}
		}
	}

	// fmt.Printf("DEBUG]\n |- robotCoords = %v\n", robotCoords)
	// fmt.Printf("%v\n", utils.MatrixToString(warehouse, false, false))

	for _, move := range robotMoves {
		var shiftRobotCoords utils.MatrixCoord
		switch move {
		case "^":
			shiftRobotCoords = utils.MatrixCoord{Row: -1, Col: 0}
		case ">":
			shiftRobotCoords = utils.MatrixCoord{Row: 0, Col: +1}
		case "v":
			shiftRobotCoords = utils.MatrixCoord{Row: +1, Col: 0}
		case "<":
			shiftRobotCoords = utils.MatrixCoord{Row: 0, Col: -1}
		}
		robotCoords = moveRobot(warehouse, robotCoords, shiftRobotCoords)

		// fmt.Printf("%v\n", utils.MatrixToString(warehouse, false, false))
	}

	var gpsSum int
	for r := 0; r < len(warehouse); r++ {
		for c := 0; c < len(warehouse[r]); c++ {
			if warehouse[r][c] == "O" {
				gpsSum += r*100 + c
			}
		}
	}

	PrintSolution(gpsSum)
}

func D15P2() {
	day, part = 15, 2
	input := utils.SplitDoubleLines(ReadFile())

	oldWarehouse := utils.SplitMatrixLinesCharacters(input[0])
	robotMoves := utils.SplitCharacters(strings.ReplaceAll(input[1], "\r\n", ""))

	var robotCoords utils.MatrixCoord
	var newWarehouse = make([][]string, len(oldWarehouse))
	for r := 0; r < len(oldWarehouse); r++ {
		newWarehouse[r] = make([]string, 2*len(oldWarehouse[r]))
		var newWarehouseRow string
		for c := 0; c < len(oldWarehouse[r]); c++ {
			switch oldWarehouse[r][c] {
			case "#":
				newWarehouseRow += "##"
			case ".":
				newWarehouseRow += ".."
			case "@":
				robotCoords = utils.MatrixCoord{Row: r, Col: c * 2}
				newWarehouseRow += "@."
			case "O":
				newWarehouseRow += "[]"
			}
		}
		newWarehouse[r] = strings.Split(newWarehouseRow, "")
	}

	// fmt.Printf("DEBUG]\n |- robotCoords = %v\n", robotCoords)
	// fmt.Printf("%v\n", utils.MatrixToString(newWarehouse, false, false))

	for _, move := range robotMoves {
		var shiftRobotCoords utils.MatrixCoord
		switch move {
		case "^":
			shiftRobotCoords = utils.MatrixCoord{Row: -1, Col: 0}
		case ">":
			shiftRobotCoords = utils.MatrixCoord{Row: 0, Col: +1}
		case "v":
			shiftRobotCoords = utils.MatrixCoord{Row: +1, Col: 0}
		case "<":
			shiftRobotCoords = utils.MatrixCoord{Row: 0, Col: -1}
		}
		robotCoords = moveRobot2(newWarehouse, robotCoords, shiftRobotCoords)

		// fmt.Printf("%v\n", utils.MatrixToString(newWarehouse, false, false))
	}

	var gpsSum int
	for r := 0; r < len(newWarehouse); r++ {
		for c := 0; c < len(newWarehouse[r]); c++ {
			if newWarehouse[r][c] == "[" {
				gpsSum += r*100 + c
			}
		}
	}

	PrintSolution(gpsSum)
}

func moveRobot(warehouse [][]string, robotCoords utils.MatrixCoord, shiftRobotCoords utils.MatrixCoord) utils.MatrixCoord {
	// fmt.Printf("DEBUG]\n |- robotCoords = %v\n |- shiftRobotCoords = %v\n |- nextRobotCoords = %v\n", robotCoords, shiftRobotCoords, utils.MatrixCoord{Row: robotCoords.Row + shiftRobotCoords.Row, Col: robotCoords.Col + shiftRobotCoords.Col})
	if warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col+shiftRobotCoords.Col] == "#" {
		// fmt.Printf(" |- a wall prevents robot from moving\n")
		return robotCoords
	} else if warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col+shiftRobotCoords.Col] == "." {
		// fmt.Printf(" |- robot moves to an empty space\n")
		warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col+shiftRobotCoords.Col] = "@"
		warehouse[robotCoords.Row][robotCoords.Col] = "."
		return utils.MatrixCoord{Row: robotCoords.Row + shiftRobotCoords.Row, Col: robotCoords.Col + shiftRobotCoords.Col}
	} else {
		// fmt.Printf(" |- robot try to moves onto crates\n")
		for i := 2; ; i++ {
			if warehouse[robotCoords.Row+shiftRobotCoords.Row*i][robotCoords.Col+shiftRobotCoords.Col*i] == "#" {
				// fmt.Printf(" |- a wall prevents crates to be shifted\n")
				return robotCoords
			} else if warehouse[robotCoords.Row+shiftRobotCoords.Row*i][robotCoords.Col+shiftRobotCoords.Col*i] == "." {
				// fmt.Printf(" |- crates are shifted\n")
				warehouse[robotCoords.Row+shiftRobotCoords.Row*i][robotCoords.Col+shiftRobotCoords.Col*i] = "O"
				warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col+shiftRobotCoords.Col] = "@"
				warehouse[robotCoords.Row][robotCoords.Col] = "."
				return utils.MatrixCoord{Row: robotCoords.Row + shiftRobotCoords.Row, Col: robotCoords.Col + shiftRobotCoords.Col}
			}
		}
	}
}

func moveRobot2(warehouse [][]string, robotCoords utils.MatrixCoord, shiftRobotCoords utils.MatrixCoord) utils.MatrixCoord {
	// fmt.Printf("DEBUG]\n |- robotCoords = %v\n |- shiftRobotCoords = %v\n |- nextRobotCoords = %v\n", robotCoords, shiftRobotCoords, utils.MatrixCoord{Row: robotCoords.Row + shiftRobotCoords.Row, Col: robotCoords.Col + shiftRobotCoords.Col})
	if warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col+shiftRobotCoords.Col] == "#" {
		// fmt.Printf(" |- a wall prevents robot from moving\n")
		return robotCoords
	} else if warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col+shiftRobotCoords.Col] == "." {
		// fmt.Printf(" |- robot moves to an empty space\n")
		warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col+shiftRobotCoords.Col] = "@"
		warehouse[robotCoords.Row][robotCoords.Col] = "."
		return utils.MatrixCoord{Row: robotCoords.Row + shiftRobotCoords.Row, Col: robotCoords.Col + shiftRobotCoords.Col}
	} else {
		if shiftRobotCoords.Row == 0 {
			// fmt.Printf(" |- robot try to moves horizontally onto crates\n")
			for i := 2; ; i++ {
				if warehouse[robotCoords.Row+shiftRobotCoords.Row*i][robotCoords.Col+shiftRobotCoords.Col*i] == "#" {
					// fmt.Printf(" |- a wall prevents crates to be shifted\n")
					return robotCoords
				} else if warehouse[robotCoords.Row+shiftRobotCoords.Row*i][robotCoords.Col+shiftRobotCoords.Col*i] == "." {
					// fmt.Printf(" |- crates are shifted\n")
					for ; i > 0; i-- {
						warehouse[robotCoords.Row+shiftRobotCoords.Row*i][robotCoords.Col+shiftRobotCoords.Col*i] = warehouse[robotCoords.Row+shiftRobotCoords.Row*(i-1)][robotCoords.Col+shiftRobotCoords.Col*(i-1)]
					}
					warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col+shiftRobotCoords.Col] = "@"
					warehouse[robotCoords.Row][robotCoords.Col] = "."
					return utils.MatrixCoord{Row: robotCoords.Row + shiftRobotCoords.Row, Col: robotCoords.Col + shiftRobotCoords.Col}
				}
			}
		} else {
			// fmt.Printf(" |- robot try to moves vertically onto crates\n")
			colsToCheck := make([][]int, 1)
			if warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col] == "[" {
				colsToCheck[0] = append(colsToCheck[0], robotCoords.Col, robotCoords.Col+1)
			} else if warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col] == "]" {
				colsToCheck[0] = append(colsToCheck[0], robotCoords.Col, robotCoords.Col-1)
			}

			for i := 2; ; i++ {
				colsToCheckAtNextIteration := make([]int, 0)

				var numOfFreeSpaceAhead int
				for colIndex := 0; colIndex < len(colsToCheck[i-2]); colIndex++ {
					// fmt.Printf(" |- DEBUG]\n |   |- i = %v / colIndex = %v\n |   |- colsToCheck[%v-2] = %v\n", i, colIndex, i, colsToCheck[i-2])
					if warehouse[robotCoords.Row+shiftRobotCoords.Row*i][colsToCheck[i-2][colIndex]] == "#" {
						// fmt.Printf(" |---|- a wall prevents crates to be shifted\n")
						return robotCoords
					} else if warehouse[robotCoords.Row+shiftRobotCoords.Row*i][colsToCheck[i-2][colIndex]] == "[" {
						// fmt.Printf(" |   |- a crates is found in colsToCheck[%v-2][%v] = %v\n", i, colIndex, colsToCheck[i-2][colIndex])
						if !slices.Contains(colsToCheckAtNextIteration, colsToCheck[i-2][colIndex]) {
							colsToCheckAtNextIteration = append(colsToCheckAtNextIteration, colsToCheck[i-2][colIndex])
						}
						if !slices.Contains(colsToCheckAtNextIteration, colsToCheck[i-2][colIndex]+1) {
							colsToCheckAtNextIteration = append(colsToCheckAtNextIteration, colsToCheck[i-2][colIndex]+1)
						}
					} else if warehouse[robotCoords.Row+shiftRobotCoords.Row*i][colsToCheck[i-2][colIndex]] == "]" {
						// fmt.Printf(" |   |- a crates is found in colsToCheck[%v-2][%v] = %v\n", i, colIndex, colsToCheck[i-2][colIndex])
						if !slices.Contains(colsToCheckAtNextIteration, colsToCheck[i-2][colIndex]) {
							colsToCheckAtNextIteration = append(colsToCheckAtNextIteration, colsToCheck[i-2][colIndex])
						}
						if !slices.Contains(colsToCheckAtNextIteration, colsToCheck[i-2][colIndex]-1) {
							colsToCheckAtNextIteration = append(colsToCheckAtNextIteration, colsToCheck[i-2][colIndex]-1)
						}
					} else if warehouse[robotCoords.Row+shiftRobotCoords.Row*i][colsToCheck[i-2][colIndex]] == "." {
						// fmt.Printf(" |   |- a free space is found in colsToCheck[%v-2][%v] = %v\n", i, colIndex, colsToCheck[i-2][colIndex])
						numOfFreeSpaceAhead++
					}
					// fmt.Printf(" |   |- colsToCheckAtNextIteration = %v\n |   |- numOfFreeSpaceAhead = %v\n", colsToCheckAtNextIteration, numOfFreeSpaceAhead)
				}

				// fmt.Printf(" |- %v == %v ?= %v\n", numOfFreeSpaceAhead, len(colsToCheck[i-2]), numOfFreeSpaceAhead == len(colsToCheck[i-2]))
				if numOfFreeSpaceAhead == len(colsToCheck[i-2]) {
					// fmt.Printf(" |- crates are shifted\n")
					for ; i > 1; i-- {
						for c := 0; c < len(colsToCheck[i-2]); c++ {
							warehouse[robotCoords.Row+shiftRobotCoords.Row*i][colsToCheck[i-2][c]] = warehouse[robotCoords.Row+shiftRobotCoords.Row*(i-1)][colsToCheck[i-2][c]]
							warehouse[robotCoords.Row+shiftRobotCoords.Row*(i-1)][colsToCheck[i-2][c]] = "."
						}
					}
					warehouse[robotCoords.Row+shiftRobotCoords.Row][robotCoords.Col] = "@"
					warehouse[robotCoords.Row][robotCoords.Col] = "."
					return utils.MatrixCoord{Row: robotCoords.Row + shiftRobotCoords.Row, Col: robotCoords.Col}
				}

				// fmt.Printf(" |- copying colsToCheckAtNextIteration = %v \n |   into colsToCheck[%v-1] = ", colsToCheckAtNextIteration, i)
				colsToCheck = append(colsToCheck, make([]int, len(colsToCheckAtNextIteration)))
				copy(colsToCheck[i-1], colsToCheckAtNextIteration)
				// fmt.Printf("%v\n", colsToCheck[i-1])
			}
		}
	}
}
