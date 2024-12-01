package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
)

func readLines(day uint8) iter.Seq[string] {
	assert(day <= 25)
	return func(yield func(string) bool) {
		fpath := fmt.Sprintf("../input/%d.txt", day)
		f, _ := os.OpenFile(fpath, os.O_RDONLY, 0)
		defer f.Close()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}

func assert(cond bool) {
	if !cond {
		panic("assertion fail!")
	}
}

func abs(a int64) int64 {
	if a >= 0 {
		return a
	}
	return -a
}

func min(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func isDigit(n byte) bool {
	return '0' <= n && n <= '9'
}

func concat(n int64, c byte) int64 {
	if n == 0 {
		return int64(c - '0')
	} else {
		return n*10 + int64(c-'0')
	}
}

func parseNums(line []byte) []int64 {
	var s []int64
	i := 0
	for i < len(line) {
		isNeg := line[i] == '-'
		if isDigit(line[i]) || isNeg {
			var n int64 = 0
			if isNeg {
				i++
			}
			for i < len(line) && isDigit(line[i]) {
				n = concat(n, line[i])
				i++
			}
			if isNeg {
				n *= -1
			}
			s = append(s, n)
		} else {
			i++
		}
	}
	return s
}
