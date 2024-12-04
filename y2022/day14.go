package y2022

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

/*
type Coord struct {
	r int
	c int
}

func NewCoord(r, c int) Coord {
	return Coord{r, c}
}
*/

func D14P1() {
	//    0  170        0  100
	// r: 0->170   c: 450->550
	var R, C = 170, 100
	grid := make([][]string, R)
	for r := 0; r < R; r++ {
		grid[r] = make([]string, C)
		for c := 0; c < C; c++ {
			grid[r][c] = "."
		}
	}

	lines := utils.SplitLines(ReadFile(14))
	for l := 0; l < len(lines); l++ {
		ceils := strings.Split(lines[l], " -> ")
		for c := 0; c < len(ceils)-1; c++ {
			stringCoord1, stringCoord2 := strings.Split(ceils[c], ","), strings.Split(ceils[c+1], ",")

			c1, err := strconv.Atoi(stringCoord1[0])
			if err != nil {
				panic(err)
			}
			r1, err := strconv.Atoi(stringCoord1[1])
			if err != nil {
				panic(err)
			}
			c2, err := strconv.Atoi(stringCoord2[0])
			if err != nil {
				panic(err)
			}
			r2, err := strconv.Atoi(stringCoord2[1])
			if err != nil {
				panic(err)
			}

			coord1, coord2 := NewCoord(r1, c1-450), NewCoord(r2, c2-450)

			drawLine(grid, coord1, coord2)
		}
	}

	sandSpawner := NewCoord(0, 500-450)
	grid[sandSpawner.r][sandSpawner.c] = "+"

	printGrid(grid)

	for voidFall := false; voidFall == false; {
		currentSand := sandSpawner
		for newSand := false; newSand == false; {
			printSandCoord(currentSand)
			switch moveSand(grid, &currentSand) {
			case 1:
				fmt.Print("\n")

				drawSand(grid, currentSand)
				//printGrid(grid)

				newSand = true
			case 0:
				//printGrid(grid)
			case -1:
				//printGrid(grid)
				newSand = true
				voidFall = true
			}
		}
	}

	var staticSandCounter int
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if grid[r][c] == "o" {
				staticSandCounter++
			}
		}
	}

	printGrid(grid)
	fmt.Printf("\nstaticSandCounter = %d\n", staticSandCounter)
}

func D14P2() {
	//    0  170        0  500
	// r: 0->170   c: 250->750
	var R, C = 170, 500
	grid := make([][]string, R)
	for r := 0; r < R; r++ {
		grid[r] = make([]string, C)
		for c := 0; c < C; c++ {
			grid[r][c] = "."
		}
	}

	file := ReadFile(14) + "\r\n250,165 -> 749,165"
	lines := utils.SplitLines(file)
	for l := 0; l < len(lines); l++ {
		ceils := strings.Split(lines[l], " -> ")
		for c := 0; c < len(ceils)-1; c++ {
			stringCoord1, stringCoord2 := strings.Split(ceils[c], ","), strings.Split(ceils[c+1], ",")

			c1, err := strconv.Atoi(stringCoord1[0])
			if err != nil {
				panic(err)
			}
			r1, err := strconv.Atoi(stringCoord1[1])
			if err != nil {
				panic(err)
			}
			c2, err := strconv.Atoi(stringCoord2[0])
			if err != nil {
				panic(err)
			}
			r2, err := strconv.Atoi(stringCoord2[1])
			if err != nil {
				panic(err)
			}

			coord1, coord2 := NewCoord(r1, c1-250), NewCoord(r2, c2-250)

			drawLine(grid, coord1, coord2)
		}
	}

	sandSpawner := NewCoord(0, 500-250)
	grid[sandSpawner.r][sandSpawner.c] = "+"

	printGrid(grid)

	var staticSandCounter, sscCounter int
	for grid[sandSpawner.r][sandSpawner.c] == "+" {
		currentSand := sandSpawner

		for newSand := false; newSand == false; {
			printSandCoord(currentSand)
			switch moveSand(grid, &currentSand) {
			case 1:
				fmt.Print("\n")

				drawSand(grid, currentSand)
				staticSandCounter++
				if staticSandCounter == 1000000000000000000 {
					sscCounter++
					staticSandCounter = 0
				}

				//printGrid(grid)

				newSand = true
			case 0:
				//printGrid(grid)
			case -1:
				//printGrid(grid)
				newSand = true
			}
		}

		var s string
		_, err := fmt.Scan(s)
		if err != nil {
			panic(err)
		}

		if s == "p" {
			printGrid(grid)
			printConcat2Int(sscCounter, staticSandCounter)
		}
	}

	printGrid(grid)
	printConcat2Int(sscCounter, staticSandCounter)
}

func printConcat2Int(i1, i2 int) {
	s1 := strconv.Itoa(i1)
	s1len := len(s1)

	l1 := 18
	for ; l1 > s1len; l1-- {
		if l1%3 == 0 && l1 != 18 {
			fmt.Printf("'")
		}
		fmt.Printf("0")
	}
	for ; l1 > 0; l1-- {
		if l1%3 == 0 {
			fmt.Printf("'")
		}
		fmt.Printf("%c", s1[s1len-l1])
	}

	s2 := strconv.Itoa(i2)
	s2len := len(s2)

	l2 := 18
	for ; l2 > s2len; l2-- {
		if l2%3 == 0 {
			fmt.Printf("'")
		}
		fmt.Printf("0")
	}
	for ; l2 > 0; l2-- {
		if l2%3 == 0 {
			fmt.Printf("'")
		}
		fmt.Printf("%c", s2[s2len-l2])
	}
	fmt.Printf("\n")
}

func printSandCoord(coord Coord) {
	fmt.Printf("Sand = {%v, %v}  ", coord.r, coord.c)
}

func printGrid(grid [][]string) {
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			fmt.Print(grid[r][c])
		}
		fmt.Print("\n")
	}
	fmt.Print("\n\n\n")
}

func drawLine(grid [][]string, p1, p2 Coord) {
	if p1.r == p2.r {
		//Same r
		startC := int(math.Min(float64(p1.c), float64(p2.c)))
		endC := int(math.Max(float64(p1.c), float64(p2.c)))
		for c := startC; c <= endC; c++ {
			grid[p1.r][c] = "#"
		}
	} else {
		//Same c
		startR := int(math.Min(float64(p1.r), float64(p2.r)))
		endR := int(math.Max(float64(p1.r), float64(p2.r)))
		for r := startR; r <= endR; r++ {
			grid[r][p1.c] = "#"
		}
	}
}

func drawSand(grid [][]string, point Coord) {
	grid[point.r][point.c] = "o"
}

func moveSand(grid [][]string, point *Coord) int {
	if point.r+1 < len(grid) {
		if grid[point.r+1][point.c] == "." {
			point.r += 1
			fmt.Printf("|\n")
			return 0
		} else if grid[point.r+1][point.c-1] == "." {
			point.r += 1
			point.c -= 1
			fmt.Printf("/\n")
			return 0
		} else if grid[point.r+1][point.c+1] == "." {
			point.r += 1
			point.c += 1
			fmt.Printf("\\\n")
			return 0
		}
		fmt.Printf("-\n")
		return 1
	}
	fmt.Printf("!\n")
	return -1
}
