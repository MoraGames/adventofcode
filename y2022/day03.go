package y2022

import (
	"fmt"

	"github.com/MoraGames/adventofcode/utils"
)

func D3P1() {
	backpacks := utils.SplitLines(ReadFile(3))

	var totPriority int
	for b := 0; b < len(backpacks); b += 3 {

		fmt.Printf("Backpack %d\n\t", b)

		var stop = false
		for i1 := 0; i1 < len(backpacks[b]) && !stop; i1++ {
			for i2 := 0; i2 < len(backpacks[b+1]) && !stop; i2++ {
				for i3 := 0; i3 < len(backpacks[b+2]) && !stop; i3++ {
					if backpacks[b][i1] == backpacks[b+1][i2] && backpacks[b+1][i2] == backpacks[b+2][i3] {
						fmt.Printf(" | SameType %c (%d)\n\t", backpacks[b][i1], int(backpacks[b][i1]))
						var priority int
						switch {
						case backpacks[b][i1] >= 'A' && backpacks[b][i1] <= 'Z':
							priority = int(backpacks[b][i1]) - 65 + 27
						case backpacks[b][i1] >= 'a' && backpacks[b][i1] <= 'z':
							priority = int(backpacks[b][i1]) - 97 + 1
						}
						fmt.Printf(" | Priority: %d\n", priority)
						totPriority += priority
						stop = true
					}
				}
			}
		}
	}

	fmt.Printf("total: %d\n", totPriority)
}

func D3P2() {
	backpacks := utils.SplitLines(ReadFile(3))

	var totPriority int
	for b := 0; b < len(backpacks); b++ {
		firstSpace := backpacks[b][:len(backpacks[b])/2]
		secondSpace := backpacks[b][len(backpacks[b])/2:]

		fmt.Printf("Backpack %d\n\t  |  FS: %v\n\t  |  SP: %v\n\t", b, firstSpace, secondSpace)

		var stop = false
		for iFS := 0; iFS < len(firstSpace) && !stop; iFS++ {
			for iSS := 0; iSS < len(secondSpace) && !stop; iSS++ {
				if firstSpace[iFS] == secondSpace[iSS] {
					fmt.Printf("SameType %c (%d)", firstSpace[iFS], int(firstSpace[iFS]))
					//Same Item's Type
					var priority int
					switch {
					case firstSpace[iFS] >= 'A' && firstSpace[iFS] <= 'Z':
						priority = int(firstSpace[iFS]) - 65 + 27
					case firstSpace[iFS] >= 'a' && firstSpace[iFS] <= 'z':
						priority = int(firstSpace[iFS]) - 97 + 1
					}
					fmt.Printf("\n\tPriority: %d\n", priority)
					totPriority += priority
					stop = true
				}
			}
		}
	}

	fmt.Printf("total: %d\n", totPriority)
}

//AAaa len = 4/2 = 2 AA
//AAaaa len = 5/2 = 2 AA
