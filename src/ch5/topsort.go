package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},
	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func topSort(prereqs map[string][]string) []string {
	var orders []string
	orders = make([]string, 0)
	seen := make(map[string]bool)
	var visitAll func(iterms []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(prereqs[item])
				orders = append(orders, item)
			}
		}
	}

	keys := make([]string, 0)
	for key := range prereqs {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return orders
}

func main() {
	for i, course := range topSort(prereqs) {
		fmt.Printf("%d:%s\n", i, course)
	}
}
