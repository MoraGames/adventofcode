package y2023

import (
	"sort"

	"github.com/MoraGames/adventofcode/utils"
)

type Hand7 struct {
	cards    []int
	typology int
	bid      int
}

var JValue int = 11

func CardToInt(card string) int {
	if card == "A" {
		return 14
	} else if card == "K" {
		return 13
	} else if card == "Q" {
		return 12
	} else if card == "J" {
		return JValue //11 for D7P1, 0 for D7P2
	} else if card == "T" {
		return 10
	} else {
		return utils.StoI(card)
	}
}

func D7P1() {
	day, part = 7, 1
	JValue = 11
	var hands []Hand7

	lines := utils.SplitLines(ReadFile())
	for l := 0; l < len(lines); l++ {
		handInfo := utils.SplitSpaces(lines[l])
		cardsStr := utils.Split(handInfo[0], "")
		bid := utils.StoI(handInfo[1])

		// get typology
		cardsAmount := map[string]int{
			"A": 0, "K": 0, "Q": 0, "J": 0,
			"T": 0, "9": 0, "8": 0, "7": 0, "6": 0, "5": 0, "4": 0, "3": 0, "2": 0,
		}
		var cards []int
		for c := 0; c < len(cardsStr); c++ {
			cardsAmount[cardsStr[c]]++
			cards = append(cards, CardToInt(cardsStr[c]))
		}
		var typology int
		setOfFive, setOfFour, setOfThree, setOfTwo, setOfOne := 0, 0, 0, 0, 0
		for _, v := range cardsAmount {
			if v == 5 {
				setOfFive++
			} else if v == 4 {
				setOfFour++
			} else if v == 3 {
				setOfThree++
			} else if v == 2 {
				setOfTwo++
			} else if v == 1 {
				setOfOne++
			}
		}
		if setOfFive == 1 {
			typology = 7
		} else if setOfFour == 1 && setOfOne == 1 {
			typology = 6
		} else if setOfThree == 1 && setOfTwo == 1 {
			typology = 5
		} else if setOfThree == 1 && setOfOne == 2 {
			typology = 4
		} else if setOfTwo == 2 && setOfOne == 1 {
			typology = 3
		} else if setOfTwo == 1 && setOfOne == 3 {
			typology = 2
		} else if setOfOne == 5 {
			typology = 1
		}

		hands = append(hands, Hand7{cards, typology, bid})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].typology == hands[j].typology {
			for c := 0; c < len(hands[i].cards); c++ {
				if hands[i].cards[c] != hands[j].cards[c] {
					return hands[i].cards[c] < hands[j].cards[c]
				}
			}
		}
		return hands[i].typology < hands[j].typology
	})
	var totalWin uint64
	for h := 0; h < len(hands); h++ {
		// fmt.Printf("hands[%v]: type = %v | cards = %v | bid = %v | rank = %v\n", h, hands[h].typology, hands[h].cards, hands[h].bid, h+1)
		totalWin += uint64(hands[h].bid) * uint64(h+1)
	}
	// fmt.Printf("Total win: %v\n", totalWin)
	PrintSolution(totalWin)
}

func D7P2() {
	day, part = 7, 2
	JValue = 0
	var hands []Hand7

	lines := utils.SplitLines(ReadFile())
	for l := 0; l < len(lines); l++ {
		handInfo := utils.SplitSpaces(lines[l])
		cardsStr := utils.Split(handInfo[0], "")
		bid := utils.StoI(handInfo[1])

		// get typology
		cardsAmount := map[string]int{
			"A": 0, "K": 0, "Q": 0, "J": 0,
			"T": 0, "9": 0, "8": 0, "7": 0, "6": 0, "5": 0, "4": 0, "3": 0, "2": 0,
		}
		var cards []int
		for c := 0; c < len(cardsStr); c++ {
			cardsAmount[cardsStr[c]]++
			cards = append(cards, CardToInt(cardsStr[c]))
		}
		var typology int
		numOfJokers := 0
		setOfFive, setOfFour, setOfThree, setOfTwo, setOfOne := 0, 0, 0, 0, 0
		for k, v := range cardsAmount {
			if k == "J" {
				numOfJokers = v
			} else {
				if v == 5 {
					setOfFive++
				} else if v == 4 {
					setOfFour++
				} else if v == 3 {
					setOfThree++
				} else if v == 2 {
					setOfTwo++
				} else if v == 1 {
					setOfOne++
				}
			}
		}
		if setOfFive == 1 || setOfFour == 1 && numOfJokers == 1 || setOfThree == 1 && numOfJokers == 2 || setOfTwo == 1 && numOfJokers == 3 || setOfOne == 1 && numOfJokers == 4 || numOfJokers == 5 {
			typology = 7
		} else if setOfFour == 1 && setOfOne == 1 || setOfThree == 1 && setOfOne == 1 && numOfJokers == 1 || setOfTwo == 1 && setOfOne == 1 && numOfJokers == 2 || setOfOne == 2 && numOfJokers == 3 || setOfOne == 1 && numOfJokers == 4 {
			typology = 6
		} else if setOfThree == 1 && setOfTwo == 1 || setOfTwo == 2 && numOfJokers == 1 {
			typology = 5
		} else if setOfThree == 1 && setOfOne == 2 || setOfTwo == 1 && setOfOne == 2 && numOfJokers == 1 || setOfOne == 3 && numOfJokers == 2 {
			typology = 4
		} else if setOfTwo == 2 && setOfOne == 1 {
			typology = 3
		} else if setOfTwo == 1 && setOfOne == 3 || setOfOne == 4 && numOfJokers == 1 {
			typology = 2
		} else if setOfOne == 5 {
			typology = 1
		}

		hands = append(hands, Hand7{cards, typology, bid})
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].typology == hands[j].typology {
			for c := 0; c < len(hands[i].cards); c++ {
				if hands[i].cards[c] != hands[j].cards[c] {
					return hands[i].cards[c] < hands[j].cards[c]
				}
			}
		}
		return hands[i].typology < hands[j].typology
	})
	var totalWin uint64
	for h := 0; h < len(hands); h++ {
		// fmt.Printf("hands[%v]: type = %v | cards = %v | bid = %v | rank = %v\n", h, hands[h].typology, hands[h].cards, hands[h].bid, h+1)
		totalWin += uint64(hands[h].bid) * uint64(h+1)
	}
	// fmt.Printf("Total win: %v\n", totalWin)
	PrintSolution(totalWin)
}

//250788629 too low
//251356897 too low
//251735672 correct
