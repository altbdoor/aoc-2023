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

var digitRegex = regexp.MustCompile("\\d+")

func getDigitsFromString(input string) []int {
	var convertedInt []int
	for _, rawStr := range digitRegex.FindAllString(input, -1) {
		val, _ := strconv.Atoi(rawStr)
		convertedInt = append(convertedInt, val)
	}

	return convertedInt
}

func main() {
	file, _ := os.Open("./input/05.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// first line is seeds
	scanner.Scan()
	seedInt := getDigitsFromString(scanner.Text())
	scanner.Scan()

	seedChecked := make([]bool, len(seedInt))
	for idx := range seedChecked {
		seedChecked[idx] = false
	}

	logText := "(seed %d) from %d to %d"

	for scanner.Scan() {
		currentText := scanner.Text()

		// when blank lines, we skip
		if currentText == "" {
			continue
		}

		// if string has map, its a new section
		if strings.Contains(currentText, "map") {
			for idx := range seedChecked {
				seedChecked[idx] = false
			}

			fmt.Println("section", currentText)
			continue
		}

		// otherwise, they are range text
		ranges := getDigitsFromString(currentText)
		srcRange := ranges[1]
		destRange := ranges[0]
		rangeLength := ranges[2]

		minSrcRange := srcRange
		maxSrcRange := srcRange + rangeLength - 1

		for idx, seed := range seedInt {
			// seed was already checked, continue
			if seedChecked[idx] {
				continue
			}

			// if seed within range, recalculate the dest value
			if seed >= minSrcRange && seed <= maxSrcRange {
				seedChecked[idx] = true

				srcIndex := seed - minSrcRange
				seedInt[idx] = destRange + srcIndex
				fmt.Printf(logText, idx, seed, seedInt[idx])
				fmt.Println()
			}
		}

		fmt.Println(">> processed line...")
	}

	// fmt.Println(seedInt)
	fmt.Println(slices.Min(seedInt))
}
