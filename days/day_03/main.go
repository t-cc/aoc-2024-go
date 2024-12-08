package day_03

import (
	"fmt"
	"regexp"
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
	var sum = 0
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	for _, line := range input {
		for _, m := range re.FindAllString(line, -1) {
			numbers := strings.Split(m[:len(m)-1][4:], ",")
			a, err := strconv.Atoi(numbers[0])
			if err != nil {
				panic(err)
			}
			b, err := strconv.Atoi(numbers[1])
			if err != nil {
				panic(err)
			}
			sum += a * b
		}
	}

	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	sum := 0
	enabled := true
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
	for _, line := range input {
		for _, m := range re.FindAllString(line, -1) {
			if m == "don't()" {
				enabled = false
			} else if m == "do()" {
				enabled = true
			} else if enabled {
				numbers := strings.Split(m[:len(m)-1][4:], ",")
				a, err := strconv.Atoi(numbers[0])
				if err != nil {
					panic(err)
				}
				b, err := strconv.Atoi(numbers[1])
				if err != nil {
					panic(err)
				}
				sum += a * b
			}
		}
	}

	return strconv.Itoa(sum)
}
