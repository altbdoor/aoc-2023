package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Queue struct {
	items []string
}

func (q *Queue) Add(item string) {
	q.items = append(q.items, "."+item+".")

	if len(q.items) > 4 {
		q.items = q.items[1:]
	}
}

var digitRegex = regexp.MustCompile("\\d+")

func main() {
	file, err := os.Open("./input/03.txt")
	if err != nil {
		fmt.Println("error reading input.txt file", err)
		return
	}

	parseQueue := Queue{}
	lineCharCount := 0
	leftRightGearRegex := regexp.MustCompile("\\d+\\*\\d+")

	gearSum := 0
	lineCount := -2

	scanner := bufio.NewScanner(file)
	for {
		// manually handle scan, so we can insert a "mocked" last line
		hasNextLine := scanner.Scan()
		lineCount += 1

		line := ""
		if hasNextLine {
			// handles the current line, with left/right padding
			line = "." + scanner.Text() + "."
		} else {
			// handles the "mocked" last line padding
			line = strings.Repeat(".", lineCharCount)
		}

		if lineCharCount == 0 {
			lineCharCount = len(line)
			// handles the "mocked" first line padding
			parseQueue.Add(strings.Repeat(".", lineCharCount))
			parseQueue.Add(strings.Repeat(".", lineCharCount))
		}

		parseQueue.Add(line)

		if len(parseQueue.items) < 4 {
			continue
		}

		// process starts here
		firstLine := parseQueue.items[0]
		secondLine := parseQueue.items[1]
		thirdLine := parseQueue.items[2]
		fourthLine := parseQueue.items[3]
		// fmt.Println(firstLine)
		// fmt.Println(secondLine)
		// fmt.Println(thirdLine)
		// fmt.Println(fourthLine)

		// e.g. 100*100
		leftRightMatches := leftRightGearRegex.FindAllString(secondLine, -1)
		if leftRightMatches != nil {
			for _, leftRightMatch := range leftRightMatches {
				partNumbers := strings.Split(leftRightMatch, "*")
				partNumberInt1, _ := strconv.Atoi(partNumbers[0])
				partNumberInt2, _ := strconv.Atoi(partNumbers[1])
				fmt.Println(lineCount, "--", partNumberInt1, "*", partNumberInt2, "-- single line match")
				gearSum += partNumberInt1 * partNumberInt2
			}
		}

		secondLineDigitMatches := digitRegex.FindAllStringIndex(secondLine, -1)
		if secondLineDigitMatches == nil {
			continue
		}

		for _, secondLineDigitMatchGroup := range secondLineDigitMatches {
			secondLineDigitStr := secondLine[secondLineDigitMatchGroup[0]:secondLineDigitMatchGroup[1]]
			secondLineDigitInt, _ := strconv.Atoi(secondLineDigitStr)

			/*
				e.g.
				...*100
				100....
			*/
			checkLeftGear := secondLine[secondLineDigitMatchGroup[0]-1 : secondLineDigitMatchGroup[0]]
			if checkLeftGear == "*" {
				thirdLineDigit := checkNextLineBasedOnGearIndex(thirdLine, secondLineDigitMatchGroup[0]-1, false, false)
				if thirdLineDigit != -1 {
					fmt.Println(lineCount, "--", secondLineDigitInt, "*", thirdLineDigit, "-- left gear")
					gearSum += secondLineDigitInt * thirdLineDigit
					continue
				}
			}

			/*
				e.g.
				100*...
				....100
			*/
			checkRightGear := secondLine[secondLineDigitMatchGroup[1] : secondLineDigitMatchGroup[1]+1]
			if checkRightGear == "*" {
				thirdLineDigit := checkNextLineBasedOnGearIndex(thirdLine, secondLineDigitMatchGroup[1]+1, false, false)
				if thirdLineDigit != -1 {
					fmt.Println(lineCount, "--", secondLineDigitInt, "*", thirdLineDigit, "-- right gear")
					gearSum += secondLineDigitInt * thirdLineDigit
					continue
				}
			}

			/*
				e.g.
				....100
				100*...
			*/
			checkLowerLeftGear := thirdLine[secondLineDigitMatchGroup[0]-1 : secondLineDigitMatchGroup[0]]
			if checkLowerLeftGear == "*" {
				thirdLineDigit := checkNextLineBasedOnGearIndex(thirdLine, secondLineDigitMatchGroup[0]-1, false, false)
				if thirdLineDigit != -1 {
					fmt.Println(lineCount, "--", secondLineDigitInt, "*", thirdLineDigit, "-- lower left gear")
					gearSum += secondLineDigitInt * thirdLineDigit
					continue
				}
			}

			/*
				e.g.
				100....
				...*100
			*/
			checkLowerRightGear := thirdLine[secondLineDigitMatchGroup[1] : secondLineDigitMatchGroup[1]+1]
			if checkLowerRightGear == "*" {
				thirdLineDigit := checkNextLineBasedOnGearIndex(thirdLine, secondLineDigitMatchGroup[1], false, false)
				if thirdLineDigit != -1 {
					fmt.Println(lineCount, "--", secondLineDigitInt, "*", thirdLineDigit, "-- lower right gear")
					gearSum += secondLineDigitInt * thirdLineDigit
					continue
				}
			}

			/*
				e.g.
				...*...
				100.100
			*/
			checkUpperBumpGear := firstLine[secondLineDigitMatchGroup[1] : secondLineDigitMatchGroup[1]+1]
			if checkUpperBumpGear == "*" {
				secondLineDigitFound := checkNextLineBasedOnGearIndex(secondLine, secondLineDigitMatchGroup[1], true, false)
				if secondLineDigitFound != -1 {
					fmt.Println(lineCount, "--", secondLineDigitInt, "*", secondLineDigitFound, "-- upper bump gear")
					gearSum += secondLineDigitInt * secondLineDigitFound
					continue
				}
			}

			/*
				e.g.
				100.100
				...*...
			*/
			checkLowerBumpGear := thirdLine[secondLineDigitMatchGroup[1] : secondLineDigitMatchGroup[1]+1]
			if checkLowerBumpGear == "*" {
				secondLineDigitFound := checkNextLineBasedOnGearIndex(secondLine, secondLineDigitMatchGroup[1], true, false)
				if secondLineDigitFound != -1 {
					fmt.Println(lineCount, "--", secondLineDigitInt, "*", secondLineDigitFound, "-- lower bump gear")
					gearSum += secondLineDigitInt * secondLineDigitFound
					continue
				}
			}

			thirdLineSeek := thirdLine[secondLineDigitMatchGroup[0]-1 : secondLineDigitMatchGroup[1]+1]
			thirdLineGearIndex := strings.Index(thirdLineSeek, "*")

			if thirdLineGearIndex == -1 {
				continue
			}

			thirdLineGearIndex += secondLineDigitMatchGroup[0] - 1

			fourthLineDigitInt := checkNextLineBasedOnGearIndex(fourthLine, thirdLineGearIndex, false, false)
			if fourthLineDigitInt != -1 {
				fmt.Println(lineCount, "--", secondLineDigitInt, "*", fourthLineDigitInt, "-- fourth line match")
				gearSum += secondLineDigitInt * fourthLineDigitInt
			}
		}

		fmt.Println("=====")

		if !hasNextLine {
			break
		}
	}

	fmt.Println(gearSum)

}

