package y2022

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

type Coord struct {
	r int
	c int
}

func NewCoord(r, c int) Coord {
	return Coord{r, c}
}

func D9P1() {
	const dim int = 3001 // dim/2 = 1500

	grid := make([][]string, dim)
	for r := 0; r < dim; r++ {
		grid[r] = make([]string, dim)
	}

	var head, tail = Coord{dim / 2, dim / 2}, Coord{dim / 2, dim / 2}
	grid[tail.r][tail.c] = "#"

	lines := utils.SplitLines(ReadFile(9))
	for l := 0; l < len(lines); l++ {
		line := strings.Split(lines[l], " ")
		length, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		for t := 0; t < length; t++ {
			switch line[0] {
			case "U": //Up
				head.r--
			case "D": //Down
				head.r++
			case "R": //Right
				head.c++
			case "L": //Left
				head.c--
			}
			if !closeEnough(head, tail) {
				moveTail(&head, &tail)
				grid[tail.r][tail.c] = "#"
			}
		}
	}

	var counter int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == "#" {
				counter++
			}
		}
	}

	fmt.Printf("counter = %d\n", counter)
}

func D9P2() {
	const dim int = 3001 // dim/2 = 1500

	grid := make([][]string, dim)
	for r := 0; r < dim; r++ {
		grid[r] = make([]string, dim)
	}

	var rope = [10]Coord{{dim / 2, dim / 2}, {dim / 2, dim / 2}, {dim / 2, dim / 2}, {dim / 2, dim / 2}, {dim / 2, dim / 2}, {dim / 2, dim / 2}, {dim / 2, dim / 2}, {dim / 2, dim / 2}, {dim / 2, dim / 2}, {dim / 2, dim / 2}}
	grid[rope[9].r][rope[9].c] = "#"

	lines := utils.SplitLines(ReadFile(9))
	for l := 0; l < len(lines); l++ {
		line := strings.Split(lines[l], " ")
		length, err := strconv.Atoi(line[1])
		if err != nil {
			panic(err)
		}

		for t := 0; t < length; t++ {
			switch line[0] {
			case "U": //Up
				rope[0].r--
			case "D": //Down
				rope[0].r++
			case "R": //Right
				rope[0].c++
			case "L": //Left
				rope[0].c--
			}
			for i := 0; i < 9; i++ {
				if !closeEnough(rope[i], rope[i+1]) {
					moveTail(&rope[i], &rope[i+1])
					grid[rope[9].r][rope[9].c] = "#"
				}
			}
		}
	}

	var counter int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == "#" {
				counter++
			}
		}
	}

	fmt.Printf("counter = %d\n", counter)
}

func closeEnough(head, tail Coord) bool {
	if head.r < tail.r-1 || head.r > tail.r+1 || head.c < tail.c-1 || head.c > tail.c+1 {
		return false
	}
	return true
}

func moveTail(head, tail *Coord) {
	if head.r == tail.r {
		if head.c < tail.c {
			tail.c--
		} else {
			tail.c++
		}
	} else if head.c == tail.c {
		if head.r < tail.r {
			tail.r--
		} else {
			tail.r++
		}
	} else {
		if head.c < tail.c {
			tail.c--
		} else {
			tail.c++
		}
		if head.r < tail.r {
			tail.r--
		} else {
			tail.r++
		}
	}
}
