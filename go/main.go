package main

import (
	"fmt"
)

func main() {
	fmt.Printf("day1 part1: %d\n", day1Part1(readLines(1)))
	fmt.Printf("day1 part2: %d\n", day1Part2(readLines(1)))

	fmt.Printf("day2 part1: %d\n", day2Part1(readLines(2)))
	fmt.Printf("day2 part2: %d\n", day2Part2(readLines(2)))

	fmt.Printf("day3 part1: %d\n", day3Part1(readLines(3)))
	fmt.Printf("day3 part2: %d\n", day3Part2(readLines(3)))

	fmt.Printf("day4 part1: %d\n", day4Part1(readLines(4)))
	fmt.Printf("day4 part2: %d\n", day4Part2(readLines(4)))

	fmt.Printf("day5 part1: %d\n", day5Part1(readLines(5)))
	fmt.Printf("day5 part2: %d\n", day5Part2(readLines(5)))
}
