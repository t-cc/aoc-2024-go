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

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	safeCount := 0

	for _, lineStr := range input {
		var line = []int{}
		for _, char := range strings.Split(lineStr, " ") {
			number, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			line = append(line, number)
		}

		isDecreasing := line[0] > line[1]

		if isDecreasing {
			isSafe := true
			for index, number := range line {
				if index > 0 {
					diff := number - line[index-1]
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
			if isSafe {
				safeCount++
			}
		} else {
			isSafe := true
			for index, number := range line {
				if index > 0 {
					diff := number - line[index-1]
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
			if isSafe {
				safeCount++
			}
		}
	}

	return strconv.Itoa(safeCount)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	safeCount := 0

	for _, lineStr := range input {
		var line = []int{}
		for _, char := range strings.Split(lineStr, " ") {
			number, err := strconv.Atoi(char)
			if err != nil {
				panic(err)
			}
			line = append(line, number)
		}

		isDecreasing := line[0] > line[1]

		if isDecreasing {
			errorCount := 0
			for index, number := range line {
				if index > 0 {
					diff := number - line[index-1]
					if diff > 0 {
						// different directions
						errorCount++
					}
					if diff < -3 || diff == 0 {
						errorCount++
						break
					}
				}
			}
			if errorCount <= 1 {
				safeCount++
			}
		} else {
			errorCount := 0
			for index, number := range line {
				if index > 0 {
					diff := number - line[index-1]
					if diff < 0 {
						// different directions
						errorCount++
					}
					if diff > 3 || diff == 0 {
						errorCount++
					}
				}
			}
			if errorCount <= 1 {
				safeCount++
			}
		}
	}

	return strconv.Itoa(safeCount)
}
