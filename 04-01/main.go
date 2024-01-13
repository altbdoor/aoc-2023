package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input/04.txt")
	if err != nil {
		fmt.Println("error reading input.txt file", err)
		return
	}
	defer file.Close()

	sum := 0
	digitsRegex := regexp.MustCompile("\\d+")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " | ")

		winningStr := digitsRegex.FindAllString(lineSplit[0][9:], -1)
		var winningInt []int
		for _, win := range winningStr {
			tmpInt, _ := strconv.Atoi(win)
			winningInt = append(winningInt, tmpInt)
		}

		lineMatchCount := 0
		cardStr := digitsRegex.FindAllString(lineSplit[1], -1)

		for _, card := range cardStr {
			cardInt, _ := strconv.Atoi(card)

			if slices.Contains(winningInt, cardInt) {
				if lineMatchCount == 0 {
					lineMatchCount += 1
				} else {
					lineMatchCount *= 2
				}
			}
		}

		sum += lineMatchCount
	}

	fmt.Println(sum)
}
