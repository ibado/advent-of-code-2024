package main

import (
	"iter"
	"maps"
	"strings"
)

func day19Part1(lines iter.Seq[string]) int {
	towels, patterns := inputDay19(lines)
	count := 0
	for _, pattern := range patterns {
		if isPossible(pattern, towels) {
			count++
		}
	}
	return count
}

func day19Part2(lines iter.Seq[string]) int {
	towels, patterns := inputDay19(lines)
	count := 0
	for _, pattern := range patterns {
		c := countPossible(pattern, towels)
		count += c
	}
	return count

}

func countPossible(comb string, towels map[string]bool) int {
	var rec func(string, map[string]bool) int
	m := make(map[string]int)
	rec = func(comb string, towels map[string]bool) int {
		if c, ok := m[comb]; ok {
			return c
		}
		if comb == "" {
			return 1
		}
		res := 0
		for t := range maps.Keys(towels) {
			if !strings.HasPrefix(comb, t) {
				continue
			}
			r := rec(comb[len(t):], towels)
			m[comb[len(t):]] = r
			res += r

		}

		return res
	}
	return rec(comb, towels)
}

func isPossible(comb string, towels map[string]bool) bool {
	var rec func(string, map[string]bool) bool
	m := make(map[string]bool)
	rec = func(comb string, towels map[string]bool) bool {
		if v, ok := m[comb]; ok {
			return v
		}
		if _, ok := towels[comb]; ok {
			return true
		}

		for t := range maps.Keys(towels) {
			if strings.HasPrefix(comb, t) && rec(comb[len(t):], towels) {
				m[comb] = true
				return true
			}
		}

		m[comb] = false
		return false

	}
	return rec(comb, towels)
}

func inputDay19(lines iter.Seq[string]) (towels map[string]bool, patterns []string) {
	towels = make(map[string]bool)
	next, stop := iter.Pull(lines)
	defer stop()

	firstLine, _ := next()
	for _, t := range strings.Split(firstLine, ", ") {
		towels[t] = true
	}

	for {
		pattern, ok := next()
		if !ok {
			break
		}
		patterns = append(patterns, pattern)
	}

	return towels, patterns
}
