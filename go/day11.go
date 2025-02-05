package main

import (
	"iter"
	"strconv"
)

type StoneBlinks struct {
	stoneNum, blinks int
}

type day11 struct{}

func (d day11) Part1(lines iter.Seq[string]) any {
	return solveDay11(lines, 25)
}

func (d day11) Part2(lines iter.Seq[string]) any {
	return solveDay11(lines, 75)
}

func solveDay11(lines iter.Seq[string], blinks int) int {
	next, stop := iter.Pull(lines)
	defer stop()
	l, _ := next()
	nums := parseNums([]byte(l))
	sum := 0
	for _, n := range nums {
		sum += computeStones(n, blinks)
	}
	return sum
}

func computeStones(stoneNum, blinks int) int {
	cache := make(map[StoneBlinks]int)
	var computeRec func(int, int) int

	computeRec = func(stoneNum, blinks int) int {
		if blinks == 0 {
			return 1
		}
		sb := StoneBlinks{stoneNum, blinks}
		if val, ok := cache[sb]; ok {
			return val
		}
		var res int
		if stoneNum == 0 {
			res = computeRec(1, blinks-1)
		} else if nstr := strconv.Itoa(stoneNum); len(nstr)%2 == 0 {
			fh, _ := strconv.Atoi(nstr[:len(nstr)/2])
			sh, _ := strconv.Atoi(nstr[len(nstr)/2:])
			res = computeRec(fh, blinks-1) + computeRec(sh, blinks-1)
		} else {
			res = computeRec(stoneNum*2024, blinks-1)
		}
		cache[sb] = res
		return res

	}
	return computeRec(stoneNum, blinks)
}
