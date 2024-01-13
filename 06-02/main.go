package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func getDigitsFromString(input string) int64 {
	var digitRegex = regexp.MustCompile("\\d+")
	combinedStr := ""
	for _, rawStr := range digitRegex.FindAllString(input, -1) {
		combinedStr += rawStr
	}

	convertedInt, _ := strconv.ParseInt(combinedStr, 10, 64)
	return convertedInt
}

func main() {
	file, _ := os.Open("./input/06.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timeVal := getDigitsFromString(scanner.Text())
	scanner.Scan()
	distVal := getDigitsFromString(scanner.Text())

	var currentIterCount int64 = 0
	var holdBtn int64 = 1

	for holdBtn < timeVal {
		currentDist := (timeVal - holdBtn) * holdBtn
		if currentDist > distVal {
			currentIterCount++
		}

		holdBtn++
	}

	fmt.Println(currentIterCount)
}
