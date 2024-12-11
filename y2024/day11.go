package y2024

import (
	"github.com/MoraGames/adventofcode/utils"
)

func D11P1() {
	day, part = 11, 1
	stoneList := utils.SliceStoI(utils.SplitSpaces(ReadFile()))

	for t := 0; t < 25; t++ {
		newStoneList := make([]int, 0)
		for i := 0; i < len(stoneList); i++ {
			if stoneList[i] == 0 {
				newStoneList = append(newStoneList, 1)
				continue
			}
			if strStoneNum := utils.ItoS(stoneList[i]); len(strStoneNum)%2 == 0 {
				newStoneList = append(newStoneList, utils.StoI(strStoneNum[:len(strStoneNum)/2]), utils.StoI(strStoneNum[len(strStoneNum)/2:]))
				continue
			}
			newStoneList = append(newStoneList, stoneList[i]*2024)
		}
		stoneList = newStoneList
	}

	PrintSolution(len(stoneList))
}

func D11P2() {
	day, part = 11, 2
	stoneMap := utils.NewCache[int]()
	for _, stone := range utils.SliceStoI(utils.SplitSpaces(ReadFile())) {
		stoneMap[stone] = 1
	}

	for b := 0; b < 75; b++ {
		newStoneMap := utils.NewCache[int]()
		for stone, count := range stoneMap {
			if stone == 0 {
				utils.CacheAdd(newStoneMap, 1, count)
				continue
			}
			if stoneStr := utils.ItoS(stone); len(stoneStr)%2 == 0 {
				utils.CacheAdd(newStoneMap, utils.StoI(stoneStr[:len(stoneStr)/2]), count)
				utils.CacheAdd(newStoneMap, utils.StoI(stoneStr[len(stoneStr)/2:]), count)
				continue
			}
			utils.CacheAdd(newStoneMap, stone*2024, count)
		}
		stoneMap = newStoneMap
	}

	var sum int
	for _, count := range stoneMap {
		sum += count
	}

	PrintSolution(sum)
}
