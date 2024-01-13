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

var cards = []string{"A", "K", "Q", "T", "9", "8", "7", "6", "5", "4", "3", "2", "J"}

type CardSet struct {
	cards string
	bid   int
	rank  int
}

func getRank(cards string) int {
	var cardTypes []string
	var cardCount []int

	for _, currentCard := range strings.Split(cards, "") {
		idx := slices.Index(cardTypes, currentCard)
		if idx == -1 {
			cardTypes = append(cardTypes, currentCard)
			cardCount = append(cardCount, 1)
			continue
		}

		cardCount[idx]++
	}

	mostCountIdx := -1
	for idx, cardType := range cardTypes {
		if cardType == "J" {
			continue
		}

		if mostCountIdx == -1 {
			mostCountIdx = idx
		}

		if cardCount[idx] > cardCount[mostCountIdx] {
			mostCountIdx = idx
		}
	}

	if mostCountIdx == -1 {
		// its all jokers!
		return 6
	}

	if slices.Contains(cardTypes, "J") {
		fixedCards := strings.ReplaceAll(cards, "J", cardTypes[mostCountIdx])
		return getRank(fixedCards)
	}

	if slices.Contains(cardCount, 5) {
		// five of a kind
		return 6
	} else if slices.Contains(cardCount, 4) {
		// four of a kind
		return 5
	} else if slices.Contains(cardCount, 3) {
		if slices.Contains(cardCount, 2) {
			// full house
			return 4
		}

		// three of a kind
		return 3
	} else if slices.Contains(cardCount, 2) {
		if len(cardCount) == 3 {
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

		cardSet.rank = getRank(cardSet.cards)
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
