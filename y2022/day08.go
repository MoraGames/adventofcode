package y2022

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D8P1() {
	var grid [][]int
	var spotted [][]bool

	lines := utils.SplitLines(ReadFile(8))
	//lines :=utils.SplitLines(testReadFile(8))

	grid = make([][]int, len(lines))
	spotted = make([][]bool, len(lines))
	for r := 0; r < len(lines); r++ {
		line := strings.Split(lines[r], "")

		grid[r] = make([]int, len(line))
		spotted[r] = make([]bool, len(line))
		for c := 0; c < len(line); c++ {
			height, err := strconv.Atoi(line[c])
			if err != nil {
				panic(err)
			}
			grid[r][c] = height
		}
	}

	var visibleTree, currentHeight int

	//printIntMatrix(grid)
	//printBoolMatrix(spotted)
	//fmt.Println()

	//Top down
	currentHeight = -1
	for c := 0; c < len(grid[0]); c++ {
		for r := 0; r < len(grid); r++ {
			//fmt.Printf("[DEBUG] (%d x %d):\n\theight = %v\n\tgrid = %v\n\tspotted = %v\n", r, c, currentHeight, grid[r][c], spotted[r][c])
			if grid[r][c] > currentHeight {
				currentHeight = grid[r][c]
				if !spotted[r][c] {
					visibleTree++
					spotted[r][c] = true

					//fmt.Printf("\tSpotted[%d][%d] = %v\n", r, c, spotted[r][c])
				}
			}
		}
		currentHeight = -1
	}

	//printIntMatrix(grid)
	//printBoolMatrix(spotted)
	//fmt.Println()

	//Bottom up
	currentHeight = -1
	for c := 0; c < len(grid[0]); c++ {
		for r := len(grid) - 1; r >= 0; r-- {
			//fmt.Printf("[DEBUG] (%d x %d):\n\theight = %v\n\tgrid = %v\n\tspotted = %v\n", r, c, currentHeight, grid[r][c], spotted[r][c])
			if grid[r][c] > currentHeight {
				currentHeight = grid[r][c]
				if !spotted[r][c] {
					visibleTree++
					spotted[r][c] = true

					//fmt.Printf("\tSpotted[%d][%d] = %v\n", r, c, spotted[r][c])
				}
			}
		}
		currentHeight = -1
	}

	//printIntMatrix(grid)
	//printBoolMatrix(spotted)
	//fmt.Println()

	//Left right
	currentHeight = -1
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			//fmt.Printf("[DEBUG] (%d x %d):\n\theight = %v\n\tgrid = %v\n\tspotted = %v\n", r, c, currentHeight, grid[r][c], spotted[r][c])
			if grid[r][c] > currentHeight {
				currentHeight = grid[r][c]
				if !spotted[r][c] {
					visibleTree++
					spotted[r][c] = true

					//fmt.Printf("\tSpotted[%d][%d] = %v\n", r, c, spotted[r][c])
				}
			}
		}
		currentHeight = -1
	}

	//printIntMatrix(grid)
	//printBoolMatrix(spotted)
	//fmt.Println()

	//Right left
	currentHeight = -1
	for r := 0; r < len(grid); r++ {
		for c := len(grid[r]) - 1; c >= 0; c-- {
			//fmt.Printf("[DEBUG] (%d x %d):\n\theight = %v\n\tgrid = %v\n\tspotted = %v\n", r, c, currentHeight, grid[r][c], spotted[r][c])
			if grid[r][c] > currentHeight {
				currentHeight = grid[r][c]
				if !spotted[r][c] {
					visibleTree++
					spotted[r][c] = true

					//fmt.Printf("\tSpotted[%d][%d] = %v\n", r, c, spotted[r][c])
				}
			}
		}
		currentHeight = -1
	}

	utils.PrintMatrix(grid)
	utils.PrintMatrix(spotted)
	fmt.Println()

	fmt.Printf("VisibleTree: %v\n", visibleTree)
}

func D8P2() {
	var grid [][]int
	var scores [][]int

	lines := utils.SplitLines(ReadFile(8))
	//lines :=utils.SplitLines(testReadFile(8))

	grid = make([][]int, len(lines))
	scores = make([][]int, len(lines))
	for r := 0; r < len(lines); r++ {
		line := strings.Split(lines[r], "")

		grid[r] = make([]int, len(line))
		scores[r] = make([]int, len(line))
		for c := 0; c < len(line); c++ {
			height, err := strconv.Atoi(line[c])
			if err != nil {
				panic(err)
			}
			grid[r][c] = height
		}
	}

	var visibleTree [4]int
	var currentHeight, maxScore = -1, -1

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			currentHeight = grid[r][c]
			visibleTree = [4]int{0, 0, 0, 0}

			//fmt.Printf("[%d x %d] See:\n", r+1, c+1)

			//South:
			//fmt.Printf("\tSouth:\n\t  ")
			for rs := r + 1; rs < len(grid); rs++ {
				//fmt.Printf("%d ", grid[rs][c])
				visibleTree[0]++
				if grid[rs][c] >= currentHeight {
					break
				}
			}
			//North:
			//fmt.Printf("\tSouth:\n\t  ")
			for rs := r - 1; rs >= 0; rs-- {
				//fmt.Printf("%d ", grid[rs][c])
				visibleTree[1]++
				if grid[rs][c] >= currentHeight {
					break
				}
			}
			//Est:
			//fmt.Printf("\tSouth:\n\t  ")
			for cs := c + 1; cs < len(grid[r]); cs++ {
				//fmt.Printf("%d ", grid[r][cs])
				visibleTree[2]++
				if grid[r][cs] >= currentHeight {
					break
				}
			}
			//West:
			//fmt.Printf("\tSouth:\n\t  ")
			for cs := c - 1; cs >= 0; cs-- {
				//fmt.Printf("%d ", grid[r][cs])
				visibleTree[3]++
				if grid[r][cs] >= currentHeight {
					break
				}
			}

			score := visibleTree[0] * visibleTree[1] * visibleTree[2] * visibleTree[3]
			//fmt.Printf("[%d x %d] scored %d (%d x %d x %d x %d)\n", r, c, score, visibleTree[0], visibleTree[1], visibleTree[2], visibleTree[3])

			if score > maxScore {
				//fmt.Printf("\t and is a new maxScore!\n")
				maxScore = score
			}
		}
	}

	fmt.Printf("maxScore: %v\n", maxScore)
}
