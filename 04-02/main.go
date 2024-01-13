package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strings"
)

func getWinningCardCountFromLine(index int, line string) []int {
	var nextLines []int
	lineSplit := strings.Split(line, " | ")

	winningCardCount := 0
	winningStr := digitRegex.FindAllString(lineSplit[0][9:], -1)
	cardStr := digitRegex.FindAllString(lineSplit[1], -1)

	for _, card := range cardStr {
		if slices.Contains(winningStr, card) {
			winningCardCount += 1
			nextLines = append(nextLines, index+winningCardCount)
		}
	}

	return nextLines
}

var digitRegex = regexp.MustCompile("\\d+")

func main() {
	file, _ := os.Open("./input/04.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	lineCount := 0
	nextLinesMap := make(map[int][]int)
	winningCardCountMap := make(map[int]int)

	for scanner.Scan() {
		lineCount++
		line := scanner.Text()

		nextLines := getWinningCardCountFromLine(lineCount, line)
		nextLinesMap[lineCount] = nextLines
	}

	// fmt.Println(nextLinesMap)

	for key := range nextLinesMap {
		winningCardCountMap[key] = 1
	}

	for key := range nextLinesMap {
		recurseSum(key, nextLinesMap, winningCardCountMap)
	}

	sum := 0
	for _, val := range winningCardCountMap {
		sum += val
	}

	fmt.Println(winningCardCountMap)
	fmt.Println(sum)
}

func recurseSum(currentIdx int, nextLinesMap map[int][]int, winningCardCountMap map[int]int) {
	for _, nextLine := range nextLinesMap[currentIdx] {
		winningCardCountMap[nextLine]++
		recurseSum(nextLine, nextLinesMap, winningCardCountMap)
	}
}
