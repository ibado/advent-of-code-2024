package main

import (
	"iter"
	"math"
	"strings"
)

type Game struct {
	A, B, Prize Point
}

const maxBtnPresses = 100

func day13Part1(lines iter.Seq[string]) int {
	games := parseGames(lines)
	sum := 0
	for _, g := range games {
		x, y := solveGame(g)
		if x > 0 && y > 0 && x < maxBtnPresses && y < maxBtnPresses && g.A.x*x+g.B.x*y == g.Prize.x && g.A.y*x+g.B.y*y == g.Prize.y {
			sum += x*3 + y
		}
	}
	return sum
}

func day13Part2(lines iter.Seq[string]) int {
	games := parseGames(lines)
	sum := 0
	for _, g := range games {
		g.Prize.x += 10000000000000
		g.Prize.y += 10000000000000
		x, y := solveGame(g)
		if g.A.x*x+g.B.x*y == g.Prize.x && g.A.y*x+g.B.y*y == g.Prize.y {
			sum += x*3 + y
		}
	}
	return sum

}

func solveGame(game Game) (int, int) {
	mx := make([][]float64, 2)
	mx[0] = append(mx[0], float64(game.A.x))
	mx[0] = append(mx[0], float64(game.B.x))
	mx[0] = append(mx[0], float64(game.Prize.x))
	mx[1] = append(mx[1], float64(game.A.y))
	mx[1] = append(mx[1], float64(game.B.y))
	mx[1] = append(mx[1], float64(game.Prize.y))
	GaussianElimination(mx, 2)

	x := mx[0][2] / mx[0][0]
	y := mx[1][2] / mx[1][1]
	return int(math.Round(x)), int(math.Round(y))
}

func Swap(a *float64, b *float64) {
	temp := *a
	*a = *b
	*b = temp
}

func GaussianElimination(a [][]float64, n int) {
	for i := 0; i < n; i++ {
		if a[i][i] == 0.0 {
			c := 1
			for (a[i][i] == 0) && (c < n) {
				Swap(&a[i][i], &a[i+c][i])
				c++
			}
		}
		for j := 0; j < n; j++ {
			if i != j {
				ratio := a[j][i] / a[i][i]
				for k := 0; k <= n; k++ {
					a[j][k] = a[j][k] - ratio*a[i][k]
				}
			}
		}
	}
}

func parseGames(lines iter.Seq[string]) []Game {
	var btnA, btnB, prize Point
	var games []Game
	for l := range lines {
		if l == "" {
			continue
		}
		xy := parseNums([]byte(l))
		p := Point{xy[0], xy[1]}
		if strings.Contains(l, "A:") {
			btnA = p
			continue
		}
		if strings.Contains(l, "B:") {
			btnB = p
			continue
		}
		if strings.Contains(l, "Prize:") {
			prize = p
			games = append(games, Game{btnA, btnB, prize})
		}
	}
	return games
}
