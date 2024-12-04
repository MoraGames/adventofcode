package y2022

import (
	"fmt"

	"github.com/MoraGames/adventofcode/utils"
)

func D15P1() {
	file := ReadFile(15)

	lines := utils.SplitLines(ReadFile(15))
	for l := 0; l < len(lines); l++ {
		fmt.Printf("%v", lines[l])
	}

	fmt.Printf("%v", file)
}

func D15P2() {
	file := ReadFile(15)

	lines := utils.SplitLines(ReadFile(15))
	for l := 0; l < len(lines); l++ {
		fmt.Printf("%v", lines[l])
	}

	fmt.Printf("%v", file)
}
