package y2022

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

// This comment contains the first part of the real input file, which isn't parsed by the functions and the solution are based on the hardcoded values in the stacks variable.
/*
[Q]         [N]             [N]
[H]     [B] [D]             [S] [M]
[C]     [Q] [J]         [V] [Q] [D]
[T]     [S] [Z] [F]     [J] [J] [W]
[N] [G] [T] [S] [V]     [B] [C] [C]
[S] [B] [R] [W] [D] [J] [Q] [R] [Q]
[V] [D] [W] [G] [P] [W] [N] [T] [S]
[B] [W] [F] [L] [M] [F] [L] [G] [J]
 1   2   3   4   5   6   7   8   9
*/

var startingStacks = [][]string{
	{"B", "V", "S", "N", "T", "C", "H", "Q"},
	{"W", "D", "B", "G"},
	{"F", "W", "R", "T", "S", "Q", "B"},
	{"L", "G", "W", "S", "Z", "J", "D", "N"},
	{"M", "P", "D", "V", "F"},
	{"F", "W", "J"},
	{"L", "N", "Q", "B", "J", "V"},
	{"G", "T", "R", "C", "J", "Q", "S", "N"},
	{"J", "S", "Q", "C", "W", "D", "M"},
}

func D5P1() {
	stacks := utils.MatrixCopy(startingStacks)
	moves := utils.SplitLines(ReadFile(5))

	for m := 0; m < len(moves); m++ {
		// fmt.Printf("[DEBUG] moves %d:\n > %s\n", m, moves[m])

		move := strings.Split(moves[m], " ")
		// fmt.Printf("[DEBUG] move:\n > %s\n", move)
		quantity, err := strconv.Atoi(move[1])
		if err != nil {
			panic(err)
		}
		start, err := strconv.Atoi(move[3])
		if err != nil {
			panic(err)
		}
		destination, err := strconv.Atoi(move[5])
		if err != nil {
			panic(err)
		}

		start--
		destination--

		// fmt.Printf("Start: %v\n\tDestination: %v\n\tQuantity: %v\n", stacks[start], stacks[destination], quantity)

		for q := 0; q < quantity; q++ {
			stacks[destination] = append(stacks[destination], stacks[start][len(stacks[start])-1])
			stacks[start] = stacks[start][:len(stacks[start])-1]
		}

		// fmt.Printf("[AFTER]\n\tStart: %v\n\tDestination: %v\n\tQuantity: %v\n", stacks[start], stacks[destination], quantity)
	}

	// fmt.Println()
	// for i := 0; i < 9; i++ {
	// 	fmt.Printf("%v\n", stacks[i])
	// }
	fmt.Print("top-stacks: ")
	for i := 0; i < len(stacks); i++ {
		fmt.Printf("%v", stacks[i][len(stacks[i])-1])
	}
	fmt.Println()
}

func D5P2() {
	stacks := utils.MatrixCopy(startingStacks)
	moves := utils.SplitLines(ReadFile(5))

	for m := 0; m < len(moves); m++ {
		// fmt.Printf("Move: %v\n\t", moves[m])

		move := strings.Split(moves[m], " ")
		quantity, err := strconv.Atoi(move[1])
		if err != nil {
			panic(err)
		}
		start, err := strconv.Atoi(move[3])
		if err != nil {
			panic(err)
		}
		destination, err := strconv.Atoi(move[5])
		if err != nil {
			panic(err)
		}

		start--
		destination--

		// fmt.Printf("[BEFOR]\n\t |- Start: %v\n\t |- Destination: %v\n\t |- Quantity: %v\n\t |- StacksMove: %v\n", stacks[start], stacks[destination], quantity, stacks[start][len(stacks[start])-quantity:len(stacks[start])])

		stacks[destination] = append(stacks[destination], stacks[start][len(stacks[start])-quantity:len(stacks[start])]...)
		stacks[start] = stacks[start][:len(stacks[start])-quantity]

		// fmt.Printf("[AFTER]\n\t |- Start: %v\n\t |- Destination: %v\n\t |- Quantity: %v\n", stacks[start], stacks[destination], quantity)
	}

	// fmt.Println()
	// for i := 0; i < 9; i++ {
	// 	fmt.Printf("%v\n", stacks[i])
	// }
	fmt.Print("top-stacks_2: ")
	for i := 0; i < len(stacks); i++ {
		fmt.Printf("%v", stacks[i][len(stacks[i])-1])
	}
	fmt.Println()
}
