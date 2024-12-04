package y2022

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D4P1() {
	var contained int

	lines := utils.SplitLines(ReadFile(4))
	for l := 0; l < len(lines); l++ {
		pairs := strings.Split(lines[l], ",")

		Elf1 := strings.Split(pairs[0], "-")
		Elf2 := strings.Split(pairs[1], "-")

		elf1F, err := strconv.Atoi(Elf1[0])
		if err != nil {
			panic(err)
		}
		elf1L, err := strconv.Atoi(Elf1[1])
		if err != nil {
			panic(err)
		}
		elf2F, err := strconv.Atoi(Elf2[0])
		if err != nil {
			panic(err)
		}
		elf2L, err := strconv.Atoi(Elf2[1])
		if err != nil {
			panic(err)
		}

		//first contains second
		if elf1F <= elf2F && elf1L >= elf2L {
			contained++
			//second contains first
		} else if elf2F <= elf1F && elf2L >= elf1L {
			contained++
		}
	}

	fmt.Printf("contained: %d\n", contained)
}

func D4P2() {
	var overlapped int

	lines := utils.SplitLines(ReadFile(4))
	for l := 0; l < len(lines); l++ {
		pairs := strings.Split(lines[l], ",")

		Elf1 := strings.Split(pairs[0], "-")
		Elf2 := strings.Split(pairs[1], "-")

		elf1F, err := strconv.Atoi(Elf1[0])
		if err != nil {
			panic(err)
		}
		elf1L, err := strconv.Atoi(Elf1[1])
		if err != nil {
			panic(err)
		}
		elf2F, err := strconv.Atoi(Elf2[0])
		if err != nil {
			panic(err)
		}
		elf2L, err := strconv.Atoi(Elf2[1])
		if err != nil {
			panic(err)
		}

		//fmt.Printf("Pair %d\n\tElf1: %d-%d\n\tEf2: %d-%d\n", l, elf1F, elf1L, elf2F, elf2L)

		if elf1F <= elf2F && elf1L >= elf2F {
			overlapped++
		} else if elf2F <= elf1F && elf2L >= elf1F {
			overlapped++
		}
	}

	fmt.Printf("overlapped: %d\n", overlapped)
}
