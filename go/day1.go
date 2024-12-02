package main

import (
	"iter"
	"slices"
)

func day1Part1(lines iter.Seq[string]) int64 {
	left, right := parseInput(lines)

	slices.Sort(left)
	slices.Sort(right)
	var res int64 = 0
	for i := 0; i < len(left); i++ {
		res += abs(left[i] - right[i])
	}

	return res
}

func day1Part2(lines iter.Seq[string]) int64 {
	left, right := parseInput(lines)

	m := make(map[int64]int64, len(left))
	for _, nr := range right {
		m[nr] += 1
	}
	var res int64 = 0
	for _, nl := range left {
		res += nl * m[nl]
	}

	return res
}

func parseInput(input iter.Seq[string]) ([]int64, []int64) {
	var left, right []int64
	for l := range input {
		nums := parseNums([]byte(l))
		left = append(left, nums[0])
		right = append(right, nums[1])
	}
	return left, right
}
