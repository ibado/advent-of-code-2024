package main

import (
	"iter"
	"slices"
)

type Rule struct {
	first, second int
}

type day5 struct{}

func (d day5) Part1(lines iter.Seq[string]) any {
	rules, pages := parseInputDay5(lines)

	sum := 0
	for _, page := range pages {
		valid := true
		for _, rule := range rules {
			fidx := slices.Index(page, rule.first)
			sidx := slices.Index(page, rule.second)
			if fidx != -1 && sidx != -1 && fidx > sidx {
				valid = false
				break
			}
		}
		if valid {
			sum += page[len(page)/2]
		}
	}
	return sum
}

func (d day5) Part2(lines iter.Seq[string]) any {
	rules, pages := parseInputDay5(lines)

	sum := 0
	for _, page := range pages {
		for _, rule := range rules {
			fidx := slices.Index(page, rule.first)
			sidx := slices.Index(page, rule.second)
			if fidx != -1 && sidx != -1 && fidx > sidx {
				slices.SortFunc(page, func(a, b int) int {
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
	return sum
}

func parseInputDay5(lines iter.Seq[string]) ([]Rule, [][]int) {
	var rules []Rule
	var pages [][]int
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
