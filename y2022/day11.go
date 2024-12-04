package y2022

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Monkey struct {
	items     []int
	operation [2]string
	test      int
	results   [2]int

	inspectionsNumer int
}

func D11P1() {
	var monkeys []Monkey

	monkeysString := strings.Split(ReadFile(11), "\r\n\r\n")
	fmt.Printf("monkeysString  = %q\n\n", monkeysString)

	for m := 0; m < len(monkeysString); m++ {
		fmt.Printf("monkeysString[%d]  = %q\n\n", m, monkeysString[m])
		var monkey Monkey

		monkeyLines := strings.Split(monkeysString[m], "\r\n")
		fmt.Printf("monkeyLines  = %q\n\n", monkeyLines)

		for ml := 1; ml < len(monkeyLines); ml++ {
			mlSplitted := strings.Split(monkeyLines[ml], ": ")
			fmt.Printf("mlSplitted  = %q\n\n", mlSplitted)

			switch ml {
			case 1: //Starting items
				stringParameters := strings.Split(mlSplitted[1], ", ")
				for i := 0; i < len(stringParameters); i++ {
					parameter, err := strconv.Atoi(stringParameters[i])
					if err != nil {
						panic(err)
					}
					monkey.items = append(monkey.items, parameter)
				}
			case 2: //Operation
				parameters := strings.Split(mlSplitted[1], " ")
				fmt.Printf("parameters[case 2]  = %q\n\n", parameters)
				monkey.operation = [2]string{parameters[3], parameters[4]}
			case 3: //Test
				stringParameters := strings.Split(mlSplitted[1], " ")
				fmt.Printf("stringParameters[case 3]  = %q\n\n", stringParameters)
				parameter, err := strconv.Atoi(stringParameters[2])
				if err != nil {
					panic(err)
				}
				monkey.test = parameter
			case 4: //T=true
				stringParameters := strings.Split(mlSplitted[1], " ")
				fmt.Printf("stringParameters[case 4]  = %q\n\n", stringParameters)
				parameter, err := strconv.Atoi(stringParameters[3])
				if err != nil {
					panic(err)
				}
				monkey.results[0] = parameter
			case 5: //T=false
				stringParameters := strings.Split(mlSplitted[1], " ")
				fmt.Printf("stringParameters[case 5]  = %q\n\n", stringParameters)
				parameter, err := strconv.Atoi(stringParameters[3])
				if err != nil {
					panic(err)
				}
				monkey.results[1] = parameter
			}
		}

		monkeys = append(monkeys, monkey)
	}

	stopRound := 20
	for round := 0; round < stopRound; round++ {
		for m := 0; m < len(monkeys); m++ {
			times := len(monkeys[m].items)
			for t := 0; t < times; t++ {
				monkeys[m].inspectionsNumer++

				var item int
				item, monkeys[m].items = monkeys[m].items[0], monkeys[m].items[1:]
				doOperation(monkeys[m].operation, &item)
				item = int(math.Floor(float64(item) / 3))
				if item%monkeys[m].test == 0 {
					monkeys[monkeys[m].results[0]].items = append(monkeys[monkeys[m].results[0]].items, item)
				} else {
					monkeys[monkeys[m].results[1]].items = append(monkeys[monkeys[m].results[1]].items, item)
				}
			}
		}
	}

	max1InspectionsNumber, max2InspectionsNumber := -1, -1
	for m := 0; m < len(monkeys); m++ {
		if monkeys[m].inspectionsNumer > max1InspectionsNumber {
			max2InspectionsNumber = max1InspectionsNumber
			max1InspectionsNumber = monkeys[m].inspectionsNumer
		} else if monkeys[m].inspectionsNumer > max2InspectionsNumber {
			max2InspectionsNumber = monkeys[m].inspectionsNumer
		}
	}

	monkeyBusiness := max1InspectionsNumber * max2InspectionsNumber
	fmt.Printf("monkeyBusiness = %v\n", monkeyBusiness)
}

