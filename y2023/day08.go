package y2023

import (
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

type Node struct {
	left  string
	right string
}

var Tree map[string]Node = make(map[string]Node)

func D8P1() {
	day, part = 8, 1
	sections := utils.SplitDoubleLines(ReadFile())
	instructions := utils.Split(sections[0], "")
	treeNodes := utils.SplitLines(sections[1])

	for tn := 0; tn < len(treeNodes); tn++ {
		tnInfo := utils.Split(treeNodes[tn], " = ")
		paths := utils.Split(strings.Trim(tnInfo[1], "()"), ", ")
		Tree[tnInfo[0]] = Node{paths[0], paths[1]}
	}

	var finished bool = false
	var current string = "AAA"
	var lenInstrs, lenPath int = len(instructions), 0
	for !finished {
		for i := 0; ; i++ {
			if instructions[i%lenInstrs] == "L" {
				current = Tree[current].left
				lenPath++
			} else if instructions[i%lenInstrs] == "R" {
				current = Tree[current].right
				lenPath++
			}
			if current == "ZZZ" {
				finished = true
				break
			}
		}
	}

	// fmt.Println(lenPath)
	PrintSolution(lenPath)
}

func D8P2() {
	day, part = 8, 2
	sections := utils.SplitDoubleLines(ReadFile())
	instructions := utils.Split(sections[0], "")
	treeNodes := utils.SplitLines(sections[1])

	var currentNodes []string
	for tn := 0; tn < len(treeNodes); tn++ {
		tnInfo := utils.Split(treeNodes[tn], " = ")
		if tnInfo[0][2] == 'A' {
			currentNodes = append(currentNodes, tnInfo[0])
		}

		paths := utils.Split(strings.Trim(tnInfo[1], "()"), ", ")
		Tree[tnInfo[0]] = Node{paths[0], paths[1]}
	}

	var finished bool = false
	var lenPath uint64 = 0
	for !finished {
		var i int
		for i = 0; ; i++ {
			// fmt.Println(currentNodes)
			if i == len(instructions) {
				i = 0
			}
			if instructions[i] == "L" {
				for cn := 0; cn < len(currentNodes); cn++ {
					currentNodes[cn] = Tree[currentNodes[cn]].left
				}
				lenPath++
			} else if instructions[i] == "R" {
				for cn := 0; cn < len(currentNodes); cn++ {
					currentNodes[cn] = Tree[currentNodes[cn]].right
				}
				lenPath++
			}

			var finishedPath int
			for cn := 0; cn < len(currentNodes); cn++ {
				if currentNodes[cn][2] == 'Z' {
					finishedPath++
				}
			}
			if finishedPath == len(currentNodes) {
				finished = true
				break
			}
			// } else if finishedPath > 3 || lenPath%1000000000 == 0 {
			// 	fmt.Printf("finishedPath = %v | lenPath = %v | current = %v\n", finishedPath, lenPath, currentNodes)
			// }
		}
	}

	// fmt.Println(lenPath)
	PrintSolution(lenPath)
}
