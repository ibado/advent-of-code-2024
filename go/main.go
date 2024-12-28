package main

import (
	"fmt"
	"time"
)

func run(msg string, result func() any) {
	start := time.Now()
	fmt.Println(msg, result(), "-", "time:", time.Since(start))
}

func main() {
	run("day1 part1:", func() any { return day1Part1(readLines(1)) })
	run("day1 part2:", func() any { return day1Part2(readLines(1)) })

	run("day2 part1:", func() any { return day2Part1(readLines(2)) })
	run("day2 part2:", func() any { return day2Part2(readLines(2)) })

	run("day3 part1:", func() any { return day3Part1(readLines(3)) })
	run("day3 part2:", func() any { return day3Part2(readLines(3)) })

	run("day4 part1:", func() any { return day4Part1(readLines(4)) })
	run("day4 part2:", func() any { return day4Part2(readLines(4)) })

	run("day5 part1:", func() any { return day5Part1(readLines(5)) })
	run("day5 part2:", func() any { return day5Part2(readLines(5)) })

	run("day6 part1:", func() any { return day6Part1(readLines(6)) })
	run("day6 part2:", func() any { return day6Part2(readLines(6)) })

	run("day7 part1:", func() any { return day7Part1(readLines(7)) })
	run("day7 part2:", func() any { return day7Part2(readLines(7)) })

	run("day8 part1:", func() any { return day8Part1(readLines(8)) })
	run("day8 part2:", func() any { return day8Part2(readLines(8)) })

	run("day9 part1:", func() any { return day9Part1(readLines(9)) })
	run("day9 part2:", func() any { return day9Part2(readLines(9)) })

	run("day10 part1:", func() any { return day10Part1(readLines(10)) })
	run("day10 part2:", func() any { return day10Part2(readLines(10)) })

	run("day11 part1:", func() any { return day11Part1(readLines(11)) })
	run("day11 part2:", func() any { return day11Part2(readLines(11)) })

	run("day12 part1:", func() any { return day12Part1(readLines(12)) })
	run("day12 part2:", func() any { return day12Part2(readLines(12)) })

	run("day13 part1:", func() any { return day13Part1(readLines(13)) })
	run("day13 part2:", func() any { return day13Part2(readLines(13)) })

	run("day14 part1:", func() any { return day14Part1(readLines(14)) })
	run("day14 part2:", func() any { return day14Part2(readLines(14)) })

	run("day15 part1:", func() any { return day15Part1(readLines(15)) })
	run("day15 part2:", func() any { return day15Part2(readLines(15)) })

	run("day16 part1:", func() any { return day16Part1(readLines(16)) })
	run("day16 part2:", func() any { return day16Part2(readLines(16)) })

	run("day17 part1:", func() any { return day17Part1(readLines(17)) })
	run("day17 part2:", func() any { return day17Part2(readLines(17)) })

	run("day18 part1:", func() any { return day18Part1(readLines(18)) })
	run("day18 part2:", func() any { return day18Part2(readLines(18)) })

	run("day19 part1:", func() any { return day19Part1(readLines(19)) })
	run("day19 part2:", func() any { return day19Part2(readLines(19)) })
}
