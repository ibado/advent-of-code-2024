package main

import (
	"iter"
	"slices"
)

type Rule struct {
	first, second int64
}

func day5Part1(lines iter.Seq[string]) int64 {
	rules, pages := parseInputDay5(lines)

	var sum int64 = 0
	for _, page := range pages {
		valid := true
		for _, rule := range rules {
			if slices.Contains(page, rule.first) && slices.Contains(page, rule.second) {
				fidx := slices.Index(page, rule.first)
				sidx := slices.Index(page, rule.second)
				if fidx > sidx {
					valid = false
					break
				}
			}
		}
		if valid {
			sum += page[len(page)/2]
		}
	}
	return sum
}

func day5Part2(lines iter.Seq[string]) int64 {
	rules, pages := parseInputDay5(lines)

	var sum int64 = 0
	for _, page := range pages {
		for _, rule := range rules {
			if slices.Contains(page, rule.first) && slices.Contains(page, rule.second) {
				fidx := slices.Index(page, rule.first)
				sidx := slices.Index(page, rule.second)
				if fidx > sidx {
					slices.SortFunc(page, func(a, b int64) int {
						idx := slices.IndexFunc(rules, func(r Rule) bool {
							return r.first == a && r.second == b || r.first == b && r.second == a
						})
						if idx != -1 {
							rule := rules[idx]
							if rule.first == a {
								return -1
							} else {
								return 1
							}

						}
						return 0
					})
					sum += page[len(page)/2]
					break
				}
			}
		}
	}
	return sum
}

func parseInputDay5(lines iter.Seq[string]) ([]Rule, [][]int64) {
	var rules []Rule
	var pages [][]int64
	rulesEnd := false
	for l := range lines {
		if l == "" {
			rulesEnd = true
			continue
		}
		if !rulesEnd {
			nums := parseNums([]byte(l))
			rules = append(rules, Rule{nums[0], nums[1]})
		} else {
			pages = append(pages, parseNums([]byte(l)))
		}
	}
	return rules, pages
}
