package y2023

import (
	"strconv"
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

type Bag map[string]int
type Game struct {
	hands    []Hand2
	possible bool
}

type Hand2 struct {
	cubes    map[string]int
	possible bool
}

var gameBags = make(map[int]Bag)
var games = make(map[int]Game)

func D2P1() {
	// Read from input2.txt file
	day, part = 2, 1
	lines := utils.SplitLines(ReadFile())
	bag := Bag{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	// Get game infos
	sumOfGameIDs := 0
	for l := 0; l < len(lines); l++ {
		gameInfos := strings.Split(lines[l], ": ")

		// Get gameID
		gameID, err := strconv.Atoi(strings.Split(gameInfos[0], " ")[1])
		if err != nil {
			panic(err)
		}

		// Get game hands
		isGamePossible := true
		var hands []Hand2
		gameHands := strings.Split(gameInfos[1], "; ")
		// fmt.Printf("DEBUG >> gameHands = %+v\n", gameHands)
		// fmt.Printf("Game %v =\n", gameID)
		for h := 0; h < len(gameHands); h++ {

			// Get hand set
			hand := Hand2{map[string]int{"red": 0, "green": 0, "blue": 0}, true}
			handSets := strings.Split(gameHands[h], ", ")
			for s := 0; s < len(handSets); s++ {

				// Get color and quantity of a hand set
				handSet := strings.Split(handSets[s], " ")

				amount, err := strconv.Atoi(strings.Trim(handSet[0], "\n\r"))
				if err != nil {
					panic(err)
				}
				color := strings.Trim(handSet[1], "\n\r")

				hand.cubes[color] += amount
			}
			if hand.cubes["red"] > bag["red"] || hand.cubes["green"] > bag["green"] || hand.cubes["blue"] > bag["blue"] {
				hand.possible = false
				isGamePossible = false
			}
			// fmt.Printf(" |- Hand %v: %q\n |       %v: %+v (%v)\n", h+1, gameHands[h], h+1, hand.cubes, hand.possible)
			hands = append(hands, hand)
		}
		games[gameID] = Game{hands, isGamePossible}
		if isGamePossible {
			sumOfGameIDs += gameID
		}
		// fmt.Printf(" |- [%v] --> sum = %v\n", isGamePossible, sumOfGameIDs)
	}

	PrintSolution(sumOfGameIDs)
}

func D2P2() {
	// Read from input1.txt file
	day, part = 2, 2
	lines := utils.SplitLines(ReadFile())

	// Get game infos
	sumOfGamesBag := 0
	for l := 0; l < len(lines); l++ {
		gameInfos := strings.Split(lines[l], ": ")

		// Get gameID
		gameID, err := strconv.Atoi(strings.Split(gameInfos[0], " ")[1])
		if err != nil {
			panic(err)
		}

		// Get game hands
		isGamePossible := true
		var bag Bag = Bag{"red": -1, "green": -1, "blue": -1}
		var hands []Hand2
		gameHands := strings.Split(gameInfos[1], "; ")
		// fmt.Printf("DEBUG >> gameHands = %+v\n", gameHands)
		// fmt.Printf("Game %v =\n", gameID)
		for h := 0; h < len(gameHands); h++ {

			// Get hand set
			hand := Hand2{map[string]int{"red": 0, "green": 0, "blue": 0}, true}
			handSets := strings.Split(gameHands[h], ", ")
			for s := 0; s < len(handSets); s++ {

				// Get color and quantity of a hand set
				handSet := strings.Split(handSets[s], " ")

				amount, err := strconv.Atoi(strings.Trim(handSet[0], "\n\r"))
				if err != nil {
					panic(err)
				}
				color := strings.Trim(handSet[1], "\n\r")

				hand.cubes[color] += amount
			}

			if hand.cubes["red"] > bag["red"] {
				bag["red"] = hand.cubes["red"]
			}
			if hand.cubes["green"] > bag["green"] {
				bag["green"] = hand.cubes["green"]
			}
			if hand.cubes["blue"] > bag["blue"] {
				bag["blue"] = hand.cubes["blue"]
			}
			hands = append(hands, hand)
		}

		gameBags[gameID] = bag
		games[gameID] = Game{hands, isGamePossible}

		bagPower := bag["red"] * bag["green"] * bag["blue"]
		sumOfGamesBag += bagPower
	}

	PrintSolution(sumOfGamesBag)
}
