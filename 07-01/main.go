package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

var cards = []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}

type CardSet struct {
	cards string
	bid   int
	rank  int
}

func (c *CardSet) GetRank() int {
	var powers []int
	uniqueCards := ""

	for _, currentCard := range strings.Split(c.cards, "") {
		if !strings.Contains(uniqueCards, currentCard) {
			uniqueCards += currentCard
			currentPower := strings.Count(c.cards, currentCard)
			powers = append(powers, currentPower)
		}
	}

	if slices.Contains(powers, 5) {
		// five of a kind
		return 6
	} else if slices.Contains(powers, 4) {
		// four of a kind
		return 5
	} else if slices.Contains(powers, 3) {
		if slices.Contains(powers, 2) {
			// full house
			return 4
		}

		// three of a kind
		return 3
	} else if slices.Contains(powers, 2) {
		if len(powers) == 3 {
			// two pair
			return 2
		}

		// one pair
		return 1
	}

	return 0
}

func main() {
	file, _ := os.Open("./input/07.txt")
	defer file.Close()

	var cardSetList []CardSet

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentText := scanner.Text()
		splitText := strings.Split(currentText, " ")

		cardSet := CardSet{splitText[0], 0, 0}
		cardSet.bid, _ = strconv.Atoi(splitText[1])

		cardSet.rank = cardSet.GetRank()
		cardSetList = append(cardSetList, cardSet)
	}

	sort.Slice(cardSetList, func(i, j int) bool {
		if cardSetList[i].rank != cardSetList[j].rank {
			return cardSetList[i].rank < cardSetList[j].rank
		}

		for idx, iCard := range strings.Split(cardSetList[i].cards, "") {
			jCard := cardSetList[j].cards[idx : idx+1]

			if jCard == iCard {
				continue
			}

			iCardIdx := slices.Index(cards, iCard)
			jCardIdx := slices.Index(cards, jCard)

			return iCardIdx > jCardIdx
		}

		return false
	})

	sum := 0
	for idx, cardSet := range cardSetList {
		sum += (idx + 1) * cardSet.bid
	}
	fmt.Println(sum)
}
