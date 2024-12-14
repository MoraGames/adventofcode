package y2024

import (
	"fmt"
	"time"

	"github.com/MoraGames/adventofcode/utils"
)

type robot struct {
	position [2]int
	velocity [2]int
}

func D14P1() {
	day, part = 14, 1
	robotDatas := utils.SplitMatrixLinesSpaces(ReadFile())

	locationMapBounds := [2]int{101, 103} // [x, y]
	locationMap := make(map[utils.MatrixCoord]int)
	for _, robotData := range robotDatas {
		position := utils.SliceStoI(utils.SplitCommas(utils.Split(robotData[0], "=")[1]))
		velocity := utils.SliceStoI(utils.SplitCommas(utils.Split(robotData[1], "=")[1]))

		for s := 1; s <= 100; s++ {
			position[0] += velocity[0]
			position[1] += velocity[1]

			if position[0] < 0 {
				position[0] += locationMapBounds[0]
			}
			if position[1] < 0 {
				position[1] += locationMapBounds[1]
			}
			if position[0] >= locationMapBounds[0] {
				position[0] -= locationMapBounds[0]
			}
			if position[1] >= locationMapBounds[1] {
				position[1] -= locationMapBounds[1]
			}
		}

		utils.CacheAdd(locationMap, utils.MatrixCoord{Row: position[0], Col: position[1]}, 1)
	}

	robotCounterPerQuadrant := [4]int{0, 0, 0, 0}
	for location, count := range locationMap {
		if location.Row == locationMapBounds[0]/2 || location.Col == locationMapBounds[1]/2 {
			continue
		}
		if location.Row < locationMapBounds[0]/2 && location.Col < locationMapBounds[1]/2 {
			robotCounterPerQuadrant[0] += count
		} else if location.Row > locationMapBounds[0]/2 && location.Col < locationMapBounds[1]/2 {
			robotCounterPerQuadrant[1] += count
		} else if location.Row < locationMapBounds[0]/2 && location.Col > locationMapBounds[1]/2 {
			robotCounterPerQuadrant[2] += count
		} else {
			robotCounterPerQuadrant[3] += count
		}
	}

	PrintSolution(robotCounterPerQuadrant[0] * robotCounterPerQuadrant[1] * robotCounterPerQuadrant[2] * robotCounterPerQuadrant[3])
}

// Note:
// This part can't be totally automated because the solution is the number of seconds elapsed when the robot "arrange themselves into a picture of a Christmas tree".
// This algorithm is written to print/save-on-file the location map of robots, after 20 seconds and then each 101 seconds.
// This parameters (shift = +20, module = +101) are found by running every single location map and find a sort of pattern (robots arrange themselves in a vertical line at the center).
// In the try to skip some frames, i've build the algorithm to catch only when this pattern occurs and pray to find the solution in one of these frames.
func D14P2() {
	day, part = 14, 2
	robotDatas := utils.SplitMatrixLinesSpaces(ReadFile())

	robots := make([]robot, 0, len(robotDatas))
	locationMap := [103][101]int{}
	locationMapBounds := [2]int{101, 103} // [x, y]
	for _, robotData := range robotDatas {
		position := utils.SliceStoI(utils.SplitCommas(utils.Split(robotData[0], "=")[1]))
		velocity := utils.SliceStoI(utils.SplitCommas(utils.Split(robotData[1], "=")[1]))

		robots = append(robots, robot{
			position: [2]int{position[0], position[1]},
			velocity: [2]int{velocity[0], velocity[1]},
		})
		locationMap[position[1]][position[0]]++
	}

	for s := 1; ; s++ {
		for r := range robots {
			locationMap[robots[r].position[1]][robots[r].position[0]]--

			robots[r].position[0] += robots[r].velocity[0]
			robots[r].position[1] += robots[r].velocity[1]

			if robots[r].position[0] < 0 {
				robots[r].position[0] += locationMapBounds[0]
			}
			if robots[r].position[1] < 0 {
				robots[r].position[1] += locationMapBounds[1]
			}
			if robots[r].position[0] >= locationMapBounds[0] {
				robots[r].position[0] -= locationMapBounds[0]
			}
			if robots[r].position[1] >= locationMapBounds[1] {
				robots[r].position[1] -= locationMapBounds[1]
			}

			locationMap[robots[r].position[1]][robots[r].position[0]]++
		}

		if (s-20)%101 == 0 {
			fmt.Printf("[s = %v]:\n", s)
			printLocationMap(locationMap)
			time.Sleep(1 * time.Second)

			var str string
			for i := 0; i < 103; i++ {
				str += utils.ToString(locationMap[i][:], false, false)
			}
			WriteOutput(str, fmt.Sprintf("robots_s%2d", s))
		}

	}

	PrintSolution(nil)
}

func printLocationMap(locationMap [103][101]int) {
	for row := 0; row < 103; row++ {
		for col := 0; col < 101; col++ {
			if locationMap[row][col] > 0 {
				fmt.Printf("#")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n\n")
}
