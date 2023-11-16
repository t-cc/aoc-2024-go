package day_01_test

import (
	"github.com/wlchs/advent_of_code_go_template/days/day_0"
	"github.com/wlchs/advent_of_code_go_template/loader"
	"testing"
)

func TestPartOne(t *testing.T) {
	t.Parallel()

	expectedResult := loader.LoadFirstInputLine("solution_1.txt")
	result := day_0.Part1("input_1_test.txt")

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}

func TestPartTwo(t *testing.T) {
	t.Parallel()

	expectedResult := loader.LoadFirstInputLine("solution_2.txt")
	result := day_0.Part2("input_2_test.txt")

	if result != expectedResult {
		t.Errorf("expected result was %s, but got %s instead", expectedResult, result)
	}
}
