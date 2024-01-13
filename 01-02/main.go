package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func convertToInt(input string) int {
	num, err := strconv.ParseInt(input, 10, 0)
	if err != nil {
		fmt.Println("error converting digit", err)
		return 0
	}

	return int(num)
}

func main() {
	file, err := os.Open("./input/01.txt")
	if err != nil {
		fmt.Println("error reading input.txt file", err)
		return
	}
	defer file.Close()

	sum := 0
	nonDigitRegex := regexp.MustCompile("[^\\d]")
	numberTextArr := strings.Split("one two three four five six seven eight nine", " ")

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		patchedLine := line

		// mostly not optimized?
		for idx, numberText := range numberTextArr {
			if !strings.Contains(patchedLine, numberText) {
				continue
			}

			numberInt := idx + 1
			targetStr := numberText[0:1] + strconv.Itoa(numberInt) + numberText[1:]
			patchedLine = strings.ReplaceAll(patchedLine, numberText, targetStr)
		}

		fmt.Println("-----")
		fmt.Println("before ->", line)
		fmt.Println(" after ->", patchedLine)

		digitOnlyLine := nonDigitRegex.ReplaceAllString(patchedLine, "")
		digitArr := strings.Split(digitOnlyLine, "")
		fmt.Printf("input %v ->", digitArr)

		if len(digitArr) == 0 {
			fmt.Print("nothing found")
		} else if len(digitArr) == 1 {
			num := convertToInt(digitArr[0])
			num = (num * 10) + num
			fmt.Printf(" %d", num)
			sum += num
		} else {
			firstNumber := convertToInt(digitArr[0]) * 10
			lastNumber := convertToInt(digitArr[len(digitArr)-1])
			num := firstNumber + lastNumber
			fmt.Printf(" %d", num)
			sum += num
		}

		fmt.Println()
	}

	fmt.Println("total", sum)
}
