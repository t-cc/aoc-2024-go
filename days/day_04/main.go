package day_04

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

func reverse(numbers []string) []string {
	newNumbers := make([]string, 0, len(numbers))
	for i := len(numbers) - 1; i >= 0; i-- {
		newNumbers = append(newNumbers, numbers[i])
	}
	return newNumbers
}

func buildMatrix(input []string) [][]string {
	var matrix [][]string
	for _, line := range input {
		lineArray := strings.Split(line, "")
		matrix = append(matrix, lineArray)
	}
	return matrix
}

func searchForMAS(matrix [][]string, x int, y int) int {
	return 0

}

var DIRECTIONS = [8][3][2]int{
	{{00, 01}, {00, 02}, {00, 03}},
	{{00, -1}, {00, -2}, {00, -3}},
	{{01, 00}, {02, 00}, {03, 00}},
	{{-1, 00}, {-2, 00}, {-3, 00}},
	{{01, 01}, {02, 02}, {03, 03}},
	{{-1, -1}, {-2, -2}, {-3, -3}},
	{{-1, 01}, {-2, 02}, {-3, 03}},
	{{01, -1}, {02, -2}, {03, -3}},
}

// Part1 solves the first part of the exercise
func Part1(input []string) string {
	total := 0

	matrix := buildMatrix(input)

	xMax := len(input[0])
	yMax := len(input)
	for x := 0; x < xMax; x++ {
		for y := 0; y < yMax; y++ {
			start := matrix[y][x]
			if start == "X" {
				// test for MAS
				for _, row := range DIRECTIONS {
					list := []string{}
					for _, point := range row {
						_y := point[0] + y
						_x := point[1] + x
						if _x >= 0 && _x < xMax && _y >= 0 && _y < yMax {
							list = append(list, matrix[_y][_x])
						}
					}
					if len(list) == 3 && list[0] == "M" && list[1] == "A" && list[2] == "S" {
						total++
					}
				}
			}
		}
	}

	return strconv.Itoa(total)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	return ""
}
