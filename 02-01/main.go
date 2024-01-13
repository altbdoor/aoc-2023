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
	maxColors := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	file, err := os.Open("./input/02.txt")
	if err != nil {
		fmt.Println("error reading input.txt file", err)
		return
	}
	defer file.Close()

	idSum := 0
	colorMatchRegex := regexp.MustCompile("(\\d+) (red|blue|green)")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		lineSplit := strings.Split(line, ": ")
		gameIdStr := strings.ReplaceAll(lineSplit[0], "Game ", "")
		gameIdInt, _ := strconv.Atoi(gameIdStr)
		fmt.Printf("Parsing game ID %d -> ", gameIdInt)

		colorSplit := strings.Split(lineSplit[1], "; ")
		isParsingColorsOk := true

		for _, colorGroup := range colorSplit {
			colorMatches := colorMatchRegex.FindAllString(colorGroup, -1)

			for _, match := range colorMatches {
				colorSplit := strings.Split(match, " ")
				colorCount, _ := strconv.Atoi(colorSplit[0])
				colorName := colorSplit[1]

				if colorCount > maxColors[colorName] {
					isParsingColorsOk = false
					break
				}
			}

			if !isParsingColorsOk {
				break
			}
		}

		if isParsingColorsOk {
			idSum += gameIdInt
			fmt.Print("OK")
		} else {
			fmt.Print("NOK")
		}

		fmt.Println()
	}

	fmt.Println(idSum)
}
