package main

import (
	"flag"
	"fmt"
	"time"
)

const AllDays = -1

type Day struct {
	day          uint8
	part1, part2 func() any
}

var dayarr = []Day{
	{
		day:   1,
		part1: func() any { return day1Part1(readLines(1)) },
		part2: func() any { return day1Part2(readLines(1)) },
	},
	{
		day:   2,
		part1: func() any { return day2Part1(readLines(2)) },
		part2: func() any { return day2Part2(readLines(2)) },
	},
	{
		day:   3,
		part1: func() any { return day3Part1(readLines(3)) },
		part2: func() any { return day3Part2(readLines(3)) },
	},
	{
		day:   4,
		part1: func() any { return day4Part1(readLines(4)) },
		part2: func() any { return day4Part2(readLines(4)) },
	},
	{
		day:   5,
		part1: func() any { return day5Part1(readLines(5)) },
		part2: func() any { return day5Part2(readLines(5)) },
	},
	{
		day:   6,
		part1: func() any { return day6Part1(readLines(6)) },
		part2: func() any { return day6Part2(readLines(6)) },
	},
	{
		day:   7,
		part1: func() any { return day7Part1(readLines(7)) },
		part2: func() any { return day7Part2(readLines(7)) },
	},
	{
		day:   8,
		part1: func() any { return day8Part1(readLines(8)) },
		part2: func() any { return day8Part2(readLines(8)) },
	},
	{
		day:   9,
		part1: func() any { return day9Part1(readLines(9)) },
		part2: func() any { return day9Part2(readLines(9)) },
	},
	{
		day:   10,
		part1: func() any { return day10Part1(readLines(10)) },
		part2: func() any { return day10Part2(readLines(10)) },
	},
	{
		day:   11,
		part1: func() any { return day11Part1(readLines(11)) },
		part2: func() any { return day11Part2(readLines(11)) },
	},
	{
		day:   12,
		part1: func() any { return day12Part1(readLines(12)) },
		part2: func() any { return day12Part2(readLines(12)) },
	},
	{
		day:   13,
		part1: func() any { return day13Part1(readLines(13)) },
		part2: func() any { return day13Part2(readLines(13)) },
	},
	{
		day:   14,
		part1: func() any { return day14Part1(readLines(14)) },
		part2: func() any { return day14Part2(readLines(14)) },
	},
	{
		day:   15,
		part1: func() any { return day15Part1(readLines(15)) },
		part2: func() any { return day15Part2(readLines(15)) },
	},
	{
		day:   16,
		part1: func() any { return day16Part1(readLines(16)) },
		part2: func() any { return day16Part2(readLines(16)) },
	},
	{
		day:   17,
		part1: func() any { return day17Part1(readLines(17)) },
		part2: func() any { return day17Part2(readLines(17)) },
	},
	{
		day:   18,
		part1: func() any { return day18Part1(readLines(18)) },
		part2: func() any { return day18Part2(readLines(18)) },
	},
	{
		day:   19,
		part1: func() any { return day19Part1(readLines(19)) },
		part2: func() any { return day19Part2(readLines(19)) },
	},
}

func runDay(day Day) {
	assert(day.day <= 25)
	start := time.Now()
	fmt.Println(fmt.Sprintf("day%d part1", day.day), day.part1(), "-", "time:", time.Since(start))
	start = time.Now()
	fmt.Println(fmt.Sprintf("day%d part2", day.day), day.part2(), "-", "time:", time.Since(start))
}

func main() {
	day := flag.Int("day", AllDays, "Run speciphic day")
	flag.Parse()

	if *day == AllDays {
		for _, day := range dayarr {
			runDay(day)
		}
	} else {
		assert(*day > 0)
		runDay(dayarr[*day-1])
	}
}
