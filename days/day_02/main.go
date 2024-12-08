package day_02

import (
	"fmt"
	"strconv"
	"strings"
)

// Run function of the daily challenge
func Run(input []string, mode int) {
	if mode == 1 || mode == 3 {
		fmt.Printf("Part one: %v\n", Part1(input))
	}
	if mode == 2 || mode == 3 {
		fmt.Printf("Part two: %v\n", Part2(input))
	}
}

func lineStringToInt(lineStr string) []int {
	var line []int
	for _, char := range strings.Split(lineStr, " ") {
		number, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}
		line = append(line, number)
	}
	return line
}

func IsValidRow(row []int) bool {
	isDecreasing := row[0] > row[1]

	if isDecreasing {
		isSafe := true
		for index, number := range row {
			if index > 0 {
				diff := number - row[index-1]
				if diff > 0 {
					// different directions
					isSafe = false
					break
				}
				if diff < -3 || diff == 0 {
					isSafe = false
					break
				}
			}
		}
		return isSafe
	} else {
		isSafe := true
		for index, number := range row {
			if index > 0 {
				diff := number - row[index-1]
				if diff < 0 {
					// different directions
					isSafe = false
					break
				}
				if diff > 3 || diff == 0 {
					isSafe = false
					break
				}
			}
		}
		return isSafe
	}
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	safeCount := 0

	for _, lineStr := range input {
		if IsValidRow(lineStringToInt(lineStr)) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	safeCount := 0

	for _, lineStr := range input {
		line := lineStringToInt(lineStr)
		// fmt.Printf("%v\n", line)
		if IsValidRow(line) {
			safeCount++
		} else {
			for i := range line {
				var tmp = make([]int, len(line))
				copy(tmp, line)
				var newLine = append(tmp[:i], tmp[i+1:]...)
				// fmt.Printf("try: %v (%v)\n", newLine, IsValidRow(newLine))
				if IsValidRow(newLine) {
					safeCount++
					break
				}
			}
		}

	}

	return strconv.Itoa(safeCount)
}
