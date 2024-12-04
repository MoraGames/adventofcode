package y2022

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D13P1() {
	file := ReadFile(13)
	//file := testReadFile(13)

	fmt.Printf("file = %q\n\n", file)

	var orderedPairs, sumPairsIndex int
	pairs := utils.SplitDoubleLines(file)
	for p := 0; p < len(pairs); p++ {

		fmt.Printf("=== Pair %d ===\n", p+1)

		pkgs := utils.SplitLines(pairs[p])

		fmt.Printf("Pair = {\n\t%v\n\t%v\n}\n", pkgs[0], pkgs[1])
		comparison := comparePackage(pkgs[0], pkgs[1])

		switch comparison {
		case 1:
			fmt.Printf("\tLeft package is smaller\n")
			fmt.Printf("ORDERED")
			orderedPairs++
			sumPairsIndex += p + 1
		case 0:
			if len(splitList(pkgs[1])) >= len(splitList(pkgs[0])) {
				fmt.Printf("\tLeft package is equal to Right package\n")
				fmt.Printf("ORDERED")
				orderedPairs++
				sumPairsIndex += p + 1
			}
		case -1:
			fmt.Printf("\tRight package is smaller\n")
		}
		fmt.Printf("\n\n\n")
	}

	fmt.Printf("orderedPairs = %d\nsumPairsIndex = %d\n", orderedPairs, sumPairsIndex)
}

func D13P2() {
	file := ReadFile(13)
	//file := testReadFile(13)

	packages := strings.Split(strings.Join(strings.Split(file, "\r\n\r\n"), "\r\n")+"\r\n[[2]]\r\n[[6]]", "\r\n")

	for p := 0; p < len(packages); p++ {
		fmt.Printf("%v\n", packages[p])
	}

	for isDone := false; isDone == false; {
		isDone = true
		for p := 0; p < len(packages)-1; p++ {
			if comparePackage(packages[p], packages[p+1]) == -1 {
				packages[p], packages[p+1] = packages[p+1], packages[p]
				isDone = false
			}
		}
	}

	fmt.Printf("\n\n\n\n")

	var decoderKey = 1
	for p := 0; p < len(packages); p++ {
		fmt.Printf("%v\n", packages[p])
		if packages[p] == "[[2]]" || packages[p] == "[[6]]" {
			decoderKey *= p + 1
		}
	}

	fmt.Printf("\n\n\n\n")
	fmt.Printf("decoderKey = %d\n", decoderKey)
}

func comparePackage(package0, package1 string) int {
	fmt.Printf("Compare %q vs %q\n", package0, package1)

	ep0 := isEmptyList(package0)
	ep1 := isEmptyList(package1)

	pkg0 := splitList(package0)
	pkg1 := splitList(package1)

	var comparison int
	if ep0 == ep1 {
		if ep0 == false {
			fmt.Printf("\tsplitList pkgs[0] = %q\n", pkg0)
			fmt.Printf("\tsplitList pkgs[1] = %q\n", pkg1)

			for i := 0; i < len(pkg0) && i < len(pkg1); i++ {
				comparison = compareElement(pkg0[i], pkg1[i])
				if comparison != 0 {
					break
				}
			}
			if comparison == 0 {
				if len(pkg0) < len(pkg1) {
					fmt.Printf("\tLeft package run out of elements first\n")
					comparison = 1
				} else if len(pkg0) == len(pkg1) {
					fmt.Printf("\tLeft package is equal to Right package\n")
					comparison = 0
				} else {
					fmt.Printf("\tRight package run out of elements first\n")
					comparison = -1
				}
			}
		} else {
			fmt.Printf("\tLeft package is equal to Right package (are empty)\n")
			comparison = 0
		}
	} else {
		if ep0 == true {
			fmt.Printf("\tLeft run out of elements first (is empty)\n")
			comparison = 1
		} else {
			fmt.Printf("\tRight run out of elements first (is empty)\n")
			comparison = -1
		}
	}

	return comparison
}

func compareElement(element0, element1 string) int {
	fmt.Printf("\tCompare %q vs %q\n", element0, element1)

	var se0, se1 = isSimpleElement(element0), isSimpleElement(element1)

	fmt.Printf("\t\telement0 = %q (simple? %v)\n", element0, se0)
	fmt.Printf("\t\telement1 = %q (simple? %v)\n", element1, se1)

	if se0 == se1 {
		if se0 == true {
			//Both simpleElement
			elm0, err := strconv.Atoi(element0)
			if err != nil {
				panic(err)
			}
			elm1, err := strconv.Atoi(element1)
			if err != nil {
				panic(err)
			}

			if elm0 < elm1 {
				fmt.Printf("\t\t\tLeft side is smaller\n")
				return 1
			} else if elm0 == elm1 {
				fmt.Printf("\t\t\tLeft side is equal to Right side\n")
				return 0
			} else {
				fmt.Printf("\t\t\tRight side is smaller\n")
				return -1
			}
		} else {
			//Both listElement

			el0 := isEmptyList(element0)
			el1 := isEmptyList(element1)

			if el0 == el1 {
				if el0 == false {
					elm0 := splitList(element0)
					elm1 := splitList(element1)

					for i := 0; i < len(elm0) && i < len(elm1); i++ {
						comparison := compareElement(elm0[i], elm1[i])
						if comparison != 0 {
							return comparison
						}
					}
					if len(elm0) < len(elm1) {
						fmt.Printf("\t\t\tLeft run out of elements first\n")
						return 1
					} else if len(elm0) == len(elm1) {
						fmt.Printf("\t\t\tLeft list is equal to Right list\n")
						return 0
					} else {
						fmt.Printf("\t\t\tRight run out of elements first\n")
						return -1
					}
				} else {
					fmt.Printf("\t\t\tLeft list is equal to Right list (are empty)\n")
					return 0
				}
			} else {
				if el0 == true {
					fmt.Printf("\t\t\tLeft run out of elements first (is empty)\n")
					return 1
				} else {
					fmt.Printf("\t\t\tRight run out of elements first (is empty)\n")
					return -1
				}
			}
		}
	} else {
		//Convert the simpleElement into a listElement

		if se0 == false {
			if element1 == "" {
				return -1
			}
			fmt.Printf("\t\t\tConverting Right (%q)", element1)
			element1 = "[" + element1 + "]"
			fmt.Printf(" to %q\n", element1)
		} else {
			if element0 == "" {
				return 1
			}
			fmt.Printf("\t\t\tConverting Left (%q)", element0)
			element0 = "[" + element0 + "]"
			fmt.Printf(" to %q\n", element0)
		}

		//Both listElement
		return compareElement(element0, element1)
	}
}

func isSimpleElement(element string) bool {
	return !strings.ContainsAny(element, "[],")
}

func isEmptyList(element string) bool {
	return element == "[]"
}

func splitList(list string) []string {
	var slicedList []string
	var bracketsCounter, slicingIndex int = 0, 1

	for i := 1; i < len(list)-1; i++ {
		if list[i] == '[' {
			bracketsCounter++
		} else if list[i] == ']' {
			bracketsCounter--
		} else if list[i] == ',' && bracketsCounter == 0 {
			slicedList = append(slicedList, list[slicingIndex:i])
			slicingIndex = i + 1
		}
	}
	slicedList = append(slicedList, list[slicingIndex:len(list)-1])

	return slicedList
}
