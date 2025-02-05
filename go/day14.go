package main

import (
	"iter"
)

const N = 101
const M = 103

type day14 struct{}

func (d day14) Part1(lines iter.Seq[string]) any {
	m := make(map[Point]int)
	for l := range lines {
		nums := parseNums([]byte(l))
		p := Point{nums[0], nums[1]}
		v := Point{nums[2], nums[3]}
		simulate(p, v, m, 100)
	}
	var topLeft, topRight, bottomLeft, bottomRight int
	for k, v := range m {
		if k.x < N/2 && k.y < M/2 {
			topLeft += v
		} else if k.x > N/2 && k.y < M/2 {
			topRight += v
		} else if k.x > N/2 && k.y > M/2 {
			bottomLeft += v
		} else if k.x < N/2 && k.y > M/2 {
			bottomRight += v
		}
	}
	return topLeft * topRight * bottomLeft * bottomRight
}

func (d day14) Part2(lines iter.Seq[string]) any {
	type Robot struct {
		p, v Point
	}
	var robots []Robot
	for l := range lines {
		nums := parseNums([]byte(l))
		p := Point{nums[0], nums[1]}
		v := Point{nums[2], nums[3]}
		robots = append(robots, Robot{p, v})
	}
	for i := 0; ; i++ {
		m := make(map[Point]int)
		for _, r := range robots {
			simulate(r.p, r.v, m, i)
		}
		if len(m) == len(robots) { // none of the robots are overlapping to each other
			return i
		}
	}
}

func simulate(p, v Point, m map[Point]int, iters int) {
	p.x = mod(p.x+iters*v.x, N)
	p.y = mod(p.y+iters*v.y, M)
	m[p]++
}

func mod(n, m int) int {
	return ((n)%m + m) % m
}
