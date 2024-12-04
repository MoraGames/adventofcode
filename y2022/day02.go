package y2022

import (
	"fmt"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

const (
	win  = 6
	draw = 3
	lose = 0

	rock     = 1
	paper    = 2
	scissors = 3

	//Maybe better?
	A = rock
	B = paper
	C = scissors

	X = lose
	Y = draw
	Z = win
)

func D2P1() {
	lines := utils.SplitLines(ReadFile(2))

	var totScore int
	for l := 0; l < len(lines); l++ {
		moves := strings.Split(lines[l], " ")
		fmt.Printf("[DEBUG] m[0] = %v  |  m[1] = %v\n", moves[0], moves[1])
		score := roundPoints(moves[0], moves[1])
		fmt.Printf("  |  %d\n", score)
		totScore += score
	}

	fmt.Printf("totScore = %d\n", totScore)
}

func roundPoints(elfMove, myMove string) int {
	switch elfMove {
	case "A": //Rock
		if myMove == "X" { //Rock
			fmt.Printf("\tW+R")
			return draw + rock
		} else if myMove == "Y" { //Paper
			fmt.Printf("\tW+P")
			return win + paper
		} else if myMove == "Z" { //Scissors
			fmt.Printf("\tW+S")
			return lose + scissors
		}
	case "B": //Paper
		if myMove == "X" { //Rock
			fmt.Printf("\tL+R")
			return lose + rock
		} else if myMove == "Y" { //Paper
			fmt.Printf("\tL+P")
			return draw + paper
		} else if myMove == "Z" { //Scissors
			fmt.Printf("\tL+S")
			return win + scissors
		}
	case "C": //Scissors
		if myMove == "X" { //Rock
			fmt.Printf("\tD+R")
			return win + rock
		} else if myMove == "Y" { //Paper
			fmt.Printf("\tD+P")
			return lose + paper
		} else if myMove == "Z" { //Scissors
			fmt.Printf("\tD+S")
			return draw + scissors
		}
	default:
		fmt.Println("[ERROR]")
	}
	return -1
}

func D2P2() {
	lines := utils.SplitLines(ReadFile(2))

	var totScore int
	for l := 0; l < len(lines); l++ {
		moves := strings.Split(lines[l], " ")
		fmt.Printf("[DEBUG] m[0] = %v  |  m[1] = %v\n", moves[0], moves[1])
		score := roundPoints2(moves[0], moves[1])
		fmt.Printf("  |  %d\n", score)
		totScore += score
	}

	fmt.Printf("totScore = %d\n", totScore)
}

func roundPoints2(elfMove, myMove string) int {
	switch elfMove {
	case "A": //Rock
		if myMove == "X" { //Lose
			fmt.Printf("\tW+R")
			return lose + scissors
		} else if myMove == "Y" { //Draw
			fmt.Printf("\tW+P")
			return draw + rock
		} else if myMove == "Z" { //Win
			fmt.Printf("\tW+S")
			return win + paper
		}
	case "B": //Paper
		if myMove == "X" { //Lose
			fmt.Printf("\tW+R")
			return lose + rock
		} else if myMove == "Y" { //Draw
			fmt.Printf("\tW+P")
			return draw + paper
		} else if myMove == "Z" { //Win
			fmt.Printf("\tW+S")
			return win + scissors
		}
	case "C": //Scissors
		if myMove == "X" { //Lose
			fmt.Printf("\tW+R")
			return lose + paper
		} else if myMove == "Y" { //Draw
			fmt.Printf("\tW+P")
			return draw + scissors
		} else if myMove == "Z" { //Win
			fmt.Printf("\tW+S")
			return win + rock
		}
	default:
		fmt.Println("[ERROR]")
	}
	return -1
}
