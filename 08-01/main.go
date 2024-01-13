package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	file, _ := os.Open("./input/08.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	charsOnlyRegex := regexp.MustCompile("[A-Z]{3}")
	guide := ""
	keymap := make(map[string]([2]string))

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		if guide == "" {
			guide = line
			continue
		}

		charsGroup := charsOnlyRegex.FindAllString(line, 3)
		keymap[charsGroup[0]] = [2]string{charsGroup[1], charsGroup[2]}
	}
	fmt.Println("finish parsing file")

	stepCount := 0
	nextKey := "AAA"
	nextGuide := ""
	lenGuide := len(guide)

	for nextKey != "ZZZ" {
		nextGuideIndex := stepCount % lenGuide
		nextGuide = guide[nextGuideIndex : nextGuideIndex+1]

		nextPair, _ := keymap[nextKey]
		if nextGuide == "L" {
			nextKey = nextPair[0]
		} else {
			nextKey = nextPair[1]
		}

		stepCount++
	}

	fmt.Println(stepCount)
}
