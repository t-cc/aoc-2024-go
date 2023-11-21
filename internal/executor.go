package internal

import (
	"github.com/wlchs/advent_of_code_go_template/days/day_0"
	"github.com/wlchs/advent_of_code_go_template/days/day_01"
)

// RunChallenge executes the challenge of a specific day with the provided input.
func RunChallenge(day int, inputPath string, mode int) {
	input := LoadInputLines(inputPath)
	mapping := map[int]func([]string, int){
		0: day_0.Run,
		1: day_01.Run,
	}

	mapping[day](input, mode)
}
