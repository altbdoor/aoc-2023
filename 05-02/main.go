package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var digitRegex = regexp.MustCompile("\\d+")

func getDigitsFromString(input string) []int64 {
	var convertedInt []int64
	for _, rawStr := range digitRegex.FindAllString(input, -1) {
		val, _ := strconv.ParseInt(rawStr, 10, 64)
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
	seedRaw := getDigitsFromString(scanner.Text())
	scanner.Scan()

	var seedRanges []([2]int64)

	for idx, maxSeedRange := range seedRaw {
		if idx > 0 && idx%2 == 1 {
			prevSeed := seedRaw[idx-1]
			currentSeedRange := [2]int64{prevSeed, prevSeed + maxSeedRange - 1}
			seedRanges = append(seedRanges, currentSeedRange)
		}
	}

	// seedChecked := make([]bool, len(seedRanges))
	// for idx := range seedChecked {
	// 	seedChecked[idx] = false
	// }
	var seedIsChecked []bool

	// logText := "(seed %d) from %d to %d"

	var currentSeedRanges [][2]int64
	var nextSeedRanges [][2]int64

	for scanner.Scan() {
		currentText := scanner.Text()

		// when blank lines, we skip
		if currentText == "" {
			continue
		}

		// if string has map, its a new section
		if strings.Contains(currentText, "map") {
			if currentSeedRanges == nil {
				// initial run, pull from original seeds
				currentSeedRanges = append([][2]int64{}, seedRanges...)
			} else {
				// next runs, pull from next seeds
				currentSeedRanges = append([][2]int64{}, nextSeedRanges...)
				nextSeedRanges = nil
			}

			seedIsChecked = make([]bool, len(currentSeedRanges))
			for idx := range seedIsChecked {
				seedIsChecked[idx] = false
			}

			fmt.Println("section", currentText)
			continue
		}

		// otherwise, they are range text
		ranges := getDigitsFromString(currentText)

		// i need help with some english
		srcRange := ranges[1]
		destRange := ranges[0]
		rangeLength := ranges[2]

		currentSrcRange := [2]int64{srcRange, srcRange + rangeLength - 1}

		// fmt.Println(">> init", currentSeedRanges)
		for idx, seedRange := range currentSeedRanges {
			fmt.Println(">>>> seedRange", seedRange)
			fmt.Println(">>>> currentSrcRange", currentSrcRange)

			// seed was already checked, continue
			if seedIsChecked[idx] {
				fmt.Println("skip")
				continue
			}

			// out of range
			// [0-seed-1]............
			// ...........[0-range-1]
			//
			// ............[0-seed-1]
			// [0-range-1]...........
			if seedRange[1] < currentSrcRange[0] || seedRange[0] > currentSrcRange[1] {
				continue
			}

			seedIsChecked[idx] = true

			// so after this, we're confirmed for some form of overlap
			currentDestRange := [2]int64{destRange, destRange + rangeLength - 1}
			fmt.Println(">>>> currentDestRange", currentDestRange)

			// ....[0-seed-1]
			// [0-range-1]...
			// xxxx----------
			if currentSrcRange[0] < seedRange[0] {
				// fmt.Println("=== trim left ===")
				diff := seedRange[0] - currentDestRange[0]
				if diff < 0 {
					diff *= -1
				}

				endpoint := currentDestRange[0] + diff
				split := [2]int64{currentDestRange[0], endpoint - 1}
				nextSeedRanges = append(nextSeedRanges, split)
				currentDestRange[0] = endpoint
			}

			// [0-seed-1]....
			// ...[0-range-1]
			// ----------xxxx
			if seedRange[1] < currentSrcRange[1] {
				// fmt.Println("=== trim right ===")
				diff := seedRange[1] - currentDestRange[1]
				if diff < 0 {
					diff *= -1
				}

				midpoint := currentDestRange[1] - diff
				split := [2]int64{midpoint + 1, currentDestRange[1]}
				nextSeedRanges = append(nextSeedRanges, split)
				currentDestRange[1] = midpoint
			}

			// within range
			// [----0-seed-1----]
			// ....[0-range-1]...
			nextSeedRanges = append(nextSeedRanges, currentDestRange)
			fmt.Println(">>>> nextSeedRanges", nextSeedRanges)
		}

		// fmt.Println("currentSeedRanges", currentSeedRanges)
		// fmt.Println("nextSeedRanges", nextSeedRanges)
		fmt.Println(">> processed line...")
	}

	var lowestPossibleLocation int64 = 0
	for _, foo := range nextSeedRanges {
		if lowestPossibleLocation == 0 {
			lowestPossibleLocation = foo[0]
		} else if foo[0] < lowestPossibleLocation {
			lowestPossibleLocation = foo[0]
		}
	}
	fmt.Println(lowestPossibleLocation)
	// fmt.Println(slices.Min(seedInt))
}
