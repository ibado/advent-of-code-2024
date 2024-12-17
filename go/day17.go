package main

import (
	"iter"
	"maps"
	"math"
	"slices"
	"strconv"
	"strings"
)

func day17Part1(lines iter.Seq[string]) string {
	next, stop := iter.Pull(lines)
	defer stop()
	regAInput, _ := next()
	next() // regB
	next() // regC
	next() // empty line
	opcodesInput, _ := next()
	regA := parseNums([]byte(regAInput))[0]
	opcodes := parseNums([]byte(opcodesInput))

	output := runProgram(regA, opcodes)
	var s strings.Builder

	for i, o := range output {
		s.WriteString(strconv.Itoa(o))
		if i != len(output)-1 {
			s.WriteString(",")
		}
	}

	return s.String()
}

func day17Part2(lines iter.Seq[string]) int {
	next, stop := iter.Pull(lines)
	defer stop()
	// ignore the first 4 lines
	next()
	next()
	next()
	next()
	opcodesInput, _ := next()
	opcodes := parseNums([]byte(opcodesInput))

	as := make(map[int]bool)
	as[0] = true
	for _, output := range slices.Backward(opcodes) {
		for nextA, _ := range maps.Clone(as) {
			delete(as, nextA)
			for i := 0; i < 8; i++ {
				aa := nextA*8 + i
				if runProgram(aa, opcodes)[0] == output {
					as[aa] = true
				}
			}
		}
	}

	return slices.Min(slices.Collect(maps.Keys(as)))
}

func runProgram(a int, instructions []int) []int {
	regA := a
	regB := 0
	regC := 0
	opcodes := instructions

	combo := func(op int) int {
		var res int
		switch op {
		case 0, 1, 2, 3:
			res = op
		case 4:
			res = regA
		case 5:
			res = regB
		case 6:
			res = regC
		}
		return res
	}

	var output []int
	insPtr := 0
	for insPtr < len(opcodes)-1 {
		opcode := opcodes[insPtr]
		operand := opcodes[insPtr+1]
		switch opcode {
		case 0:
			regA = int(float64(regA) / math.Pow(2, float64(combo(operand))))
		case 1:
			regB = regB ^ operand
		case 2:
			regB = combo(operand) % 8
		case 3:
			if regA != 0 {
				insPtr = operand
				continue
			}
		case 4:
			regB = regB ^ regC
		case 5:
			out := combo(operand) % 8
			output = append(output, out)
		case 6:
			regB = int(float64(regA) / math.Pow(2, float64(combo(operand))))
		case 7:
			regC = int(float64(regA) / math.Pow(2, float64(combo(operand))))
		}
		insPtr += 2
	}

	return output
}
