package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getDigitsFromString(input string) []int {
	var digitRegex = regexp.MustCompile("\\d+")
	var convertedInt []int
	for _, rawStr := range digitRegex.FindAllString(input, -1) {
		val, _ := strconv.Atoi(rawStr)
		convertedInt = append(convertedInt, val)
	}

	return convertedInt
}

func main() {
	file, _ := os.Open("./input/06.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timeVals := getDigitsFromString(scanner.Text())
	scanner.Scan()
	distVals := getDigitsFromString(scanner.Text())

	sum := 1

	for idx, timeVal := range timeVals {
		distVal := distVals[idx]
		currentIterCount := 0
		holdBtn := 1

		for holdBtn < timeVal {
			currentDist := (timeVal - holdBtn) * holdBtn
			if currentDist > distVal {
				currentIterCount++
			}

			holdBtn++
		}

		fmt.Printf("Found %d ways\n", currentIterCount)
		sum *= currentIterCount

	}
	fmt.Println(sum)
}
