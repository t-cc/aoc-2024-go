package day_01

import (
	"fmt"
	"sort"
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
	lista1 := []int{}
	lista2 := []int{}
	for _, line := range input {
		numbers := strings.Split(line, "   ")
		number1, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		lista1 = append(lista1, number1)

		number2, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		lista2 = append(lista2, number2)
	}
	sort.Ints(lista1)
	sort.Ints(lista2)

	var sum = 0
	for index, number := range lista1 {
		if number-lista2[index] < 0 {
			sum += lista2[index] - number
		} else {
			sum += number - lista2[index]
		}
	}

	return strconv.Itoa(sum)
}

// Part2 solves the second part of the exercise
func Part2(input []string) string {
	lista1 := []int{}
	lista2 := []int{}
	for _, line := range input {
		numbers := strings.Split(line, "   ")
		number1, err := strconv.Atoi(numbers[0])
		if err != nil {
			panic(err)
		}
		lista1 = append(lista1, number1)

		number2, err := strconv.Atoi(numbers[1])
		if err != nil {
			panic(err)
		}
		lista2 = append(lista2, number2)
	}

	var sum = 0
	for _, number := range lista1 {
		similarity := 0
		for _, number2 := range lista2 {
			if number2 == number {
				similarity++
			}
		}
		if similarity > 0 {
			sum += number * similarity
		}
	}

	return strconv.Itoa(sum)
}
