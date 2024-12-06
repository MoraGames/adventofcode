package y2024

import "github.com/MoraGames/adventofcode/utils"

func Y2024() {
	// D1P1()
	// D1P2()
	// D2P1()
	// D2P2()
	// D3P1()
	// D3P2()
	// D4P1()
	// D4P2()
	// D5P1()
	// D5P2()
	// D6P1()
	// D6P2()
	D7P1()
	D7P2()
}

var (
	year = 2024
	day  = 1
	part = 1
)

func ReadFile() string {
	return utils.ReadFile(year, day)
}

func ReadTestFile() string {
	return utils.ReadTestFile(year, day)
}

func ReadCustomTestFile() string {
	return utils.ReadCustomTestFile(year, day)
}

func WriteOutput(content string, filename string) {
	utils.WriteOutput(content, year, day, filename)
}

func PrintSolution(solution any) {
	utils.PrintSolution(year, day, part, solution)
}
