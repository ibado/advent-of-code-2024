package main

import (
	"iter"
)

type day9 struct{}

func (d day9) Part1(lines iter.Seq[string]) any {
	next, stop := iter.Pull(lines)
	defer stop()
	l, _ := next()
	var blocks []int
	for i, r := range l {
		val := int(r - '0')
		if i%2 == 0 {
			for j := 0; j < val; j++ {
				blocks = append(blocks, i/2)
			}
		} else {
			for j := 0; j < val; j++ {
				blocks = append(blocks, -1)
			}
		}
	}
	startptr := 0
	endptr := len(blocks) - 1
	res := make([]int, len(blocks))
	for startptr != endptr {
		if blocks[startptr] != -1 {
			res[startptr] = blocks[startptr]
			startptr++
		} else if blocks[endptr] != -1 {
			res[startptr] = blocks[endptr]
			endptr--
			startptr++

		} else {
			endptr--
		}
	}
	res[startptr] = blocks[endptr]

	var sum int = 0
	for i := 0; i < len(res) && res[i] != -1; i++ {
		sum += i * res[i]
	}

	return sum
}

func (d day9) Part2(lines iter.Seq[string]) any {
	next, stop := iter.Pull(lines)
	defer stop()
	l, _ := next()
	var blocks []int
	for i, r := range l {
		val := int(r - '0')
		if i%2 == 0 {
			for j := 0; j < val; j++ {
				blocks = append(blocks, i/2)
			}
		} else {
			for j := 0; j < val; j++ {
				blocks = append(blocks, -1)
			}
		}
	}

	endptr := len(blocks) - 1
	for endptr >= 0 {
		if blocks[endptr] != -1 {
			size := 0
			val := blocks[endptr]
			for endptr >= 0 && val == blocks[endptr] {
				size++
				endptr--
			}
			if size == 0 {
				continue
			}

			freeSize := 0
			i := 0
			for ; i < endptr; i++ {
				if blocks[i] == -1 {
					freeSize++
					if freeSize == size {
						break
					}
				} else {
					freeSize = 0
				}
			}
			if freeSize == size {
				for j := endptr + 1; j < endptr+1+size; j++ {
					blocks[j] = -1
				}
				for j := i; j > i-freeSize; j-- {
					blocks[j] = val
				}
			}
		} else {
			endptr--
		}
	}

	var sum int = 0
	for i := 0; i < len(blocks); i++ {
		if blocks[i] == -1 {
			continue
		}
		sum += i * blocks[i]
	}
	return sum
}

// 6420913947251
// 6420913943576
