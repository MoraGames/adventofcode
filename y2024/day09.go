package y2024

import (
	"github.com/MoraGames/adventofcode/utils"
)

func D9P1() {
	day, part = 9, 1
	disk := ReadFile()

	var fileID int
	diskMap := make([]string, 0)
	for i, dim := range disk {
		if i%2 == 0 {
			//file
			dimension := utils.StoI(string(dim))
			for d := 0; d < dimension; d++ {
				diskMap = append(diskMap, utils.ItoS(fileID))
			}
			fileID++
		} else {
			//free
			dimension := utils.StoI(string(dim))
			for d := 0; d < dimension; d++ {
				diskMap = append(diskMap, ".")
			}
		}
	}

	j := len(diskMap) - 1
	for i := 0; i < len(diskMap) && i <= j; i++ {
		if diskMap[i] == "." {
			for diskMap[j] == "." {
				j--
			}
			diskMap = utils.Swap(diskMap, i, j)
			j--
		}
	}

	var checksum int
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] != "." {
			checksum += utils.StoI(diskMap[i]) * i
		}
	}

	PrintSolution(checksum)
}

func D9P2() {
	day, part = 9, 2
	disk := ReadFile()

	var fileID int
	diskMap := make([]string, 0)
	for i, dim := range disk {
		if i%2 == 0 {
			//file
			dimension := utils.StoI(string(dim))
			for d := 0; d < dimension; d++ {
				diskMap = append(diskMap, utils.ItoS(fileID))
			}
			fileID++
		} else {
			//free
			dimension := utils.StoI(string(dim))
			for d := 0; d < dimension; d++ {
				diskMap = append(diskMap, ".")
			}
		}
	}

	var fileLength int
	for i := len(diskMap) - 1; i > 0; i-- {
		if diskMap[i] != "." {
			fileLength++
			if diskMap[i] != diskMap[i-1] {
				findSpaceAndMoveFile(diskMap, fileLength, i)
				fileLength = 0
			}
		} else if diskMap[i] == "." && fileLength != 0 {
			findSpaceAndMoveFile(diskMap, fileLength, i)
			fileLength = 0
		}
	}

	var checksum int
	for i := 0; i < len(diskMap); i++ {
		if diskMap[i] != "." {
			checksum += utils.StoI(diskMap[i]) * i
		}
	}

	PrintSolution(checksum)
}

func findSpaceAndMoveFile(diskMap []string, length int, startIndex int) {
	freeSpaceLength := 0
	freeSpaceStartIndex := -1
	for i := 1; i <= startIndex; i++ {
		if diskMap[i-1] != "." && diskMap[i] == "." {
			if freeSpaceLength >= length {
				for j := 0; j < length; j++ {
					utils.Swap(diskMap, freeSpaceStartIndex+j, startIndex+j)
				}
			}
			freeSpaceStartIndex = i
			freeSpaceLength = 1
			continue
		}

		if diskMap[i-1] == "." && diskMap[i] != "." {
			if freeSpaceLength >= length {
				for j := 0; j < length; j++ {
					utils.Swap(diskMap, freeSpaceStartIndex+j, startIndex+j)
				}
			}
			freeSpaceStartIndex = -1
			freeSpaceLength = 0
		}

		if diskMap[i] == "." {
			freeSpaceLength++
		}
	}
}