func D11P2() {
	globalCommonDivisor := 1
	var monkeys []Monkey

	monkeysString := strings.Split(ReadFile(11), "\r\n\r\n")
	fmt.Printf("monkeysString  = %q\n\n", monkeysString)

	for m := 0; m < len(monkeysString); m++ {
		fmt.Printf("monkeysString[%d]  = %q\n\n", m, monkeysString[m])
		var monkey Monkey

		monkeyLines := strings.Split(monkeysString[m], "\r\n")
		fmt.Printf("monkeyLines  = %q\n\n", monkeyLines)

		for ml := 1; ml < len(monkeyLines); ml++ {
			mlSplitted := strings.Split(monkeyLines[ml], ": ")
			fmt.Printf("mlSplitted  = %q\n\n", mlSplitted)

			switch ml {
			case 1: //Starting items
				stringParameters := strings.Split(mlSplitted[1], ", ")
				for i := 0; i < len(stringParameters); i++ {
					parameter, err := strconv.Atoi(stringParameters[i])
					if err != nil {
						panic(err)
					}
					monkey.items = append(monkey.items, parameter)
				}
			case 2: //Operation
				parameters := strings.Split(mlSplitted[1], " ")
				fmt.Printf("parameters[case 2]  = %q\n\n", parameters)
				monkey.operation = [2]string{parameters[3], parameters[4]}
			case 3: //Test
				stringParameters := strings.Split(mlSplitted[1], " ")
				fmt.Printf("stringParameters[case 3]  = %q\n\n", stringParameters)
				parameter, err := strconv.Atoi(stringParameters[2])
				if err != nil {
					panic(err)
				}
				monkey.test = parameter
				globalCommonDivisor *= parameter
			case 4: //T=true
				stringParameters := strings.Split(mlSplitted[1], " ")
				fmt.Printf("stringParameters[case 4]  = %q\n\n", stringParameters)
				parameter, err := strconv.Atoi(stringParameters[3])
				if err != nil {
					panic(err)
				}
				monkey.results[0] = parameter
			case 5: //T=false
				stringParameters := strings.Split(mlSplitted[1], " ")
				fmt.Printf("stringParameters[case 5]  = %q\n\n", stringParameters)
				parameter, err := strconv.Atoi(stringParameters[3])
				if err != nil {
					panic(err)
				}
				monkey.results[1] = parameter
			}
		}

		monkeys = append(monkeys, monkey)
	}

	stopRound := 10000
	for round := 0; round < stopRound; round++ {
		for m := 0; m < len(monkeys); m++ {
			times := len(monkeys[m].items)
			for t := 0; t < times; t++ {
				monkeys[m].inspectionsNumer++

				var item int
				item, monkeys[m].items = monkeys[m].items[0], monkeys[m].items[1:]
				doOperation(monkeys[m].operation, &item)
				item = item % globalCommonDivisor
				if item%monkeys[m].test == 0 {
					monkeys[monkeys[m].results[0]].items = append(monkeys[monkeys[m].results[0]].items, item)
				} else {
					monkeys[monkeys[m].results[1]].items = append(monkeys[monkeys[m].results[1]].items, item)
				}
			}
		}
	}

	max1InspectionsNumber, max2InspectionsNumber := -1, -1
	for m := 0; m < len(monkeys); m++ {
		if monkeys[m].inspectionsNumer > max1InspectionsNumber {
			max2InspectionsNumber = max1InspectionsNumber
			max1InspectionsNumber = monkeys[m].inspectionsNumer
		} else if monkeys[m].inspectionsNumer > max2InspectionsNumber {
			max2InspectionsNumber = monkeys[m].inspectionsNumer
		}
	}

	monkeyBusiness := max1InspectionsNumber * max2InspectionsNumber
	fmt.Printf("monkeyBusiness = %v\n", monkeyBusiness)
}

func doOperation(operation [2]string, item *int) {
	var operationValue int
	if operation[1] == "old" {
		operationValue = *item
	} else {
		var err error
		operationValue, err = strconv.Atoi(operation[1])
		if err != nil {
			panic(err)
		}
	}

	switch operation[0] {
	case "+":
		*item = *item + operationValue
	case "*":
		*item = *item * operationValue
	}
}
