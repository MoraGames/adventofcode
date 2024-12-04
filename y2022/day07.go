package y2022

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D7P1() {
	lines := utils.SplitLines(ReadFile(7))

	directoriesMap := make(map[string]int)
	currentDirectory := make([]string, 0)

	for line := 0; line < len(lines); line++ {
		command := strings.Split(lines[line], " ")
		if command[1] != "ls" {
			if command[1] == "cd" {
				if command[2] == "/" {
					//$ cd /
					currentDirectory = nil
				} else if command[2] == ".." {
					//$ cd ..
					currentDirectory = currentDirectory[:len(currentDirectory)-1]
				} else {
					//$ cd <name>
					currentDirectory = append(currentDirectory, command[2])
				}
			} else if command[0] == "dir" {
				//dir <name>
				if _, exist := directoriesMap[strings.Join(currentDirectory, "/")]; !exist {
					directoriesMap[strings.Join(currentDirectory, "/")] = 0
				}
			} else {
				//<weight> <name>
				if _, exist := directoriesMap[strings.Join(currentDirectory, "/")]; !exist {
					directoriesMap[strings.Join(currentDirectory, "/")] = 0
				}

				fileWeight, err := strconv.Atoi(command[0])
				if err != nil {
					panic(err)
				}

				directoriesMap[strings.Join(currentDirectory, "/")] += fileWeight
				for d := len(currentDirectory) - 1; d >= 0; d-- {
					parentDirectory := strings.Join(currentDirectory[:d], "/")
					directoriesMap[parentDirectory] += fileWeight
				}
			}
		}
	}

	var totSum int
	for _, value := range directoriesMap {
		if value <= 100000 {
			totSum += value
		}
	}

	fmt.Printf("totSum = %d\n", totSum)
}

func D7P2() {
	lines := utils.SplitLines(ReadFile(7))

	directoriesMap := make(map[string]int)
	currentDirectory := make([]string, 0)

	for line := 0; line < len(lines); line++ {
		command := strings.Split(lines[line], " ")
		if command[1] != "ls" {
			if command[1] == "cd" {
				if command[2] == "/" {
					//$ cd /
					currentDirectory = nil
				} else if command[2] == ".." {
					//$ cd ..
					currentDirectory = currentDirectory[:len(currentDirectory)-1]
				} else {
					//$ cd <name>
					currentDirectory = append(currentDirectory, command[2])
				}
			} else if command[0] == "dir" {
				//dir <name>
				if _, exist := directoriesMap[strings.Join(currentDirectory, "/")]; !exist {
					directoriesMap[strings.Join(currentDirectory, "/")] = 0
				}
			} else {
				//<weight> <name>
				if _, exist := directoriesMap[strings.Join(currentDirectory, "/")]; !exist {
					directoriesMap[strings.Join(currentDirectory, "/")] = 0
				}

				fileWeight, err := strconv.Atoi(command[0])
				if err != nil {
					panic(err)
				}

				directoriesMap[strings.Join(currentDirectory, "/")] += fileWeight
				for d := len(currentDirectory) - 1; d >= 0; d-- {
					parentDirectory := strings.Join(currentDirectory[:d], "/")
					directoriesMap[parentDirectory] += fileWeight
				}
			}
		}
	}

	unused := 70000000 - directoriesMap[""]
	spaceNeeded := 30000000 - unused

	var smallest = math.MaxInt
	for _, value := range directoriesMap {
		if value >= spaceNeeded && value < smallest {
			smallest = value
		}
	}

	fmt.Printf("smallest = %d\n", smallest)
}
