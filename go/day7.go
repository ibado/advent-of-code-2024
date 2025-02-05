package main

import (
	"iter"
	"slices"
	"strconv"
	"sync"
)

var operators1 = []byte{'+', '*'}
var operators2 = []byte{'+', '*', 'c'}

type Day7Input struct {
	nums   []int
	target int
}

type day7 struct{}

func (d day7) Part1(lines iter.Seq[string]) any {
	return solve(operators1, lines)
}

func (d day7) Part2(lines iter.Seq[string]) any {
	return solve(operators2, lines)
}

func solve(ops []byte, lines iter.Seq[string]) int {
	m := make(map[int][][]byte)
	input := parseDay7(lines)
	for _, in := range input {
		nums := in.nums
		var opsPerm [][]byte
		if val, ok := m[len(nums)]; ok {
			opsPerm = val
		} else {
			opsPerm = genPermutations(ops, len(nums))
			m[len(nums)] = opsPerm
		}
	}

	sumCh := make(chan int)
	var wg sync.WaitGroup
	for _, in := range input {
		wg.Add(1)
		go func() {
			target := in.target
			nums := in.nums
			opsPerm := m[len(nums)]
		LOOP:
			for _, ops := range opsPerm {
				result := nums[0]
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
					sumCh <- target
					break LOOP
				}
			}

			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(sumCh)
	}()

	sum := 0
	for s := range sumCh {
		sum += s
	}
	return sum
}

func concatOp(a, b int) int {
	n := a
	for _, c := range []byte(strconv.Itoa(b)) {
		n = concat(n, c)
	}
	return n
}

func genPermutations(ops []byte, n int) [][]byte {
	var rec func(ops []byte, prefix []byte, res *[][]byte, n int)

	rec = func(ops []byte, prefix []byte, res *[][]byte, n int) {
		if n == 1 {
			*res = append(*res, slices.Clone(prefix))
			return
		}

		for _, op := range ops {
			np := append(prefix, op)
			rec(ops, np, res, n-1)
		}
	}

	var res [][]byte
	rec(ops, []byte{}, &res, n)
	return res
}

func parseDay7(lines iter.Seq[string]) []Day7Input {
	var res []Day7Input
	for l := range lines {
		n := parseNums([]byte(l))
		target := n[0]
		nums := n[1:]
		res = append(res, Day7Input{nums, target})
	}
	return res
}
