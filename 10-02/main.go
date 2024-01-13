package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Point struct {
	x int
	y int
}

const DIR_TOP = "top"
const DIR_BTM = "btm"
const DIR_LEFT = "left"
const DIR_RIGHT = "right"

var directionsMap = map[string]([]string){
	"J": []string{DIR_TOP, DIR_LEFT},
	"|": []string{DIR_TOP, DIR_BTM},
	"L": []string{DIR_TOP, DIR_RIGHT},
	"7": []string{DIR_BTM, DIR_LEFT},
	"F": []string{DIR_BTM, DIR_RIGHT},
	"-": []string{DIR_RIGHT, DIR_LEFT},
}

func opposeDir(input string) string {
	if input == DIR_TOP {
		return DIR_BTM
	} else if input == DIR_BTM {
		return DIR_TOP
	} else if input == DIR_LEFT {
		return DIR_RIGHT
	} else {
		return DIR_LEFT
	}
}

func isValid(direction, input string) bool {
	matcher := ""

	if direction == DIR_TOP {
		matcher = "|7F"
	} else if direction == DIR_BTM {
		matcher = "|LJ"
	} else if direction == DIR_LEFT {
		matcher = "-LF"
	} else if direction == DIR_RIGHT {
		matcher = "-7J"
	}

	return strings.Contains(matcher, input)
}

func deduceStartChar(lines [][]string, rowIdx, colIdx int) (string, string) {
	var (
		hasTop   = isValid(DIR_TOP, lines[rowIdx-1][colIdx])
		hasBtm   = isValid(DIR_BTM, lines[rowIdx+1][colIdx])
		hasLeft  = isValid(DIR_LEFT, lines[rowIdx][colIdx-1])
		hasRight = isValid(DIR_RIGHT, lines[rowIdx][colIdx+1])
	)

	if hasTop && hasLeft {
		return "J", DIR_TOP
	} else if hasTop && hasBtm {
		return "|", DIR_TOP
	} else if hasTop && hasRight {
		return "L", DIR_TOP
	} else if hasBtm && hasLeft {
		return "7", DIR_BTM
	} else if hasBtm && hasRight {
		return "F", DIR_BTM
	} else if hasLeft && hasRight {
		return "-", DIR_RIGHT
	}

	return ".", DIR_TOP
}

func main() {
	file, _ := os.Open("./input/10.txt")
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)

	lineIdx := 0
	startRowIdx := -1

	for scanner.Scan() {
		currentText := scanner.Text()

		if strings.Contains(currentText, "S") {
			startRowIdx = lineIdx
		}

		lines = append(lines, strings.Split(currentText, ""))
		lineIdx++
	}

	// deduce the start and next dir first
	startColIdx := slices.Index(lines[startRowIdx], "S")
	startChar, nextDir := deduceStartChar(lines, startRowIdx, startColIdx)
	lines[startRowIdx][startColIdx] = startChar

	nextRowIdx := startRowIdx
	nextColIdx := startColIdx

	// var walkedCoords []Point
	var walkedCoords []string

	for true {
		char := lines[nextRowIdx][nextColIdx]

		var (
			// btmCoord   = Point{x: nextColIdx, y: nextRowIdx + 1}
			// topCoord   = Point{x: nextColIdx, y: nextRowIdx - 1}
			// leftCoord  = Point{x: nextColIdx - 1, y: nextRowIdx}
			// rightCoord = Point{x: nextColIdx + 1, y: nextRowIdx}
			topCoord   = fmt.Sprintf("%d,%d", nextColIdx, nextRowIdx-1)
			btmCoord   = fmt.Sprintf("%d,%d", nextColIdx, nextRowIdx+1)
			leftCoord  = fmt.Sprintf("%d,%d", nextColIdx-1, nextRowIdx)
			rightCoord = fmt.Sprintf("%d,%d", nextColIdx+1, nextRowIdx)
		)

		if char == "J" {
			if nextDir == DIR_TOP {
				walkedCoords = append(walkedCoords, topCoord)
				nextRowIdx--
			} else if nextDir == DIR_LEFT {
				walkedCoords = append(walkedCoords, leftCoord)
				nextColIdx--
			}
		} else if char == "|" {
			if nextDir == DIR_TOP {
				walkedCoords = append(walkedCoords, topCoord)
				nextRowIdx--
			} else if nextDir == DIR_BTM {
				walkedCoords = append(walkedCoords, btmCoord)
				nextRowIdx++
			}
		} else if char == "L" {
			if nextDir == DIR_TOP {
				walkedCoords = append(walkedCoords, topCoord)
				nextRowIdx--
			} else if nextDir == DIR_RIGHT {
				walkedCoords = append(walkedCoords, rightCoord)
				nextColIdx++
			}
		} else if char == "7" {
			if nextDir == DIR_BTM {
				walkedCoords = append(walkedCoords, btmCoord)
				nextRowIdx++
			} else if nextDir == DIR_LEFT {
				walkedCoords = append(walkedCoords, leftCoord)
				nextColIdx--
			}
		} else if char == "F" {
			if nextDir == DIR_BTM {
				walkedCoords = append(walkedCoords, btmCoord)
				nextRowIdx++
			} else if nextDir == DIR_RIGHT {
				walkedCoords = append(walkedCoords, rightCoord)
				nextColIdx++
			}
		} else if char == "-" {
			if nextDir == DIR_LEFT {
				walkedCoords = append(walkedCoords, leftCoord)
				nextColIdx--
			} else if nextDir == DIR_RIGHT {
				walkedCoords = append(walkedCoords, rightCoord)
				nextColIdx++
			}
		}

		if nextRowIdx == startRowIdx && nextColIdx == startColIdx {
			break
		}

		nextChar := lines[nextRowIdx][nextColIdx]
		for _, dir := range directionsMap[nextChar] {
			if dir != opposeDir(nextDir) {
				nextDir = dir
				break
			}
		}
	}

	// sanitize
	for rowIdx, line := range lines {
		for colIdx := range line {
			currentCoord := fmt.Sprintf("%d,%d", colIdx, rowIdx)
			indexOfWalked := slices.Index(walkedCoords, currentCoord)

			// indexOfWalked := slices.IndexFunc(walkedCoords, func(searchPoint Point) bool {
			// 	return searchPoint.x == colIdx && searchPoint.y == rowIdx
			// })

			if indexOfWalked == -1 {
				lines[rowIdx][colIdx] = "."
			}
		}
	}

	// visualize
	visualize := true
	if visualize {
		visualFile, _ := os.Create("visual.txt")

		for _, line := range lines {
			visualFile.WriteString(strings.Join(line, ""))
			visualFile.WriteString("\n")
		}
	}

	sum := 0
	for rowIdx, line := range lines {
		startLineCount := false
		lineCount := 0
		fmt.Println("===================", rowIdx)

		for _, char := range line {
			fmt.Printf("%s", char)
			if !startLineCount && strings.Contains("J|7", char) {
				startLineCount = true
				fmt.Println("start count")
				continue
			}

			if startLineCount && strings.Contains("L|F", char) {
				startLineCount = false
				fmt.Println("end count")
				continue
			}

			if char == "-" {
				fmt.Println("dash")
				continue
			}

			if startLineCount {
				fmt.Println("count")
				lineCount++
			}
		}

		if startLineCount {
			lineCount = 0
		}

		fmt.Println(lineCount)
		sum += lineCount
	}

	fmt.Println(sum)
}
