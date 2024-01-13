package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("./input/02.txt")
	if err != nil {
		fmt.Println("error reading input.txt file", err)
		return
	}
	defer file.Close()

	powerSum := 0
	colorMatchRegex := regexp.MustCompile("(\\d+) (red|blue|green)")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		lineSplit := strings.Split(line, ": ")
		gameIdStr := strings.ReplaceAll(lineSplit[0], "Game ", "")
		gameIdInt, _ := strconv.Atoi(gameIdStr)
		fmt.Printf("Parsing game ID %d -> ", gameIdInt)

		colorSplit := strings.Split(lineSplit[1], "; ")
		minColors := map[string]int{
			"red":   1,
			"green": 1,
			"blue":  1,
		}

		for _, colorGroup := range colorSplit {
			colorMatches := colorMatchRegex.FindAllString(colorGroup, -1)

			for _, match := range colorMatches {
				colorSplit := strings.Split(match, " ")
				colorCount, _ := strconv.Atoi(colorSplit[0])
				colorName := colorSplit[1]

				if colorCount > minColors[colorName] {
					minColors[colorName] = colorCount
				}
			}
		}
		fmt.Printf("%v -> ", minColors)

		power := 1
		for _, colorCount := range minColors {
			power *= colorCount
		}
		fmt.Printf("%d", power)
		powerSum += power

		fmt.Println()
	}

	fmt.Println(powerSum)
}
