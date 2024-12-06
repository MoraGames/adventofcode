package y2023

import (
	"unicode"

	"github.com/MoraGames/adventofcode/utils"
)

type (
	Checks map[Coord]bool

	//Coord struct {
	//	y, x int
	//}
)

func D3P1() {
	day, part = 3, 1
	lines := utils.SplitLines(ReadFile())
	var sumNums int
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] != '.' && !unicode.IsDigit(rune(lines[y][x])) { // A symbol
				checks := Checks{
					{y: y - 1, x: x - 1}: false,
					{y: y - 1, x: x}:     false,
					{y: y - 1, x: x + 1}: false,
					{y: y, x: x - 1}:     false,
					{y: y, x: x + 1}:     false,
					{y: y + 1, x: x - 1}: false,
					{y: y + 1, x: x}:     false,
					{y: y + 1, x: x + 1}: false,
				}

				if y == 0 {
					checks[Coord{y: y - 1, x: x - 1}] = true
					checks[Coord{y: y - 1, x: x}] = true
					checks[Coord{y: y - 1, x: x + 1}] = true
				}
				if y == len(lines)-1 {
					checks[Coord{y: y + 1, x: x - 1}] = true
					checks[Coord{y: y + 1, x: x}] = true
					checks[Coord{y: y + 1, x: x + 1}] = true
				}
				if x == 0 {
					checks[Coord{y: y - 1, x: x - 1}] = true
					checks[Coord{y: y, x: x - 1}] = true
					checks[Coord{y: y + 1, x: x - 1}] = true
				}
				if x == len(lines[y])-1 {
					checks[Coord{y: y - 1, x: x + 1}] = true
					checks[Coord{y: y, x: x + 1}] = true
					checks[Coord{y: y + 1, x: x + 1}] = true
				}

				// Logger.WithFields(logrus.Fields{
				// 	"symbol": string(lines[y][x]),
				// 	"checks": checks,
				// }).Info("Finding a symbol")

				for coord, done := range checks {
					// Logger.WithFields(logrus.Fields{
					// 	"coord": coord,
					// 	"done":  done,
					// }).Debug("Ranging on map")
					if !done && unicode.IsDigit(rune(lines[coord.y][coord.x])) {
						// Logger.WithFields(logrus.Fields{
						// 	"y": coord.y,
						// 	"x": coord.x,
						// }).Debug("Checking a position")
						numStr := string(lines[coord.y][coord.x])
						for offsetX := coord.x + 1; offsetX < len(lines[coord.y]); offsetX++ { // Check forwards
							if unicode.IsDigit(rune(lines[coord.y][offsetX])) {
								numStr += string(lines[coord.y][offsetX])
								if _, exist := checks[Coord{y: coord.y, x: offsetX}]; exist {
									checks[Coord{y: coord.y, x: offsetX}] = true
								}
							} else {
								// Logger.WithFields(logrus.Fields{
								// 	"offsetX": offsetX,
								// }).Debug("Exiting forwards loop")
								break
							}
						}
						for offsetX := coord.x - 1; offsetX >= 0; offsetX-- { // Check backwards
							if unicode.IsDigit(rune(lines[coord.y][offsetX])) {
								numStr = string(lines[coord.y][offsetX]) + numStr
								if _, exist := checks[Coord{y: coord.y, x: offsetX}]; exist {
									checks[Coord{y: coord.y, x: offsetX}] = true
								}
							} else {
								// Logger.WithFields(logrus.Fields{
								// 	"offsetX": offsetX,
								// }).Debug("Exiting backwards loop")
								break
							}
						}
						num := utils.StoI(numStr)
						sumNums += num
						// Logger.WithFields(logrus.Fields{
						// 	"num": num,
						// 	"sum": sumNums,
						// }).Debug("Found a valid number")
					}
				}
			}
		}
	}

	// Logger.WithField("sum", sumNums).Info("Sum of all valid numbers found")
	PrintSolution(sumNums)
}

