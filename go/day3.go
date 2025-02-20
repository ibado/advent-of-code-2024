package main

import (
	"iter"
	"regexp"
)

type day3 struct{}

func (d day3) Part1(lines iter.Seq[string]) any {
	var mulReg = regexp.MustCompile(`mul\(\d+,\d+\)`)

	res := 0
	for l := range lines {
		for _, mul := range mulReg.FindAll([]byte(l), -1) {
			operands := parseNums(mul)
			res += operands[0] * operands[1]
		}
	}

	return res
}

func (d day3) Part2(lines iter.Seq[string]) any {
	var mulReg = regexp.MustCompile(`mul\(\d+,\d+\)|don't\(\)|do\(\)`)

	res := 0
	do := true
	for l := range lines {
		ops := mulReg.FindAll([]byte(l), -1)
		for _, op := range ops {
			opstr := string(op)
			if opstr == "do()" {
				do = true
			} else if opstr == "don't()" {
				do = false
			} else if do {
				operands := parseNums(op)
				res += operands[0] * operands[1]
			}
		}
	}

	return res
}
