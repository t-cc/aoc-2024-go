package main

import (
	"flag"
	"github.com/wlchs/advent_of_code_go_template/days/day_0"
)

// main entry point
// A --day parameter is required to choose which daily challenge should be executed.
func main() {
	day := flag.String("day", "", "day ID to execute")

	switch *day {
	case "0":
		day_0.Run()
	}
}
