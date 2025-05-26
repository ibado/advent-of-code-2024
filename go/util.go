package main

import (
	"bufio"
	"fmt"
	"iter"
	"os"
)

type Queue[T any] []T

func (q *Queue[T]) Push(elem T) {
	*q = append(*q, elem)
}

func (q *Queue[T]) PushSlice(s []T) {
	*q = append(*q, s...)
}

func (q *Queue[T]) Pop() T {
	if len(*q) == 0 {
		panic("empty queue")
	}
	popped := (*q)[0]
	*q = (*q)[1:]
	return popped
}

func (q *Queue[T]) Len() int {
	return len(*q)
}

type Point struct {
	x, y int
}

func (p Point) String() string {
	return fmt.Sprint(p.x, ",", p.y)
}

func (p Point) isInRage(n int) bool {
	return p.x >= 0 && p.x < n && p.y >= 0 && p.y < n
}

func (p Point) Up() Point {
	return Point{p.x - 1, p.y}
}

func (p Point) Down() Point {
	return Point{p.x + 1, p.y}
}

func (p Point) Right() Point {
	return Point{p.x, p.y + 1}
}

func (p Point) Left() Point {
	return Point{p.x, p.y - 1}
}

var debug = true

func log(a ...any) {
	if debug {
		fmt.Println(a...)
	}
}

func flatten[T any](s [][]T) []T {
	var f []T
	for _, l := range s {
		for _, e := range l {
			f = append(f, e)
		}
	}
	return f
}

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

func abs[T int64 | int](a T) T {
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

func isAlpha(n byte) bool {
	return 'a' <= n && n <= 'z' || 'A' <= n && n <= 'Z'
}

func concat(n int, c byte) int {
	if n == 0 {
		return int(c - '0')
	} else {
		return n*10 + int(c-'0')
	}
}

func parseMatrix(lines iter.Seq[string]) [][]byte {
	var mx [][]byte
	for l := range lines {
		mx = append(mx, []byte(l))
	}
	return mx
}

func findPoints[T int | byte](mx [][]T, val T) []Point {
	var points []Point
	for i := 0; i < len(mx); i++ {
		for j := 0; j < len(mx[0]); j++ {
			if mx[i][j] == val {
				points = append(points, Point{i, j})
			}
		}
	}
	return points
}

func parseNums(line []byte) []int {
	var s []int
	i := 0
	for i < len(line) {
		isNeg := line[i] == '-'
		if isDigit(line[i]) || isNeg {
			var n int = 0
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
