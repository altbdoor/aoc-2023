package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// GPT-4: Function to calculate the Greatest Common Divisor (GCD) using the Euclidean algorithm
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// GPT-4: Function to calculate the Lowest Common Multiple (LCM) using the GCD
func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {
	file, _ := os.Open("./input/08.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	charsOnlyRegex := regexp.MustCompile("[A-Z0-9]{3}")
	guide := ""
	keymap := make(map[string]([2]string))
	var startKeys []string

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

		if strings.HasSuffix(charsGroup[0], "A") {
			startKeys = append(startKeys, charsGroup[0])
		}
	}
	fmt.Println("finish parsing file")

	var allStepCounts []int

	for _, startKey := range startKeys {
		stepCount := 0
		nextKey := startKey
		nextGuide := ""
		lenGuide := len(guide)

		for !strings.HasSuffix(nextKey, "Z") {
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

		fmt.Printf("%s took %d steps\n", startKey, stepCount)
		allStepCounts = append(allStepCounts, stepCount)
	}

	lcmResult := 0
	for idx, stepCount := range allStepCounts {
		if idx == 0 {
			lcmResult = stepCount
			continue
		}

		lcmResult = lcm(lcmResult, stepCount)
	}

	fmt.Println(lcmResult)
}
