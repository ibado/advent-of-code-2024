package main

import (
	"iter"
)

var WORD1 = []byte{'X', 'M', 'A', 'S'}
var WORD2 = []byte{'M', 'A', 'S'}

func day4Part1(lines iter.Seq[string]) int {
	mx := parseMatrix(lines)
	count := 0
	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx[0]); j++ {
			count += searchPart1(mx, i, j)
		}
	}
	return count
}

func day4Part2(lines iter.Seq[string]) int {
	mx := parseMatrix(lines)
	count := 0
	for i := 0; i < len(mx)-2; i++ {
		for j := 0; j < len(mx[0])-2; j++ {
			c := 0
			c += searchij(mx, WORD2, i, j, incr, incr)
			c += searchij(mx, WORD2, i, j+2, incr, decr)
			c += searchij(mx, WORD2, i+2, j, decr, incr)
			c += searchij(mx, WORD2, i+2, j+2, decr, decr)
			if c == 2 {
				count++
			}
		}
	}
	return count
}

func searchPart1(mx [][]byte, i int, j int) int {
	count := 0
	// search down
	count += searchij(mx, WORD1, i, j, incr, idle)
	// search right
	count += searchij(mx, WORD1, i, j, idle, incr)
	// search up
	count += searchij(mx, WORD1, i, j, decr, idle)
	// search left
	count += searchij(mx, WORD1, i, j, idle, decr)
	// search diagonally right down
	count += searchij(mx, WORD1, i, j, incr, incr)
	// search diagonally left up
	count += searchij(mx, WORD1, i, j, decr, decr)
	// search diagonally right up
	count += searchij(mx, WORD1, i, j, decr, incr)
	// search diagonally left down
	count += searchij(mx, WORD1, i, j, incr, decr)
	return count
}

func searchij(
	mx [][]byte,
	word []byte,
	i int, j int,
	ifunc func(int, int) int,
	jfunc func(int, int) int,
) int {
	wlen := len(word)
	mxlen := len(mx)
	for k := 0; k < wlen; k++ {
		ii := ifunc(i, k)
		jj := jfunc(j, k)
		if ii < 0 || jj < 0 || ii >= mxlen || jj >= mxlen || mx[ii][jj] != word[k] {
			return 0
		}
	}
	return 1
}

func incr(i, inc int) int {
	return i + inc
}

func decr(i, dec int) int {
	return i - dec
}

func idle(i, inc int) int {
	return i
}
