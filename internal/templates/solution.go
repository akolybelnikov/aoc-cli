package templates

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("inputs/day{{DAY}}.txt")
	if err != nil {
		fmt.Println("Failed to read input file.", err)
		os.Exit(0)
	}
	input := string(data)

	fmt.Println("--- Part One ---")
	fmt.Println("Result:", part1(input))

	fmt.Println("--- Part Two ---")
	fmt.Println("Result:", part2(input))

	os.Exit(0)
}

// part one
func part1(input string) int {
	fmt.Println(input)
	return 0
}

// part two
func part2(input string) int {
	fmt.Println(input)
	return 0
}
