package main

import (
	"iter"
	"slices"
	"strconv"
)

var operators1 = []byte{'+', '*'}
var operators2 = []byte{'+', '*', 'c'}

func day7Part1(lines iter.Seq[string]) int64 {
	return solve(operators1, lines)
}

func day7Part2(lines iter.Seq[string]) int64 {
	return solve(operators2, lines)
}

func solve(ops []byte, lines iter.Seq[string]) int64 {
	var sum int64 = 0
	for l := range lines {
		n := parseNums([]byte(l))
		target := n[0]
		nums := n[1:]
		opsComb := generateOps(ops, len(nums)-1)
	LOOP:
		for _, ops := range opsComb {
			var result int64 = nums[0]
			for i, n := range nums[1:] {
				switch ops[i] {
				case '+':
					result += n
				case '*':
					result *= n
				case 'c':
					result = concatOp(result, n)
				}
			}
			if result == target {
				sum += target
				break LOOP
			}
		}

	}
	return sum
}

func concatOp(a, b int64) int64 {
	var n int64 = a
	for _, c := range []byte(strconv.FormatInt(b, 10)) {
		n = concat(n, c)
	}
	return n
}

func generateOps(ops []byte, n int) [][]byte {
	var res [][]byte
	generateOpsRec(ops, []byte{}, &res, n)
	return res
}

func generateOpsRec(ops []byte, prefix []byte, res *[][]byte, n int) {
	if n == 0 {
		*res = append(*res, slices.Clone(prefix))
		return
	}

	for _, op := range ops {
		np := append(prefix, op)
		generateOpsRec(ops, np, res, n-1)
	}
}
