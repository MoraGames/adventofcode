package y2024

import (
	"github.com/MoraGames/adventofcode/utils"
)

var (
	directions = []struct {
		r int
		c int
	}{
		{-1, -1}, // diagonal left up
		{-1, 0},  // up
		{-1, 1},  // diagonal right up
		{0, 1},   // right
		{1, 1},   // diagonal right down
		{1, 0},   // down
		{1, -1},  // diagonal left down
		{0, -1},  // left
	}
)

func D4P1() {
	day, part = 4, 1
	matrix := utils.SplitMatrixLinesCharacters(ReadFile())

	var num int
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == "X" {
				for _, dir := range directions {
				counterLoop:
					for count := 1; count <= 3; {
						if r+dir.r*count >= len(matrix) || r+dir.r*count < 0 {
							break
						}
						if c+dir.c*count >= len(matrix[r]) || c+dir.c*count < 0 {
							break
						}
						switch count {
						case 1:
							if matrix[r+dir.r*count][c+dir.c*count] == "M" {
								count++
							} else {
								break counterLoop
							}
						case 2:
							if matrix[r+dir.r*count][c+dir.c*count] == "A" {
								count++
							} else {
								break counterLoop
							}
						case 3:
							if matrix[r+dir.r*count][c+dir.c*count] == "S" {
								count++
								num++
							} else {
								break counterLoop
							}
						}
					}
				}
			}
		}
	}

	PrintSolution(num)
}

func D4P2() {
	day, part = 4, 2
	matrix := utils.SplitMatrixLinesCharacters(ReadFile())

	var num int
	for r := 0; r < len(matrix); r++ {
		for c := 0; c < len(matrix[r]); c++ {
			if matrix[r][c] == "A" && r+1 < len(matrix) && c+1 < len(matrix[r]) && r-1 >= 0 && c-1 >= 0 {
				if matrix[r-1][c-1] == "M" && matrix[r+1][c+1] == "S" {
					if matrix[r-1][c+1] == "M" && matrix[r+1][c-1] == "S" {
						num++
					} else if matrix[r-1][c+1] == "S" && matrix[r+1][c-1] == "M" {
						num++
					}
				} else if matrix[r-1][c-1] == "S" && matrix[r+1][c+1] == "M" {
					if matrix[r-1][c+1] == "M" && matrix[r+1][c-1] == "S" {
						num++
					} else if matrix[r-1][c+1] == "S" && matrix[r+1][c-1] == "M" {
						num++
					}
				}
			}
		}
	}

	PrintSolution(num)
}
