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

	if len(q.items) > 3 {
		q.items = q.items[1:]
	}
}

func main() {
	file, err := os.Open("./input/03.txt")
	if err != nil {
		fmt.Println("error reading input.txt file", err)
		return
	}

	parseQueue := Queue{}
	lineCharCount := 0
	digitRegex := regexp.MustCompile("\\d+")
	symbolRegex := regexp.MustCompile("[^\\d\\.]")

	partNumbersSum := 0

	scanner := bufio.NewScanner(file)
	for {
		// manually handle scan, so we can insert a "mocked" last line
		hasNextLine := scanner.Scan()

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
		}

		parseQueue.Add(line)

		if len(parseQueue.items) < 3 {
			continue
		}

		// process starts here
		currentLine := parseQueue.items[1]

		digitMatches := digitRegex.FindAllStringIndex(currentLine, -1)
		if digitMatches == nil {
			continue
		}

		for _, digitMatchGroup := range digitMatches {
			digitStr := currentLine[digitMatchGroup[0]:digitMatchGroup[1]]
			digitInt, _ := strconv.Atoi(digitStr)
			fmt.Printf("Found %d -> ", digitInt)
			isDigitPart := false

			// check before, current, after for symbols
			for _, snip := range parseQueue.items {
				currentSnip := snip[digitMatchGroup[0]-1 : digitMatchGroup[1]+1]
				if symbolRegex.MatchString(currentSnip) {
					partNumbersSum += digitInt
					isDigitPart = true
					break
				}
			}

			if isDigitPart {
				fmt.Print("IS PART")
			} else {
				fmt.Print("IS NOT PART")
			}

			fmt.Println()
		}

		if !hasNextLine {
			break
		}
	}

	fmt.Println(partNumbersSum)

}
