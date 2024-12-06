package y2023

import (
	"strings"

	"github.com/MoraGames/adventofcode/utils"
)

func D4P1() {
	day, part = 4, 1
	cards := utils.SplitLines(ReadFile())

	var totalPoints int
	for c := 0; c < len(cards); c++ {
		parts := utils.Split(utils.Split(strings.ReplaceAll(cards[c], "  ", " "), ": ")[1], " | ")
		winningNumbers := utils.Split(parts[0], " ")
		possessionNumbers := utils.Split(parts[1], " ")

		finded := false
		winPtsValue := 0
		for pn := 0; pn < len(possessionNumbers); pn++ {
			for wn := 0; wn < len(winningNumbers); wn++ {
				// Logger.WithFields(logrus.Fields{
				// 	"posNum": possessionNumbers[pn],
				// 	"winNum": winningNumbers[wn],
				// }).Debug("Comparing posNum and winNum of card #" + ItoS(c+1))
				if utils.StoI(strings.TrimSpace(possessionNumbers[pn])) == utils.StoI(strings.TrimSpace(winningNumbers[wn])) {
					// Logger.WithFields(logrus.Fields{
					// 	"ptsVal": winPtsValue,
					// }).Debug("Found a match")
					if finded {
						winPtsValue *= 2
					} else {
						winPtsValue = 1
						finded = true
					}
				}
			}
		}
		// Logger.WithFields(logrus.Fields{
		// 	"ptsVal": winPtsValue,
		// }).Debug("Card #" + ItoS(c+1) + " points")
		totalPoints += winPtsValue
	}

	// Logger.WithFields(logrus.Fields{
	// 	"totalPoints": totalPoints,
	// }).Info("Total points")

	PrintSolution(totalPoints)
}

func D4P2() {
	day, part = 4, 2
	cards := utils.SplitLines(ReadFile())

	var totalScratchCards int
	for c := 0; c < len(cards); c++ {
		totalScratchCards++
		card := utils.Split(strings.ReplaceAll(strings.ReplaceAll(cards[c], "   ", " "), "  ", " "), ": ")

		cardNumber := utils.StoI(utils.SplitSpaces(card[0])[1])
		numbersParts := utils.Split(card[1], " | ")

		winningNumbers := utils.Split(numbersParts[0], " ")
		possessionNumbers := utils.Split(numbersParts[1], " ")

		var numOfCopy int
		for pn := 0; pn < len(possessionNumbers); pn++ {
			for wn := 0; wn < len(winningNumbers); wn++ {
				// Logger.WithFields(logrus.Fields{
				// 	"posNum": possessionNumbers[pn],
				// 	"winNum": winningNumbers[wn],
				// }).Debug("Comparing posNum and winNum of card #" + ItoS(cardNumber))
				if utils.StoI(strings.TrimSpace(possessionNumbers[pn])) == utils.StoI(strings.TrimSpace(winningNumbers[wn])) {
					numOfCopy++
				}
			}
		}

		var copyCards []string
		for i := 0; i < numOfCopy; i++ {
			copyCards = append(copyCards, cards[cardNumber+i])
		}

		// if len(cards)%500000 == 0 {
		// 	Logger.WithFields(logrus.Fields{
		// 		"totSchCrd": totalScratchCards,
		// 		"CrdcpyNum": numOfCopy,
		// 		"lenCpyCrd": len(copyCards),
		// 		"lenOriCrd": len(cards),
		// 	}).Debug("Card #" + ItoS(cardNumber) + " copies gifted")
		// }

		cards = append(cards, copyCards...)

		// Logger.WithFields(logrus.Fields{
		// 	"lenOriCrd": len(cards),
		// }).Debug("Card #" + ItoS(cardNumber) + " update")

		// time.Sleep(50 * time.Millisecond)
	}

	// Logger.WithFields(logrus.Fields{
	// 	"totSchCrd": totalScratchCards,
	// 	"totCrd":    len(cards),
	// }).Info("Total cards obtained")

	PrintSolution(totalScratchCards)
}