func checkNextLineBasedOnGearIndex(nextLine string, gearIndex int, isRightOnly bool, debug bool) int {
	digitMatches := digitRegex.FindAllStringIndex(nextLine, -1)
	if digitMatches == nil {
		return -1
	}

	for _, matchGroup := range digitMatches {
		if !isRightOnly && gearIndex >= matchGroup[0] && gearIndex <= matchGroup[1] {
			/*
				gear is inside the number index range
				100
				.*.
				100
			*/
			digitStr := nextLine[matchGroup[0]:matchGroup[1]]
			digitInt, _ := strconv.Atoi(digitStr)
			return digitInt

		} else if gearIndex+1 == matchGroup[0] {
			/*
				backslash pattern, add one because start index exclusive
				100....
				...*...
				....100
			*/
			digitStr := nextLine[matchGroup[0]:matchGroup[1]]
			digitInt, _ := strconv.Atoi(digitStr)
			return digitInt

		} else if !isRightOnly && gearIndex == matchGroup[1] {
			/*
				forward slash pattern, no add because end index inclusive
				....100
				...*...
				100....
			*/
			digitStr := nextLine[matchGroup[0]:matchGroup[1]]
			digitInt, _ := strconv.Atoi(digitStr)
			return digitInt
		}
	}

	return -1
}
