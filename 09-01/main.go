package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func deduceSequence(inputNumbers []int) [][]int {
	var nextSequence []int

	for idx, num := range inputNumbers {
		if idx == 0 {
			continue
		}

		prevNum := inputNumbers[idx-1]
		nextSequence = append(nextSequence, num-prevNum)
	}

	isAllZero := true
	for _, num := range nextSequence {
		if num != 0 {
			isAllZero = false
			break
		}
	}

	var recurseNextSequence [][]int
	recurseNextSequence = append(recurseNextSequence, inputNumbers)

	if !isAllZero {
		recurseNextSequence = append(recurseNextSequence, deduceSequence(nextSequence)...)
	}

	return recurseNextSequence
}

func main() {
	digitsOnlyRegex := regexp.MustCompile("-?\\d+")

	file, _ := os.Open("./input/09.txt")
	defer file.Close()

	totalSum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		currentLine := scanner.Text()

		var currentNumbers []int
		for _, digitStr := range digitsOnlyRegex.FindAllString(currentLine, -1) {
			digitInt, _ := strconv.Atoi(digitStr)
			currentNumbers = append(currentNumbers, digitInt)
		}

		sequence := deduceSequence(currentNumbers)
		for _, currentSeq := range sequence {
			lastNum := currentSeq[len(currentSeq)-1]
			totalSum += lastNum
		}
	}

	fmt.Println(totalSum)
}
