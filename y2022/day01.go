package y2022

import (
	"fmt"
	"strconv"

	"github.com/MoraGames/adventofcode/utils"
)

func D1P1() {
	input := ReadFile(1)

	var maxCalories int

	Elfs := utils.SplitDoubleLines(input)
	for e := 0; e < len(Elfs); e++ {
		var calories int

		ElfBag := utils.SplitLines(Elfs[e])
		for i := 0; i < len(ElfBag); i++ {
			c, err := strconv.Atoi(ElfBag[i])
			if err != nil {
				panic(err)
			}

			calories += c
		}

		if calories > maxCalories {
			maxCalories = calories
		}
	}

	fmt.Printf("maxCalories = %d\n", maxCalories)
}

func D1P2() {
	input := ReadFile(1)

	var maxCalories = []int{-1, -1, -1}

	Elfs := utils.SplitDoubleLines(input)
	for e := 0; e < len(Elfs); e++ {
		var calories int

		ElfBag := utils.SplitLines(Elfs[e])
		for i := 0; i < len(ElfBag); i++ {
			c, err := strconv.Atoi(ElfBag[i])
			if err != nil {
				panic(err)
			}

			calories += c
		}

		if calories > maxCalories[0] {
			maxCalories[2] = maxCalories[1]
			maxCalories[1] = maxCalories[0]
			maxCalories[0] = calories
		} else if calories > maxCalories[1] {
			maxCalories[2] = maxCalories[1]
			maxCalories[1] = calories
		} else if calories > maxCalories[2] {
			maxCalories[2] = calories
		}
	}

	sum := maxCalories[0] + maxCalories[1] + maxCalories[2]
	fmt.Printf("maxCalories = %v\n", maxCalories)
	fmt.Printf("sumTop3Calories = %d\n", sum)
}
