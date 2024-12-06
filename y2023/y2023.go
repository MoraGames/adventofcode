package y2023

import (
	"github.com/MoraGames/adventofcode/utils"
)

// var (
// 	Logger *logrus.Logger
// )

// func init() {
// 	Logger = logrus.New()
// 	Logger.SetLevel(logrus.DebugLevel)
// 	Logger.SetFormatter(&logrus.TextFormatter{
// 		TimestampFormat: "01/02/2006 15:04:05",
// 		FullTimestamp:   true,
// 		ForceColors:     true,
// 	})
// }

// TODO: Complete the tests for the following functions. Try to rewrite them to be more efficient. Complete the rest of the advent.

func Y2023() {
	D1P1()
	D1P2()
	D2P1()
	D2P2()
	D3P1()
	D3P2()
	D4P1()
	D4P2()
	D5P1()
	D5P2() // A lot of time is required to solve this problem (more than 5 minutes)
	D6P1()
	D6P2()
	D7P1()
	D7P2()
	D8P1()
	D8P2() // A lot of time is required to solve this problem (more than 20 minutes)
	D9P1()
	D9P2()   // Wrong answer
	D10P1()  // Wrong answer
	D10P2A() // Runtime Crash
	D10P2B()
	D11P1()
	D11P2()
	D12P1()
	D12P2()
}

var (
	year = 2023
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