func D3P2() {
	day, part = 3, 2
	lines := utils.SplitLines(ReadFile())
	var sumNums int
	for y := 0; y < len(lines); y++ {
		for x := 0; x < len(lines[y]); x++ {
			if lines[y][x] != '.' && string(lines[y][x]) == "*" { // Maybe a gear symbol
				checks := Checks{
					{y: y - 1, x: x - 1}: false,
					{y: y - 1, x: x}:     false,
					{y: y - 1, x: x + 1}: false,
					{y: y, x: x - 1}:     false,
					{y: y, x: x + 1}:     false,
					{y: y + 1, x: x - 1}: false,
					{y: y + 1, x: x}:     false,
					{y: y + 1, x: x + 1}: false,
				}

				if y == 0 {
					checks[Coord{y: y - 1, x: x - 1}] = true
					checks[Coord{y: y - 1, x: x}] = true
					checks[Coord{y: y - 1, x: x + 1}] = true
				}
				if y == len(lines)-1 {
					checks[Coord{y: y + 1, x: x - 1}] = true
					checks[Coord{y: y + 1, x: x}] = true
					checks[Coord{y: y + 1, x: x + 1}] = true
				}
				if x == 0 {
					checks[Coord{y: y - 1, x: x - 1}] = true
					checks[Coord{y: y, x: x - 1}] = true
					checks[Coord{y: y + 1, x: x - 1}] = true
				}
				if x == len(lines[y])-1 {
					checks[Coord{y: y - 1, x: x + 1}] = true
					checks[Coord{y: y, x: x + 1}] = true
					checks[Coord{y: y + 1, x: x + 1}] = true
				}

				// Logger.WithFields(logrus.Fields{
				// 	"symbol": string(lines[y][x]),
				// 	"checks": checks,
				// }).Info("Finding a symbol")

				var allSymbolNumbers []int
				for coord, done := range checks {
					// Logger.WithFields(logrus.Fields{
					// 	"coord": coord,
					// 	"done":  done,
					// }).Debug("Ranging on map")
					if !done && unicode.IsDigit(rune(lines[coord.y][coord.x])) {
						// Logger.WithFields(logrus.Fields{
						// 	"y": coord.y,
						// 	"x": coord.x,
						// }).Debug("Checking a position")
						numStr := string(lines[coord.y][coord.x])
						for offsetX := coord.x + 1; offsetX < len(lines[coord.y]); offsetX++ { // Check forwards
							if unicode.IsDigit(rune(lines[coord.y][offsetX])) {
								numStr += string(lines[coord.y][offsetX])
								if _, exist := checks[Coord{y: coord.y, x: offsetX}]; exist {
									checks[Coord{y: coord.y, x: offsetX}] = true
								}
							} else {
								// Logger.WithFields(logrus.Fields{
								// 	"offsetX": offsetX,
								// }).Debug("Exiting forwards loop")
								break
							}
						}
						for offsetX := coord.x - 1; offsetX >= 0; offsetX-- { // Check backwards
							if unicode.IsDigit(rune(lines[coord.y][offsetX])) {
								numStr = string(lines[coord.y][offsetX]) + numStr
								if _, exist := checks[Coord{y: coord.y, x: offsetX}]; exist {
									checks[Coord{y: coord.y, x: offsetX}] = true
								}
							} else {
								// Logger.WithFields(logrus.Fields{
								// 	"offsetX": offsetX,
								// }).Debug("Exiting backwards loop")
								break
							}
						}
						num := utils.StoI(numStr)
						allSymbolNumbers = append(allSymbolNumbers, num)
						// Logger.WithFields(logrus.Fields{
						// 	"num": num,
						// 	"asn": allSymbolNumbers,
						// }).Debug("Found a valid number")
					}
				}
				if len(allSymbolNumbers) == 2 { // A gear symbol
					sumNums += allSymbolNumbers[0] * allSymbolNumbers[1]
					// Logger.WithFields(logrus.Fields{
					// 	"sum": sumNums,
					// }).Debug("Found a gear")
				}
			}
		}
	}

	// Logger.WithFields(logrus.Fields{
	// 	"sum": sumNums,
	// }).Info("Sum of all valid numbers found")
	PrintSolution(sumNums)
}
