package y2024

import (
	"math"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D13P1() {
	day, part = 13, 1
	problems := utils.SplitDoubleLines(ReadFile())

	var price int
	for _, problem := range problems {
		lines := utils.SplitLines(problem)

		aCoords := strings.Split(strings.Split(lines[0], ": ")[1], ", ")
		bCoords := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
		pCoords := strings.Split(strings.Split(lines[2], ": ")[1], ", ")

		price += getPrice(
			utils.MatrixCoord{Row: utils.StoI(strings.Split(aCoords[0], "+")[1]), Col: utils.StoI(strings.Split(aCoords[1], "+")[1])},
			utils.MatrixCoord{Row: utils.StoI(strings.Split(bCoords[0], "+")[1]), Col: utils.StoI(strings.Split(bCoords[1], "+")[1])},
			utils.MatrixCoord{Row: utils.StoI(strings.Split(pCoords[0], "=")[1]), Col: utils.StoI(strings.Split(pCoords[1], "=")[1])},
			0,
		)
	}

	PrintSolution(price)
}

func D13P2() {
	day, part = 13, 2
	problems := utils.SplitDoubleLines(ReadFile())

	var price int
	for _, problem := range problems {
		lines := utils.SplitLines(problem)

		aCoords := strings.Split(strings.Split(lines[0], ": ")[1], ", ")
		bCoords := strings.Split(strings.Split(lines[1], ": ")[1], ", ")
		pCoords := strings.Split(strings.Split(lines[2], ": ")[1], ", ")

		price += getPrice(
			utils.MatrixCoord{Row: utils.StoI(strings.Split(aCoords[0], "+")[1]), Col: utils.StoI(strings.Split(aCoords[1], "+")[1])},
			utils.MatrixCoord{Row: utils.StoI(strings.Split(bCoords[0], "+")[1]), Col: utils.StoI(strings.Split(bCoords[1], "+")[1])},
			utils.MatrixCoord{Row: utils.StoI(strings.Split(pCoords[0], "=")[1]), Col: utils.StoI(strings.Split(pCoords[1], "=")[1])},
			10000000000000,
		)
	}

	PrintSolution(price)
}

func getPrice(a, b, p utils.MatrixCoord, offset int) int {
	p.Col += offset
	p.Row += offset

	/*
		aT * a.Col + bT * b.Col = p.Col
		aT * a.Row + bT * b.Row = p.Row
	*/

	aT := float64(p.Col*b.Row-p.Row*b.Col) / float64(a.Col*b.Row-a.Row*b.Col)
	bT := float64(p.Row*a.Col-p.Col*a.Row) / float64(a.Col*b.Row-a.Row*b.Col)

	if aT == math.Trunc(aT) && bT == math.Trunc(bT) {
		return int(aT*3 + bT)
	}
	return 0
}
