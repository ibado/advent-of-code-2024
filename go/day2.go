package main

import (
	"iter"
	"slices"
)

type day2 struct{}

func (d day2) Part1(lines iter.Seq[string]) any {
	count := 0
	for line := range lines {
		levels := parseNums([]byte(line))
		if isSafe(levels) {
			count++
		}
	}
	return count
}

func (d day2) Part2(lines iter.Seq[string]) any {
	count := 0
	for line := range lines {
		levels := parseNums([]byte(line))
		for i := 0; i < len(levels); i++ {
			l := slices.Delete(slices.Clone(levels), i, i+1)
			if isSafe(l) {
				count++
				break
			}
		}
	}
	return count

}

func isUnsafe(l int, asc bool) bool {
	if asc {
		if l >= 0 || l < -3 {
			return true
		}
	} else {
		if l <= 0 || l > 3 {
			return true
		}
	}
	return false
}

func isSafe(levels []int) bool {
	diff := levels[0] - levels[1]
	if diff == 0 || abs(diff) > 3 {
		return false
	}
	asc := diff < 0
	for i := 1; i < len(levels)-1; i++ {
		l := levels[i] - levels[i+1]
		if isUnsafe(l, asc) {
			return false
		}
	}
	return true
}
