package main

import (
	"fmt"
	"os"
	"strings"
)

type group struct {
	connections [3]string
}

func (g group) sort() group {
	sorted := g.connections[:]
	for i := 0; i < 2; i++ {
		for j := i + 1; j < 3; j++ {
			if sorted[i] > sorted[j] {
				sorted[i], sorted[j] = sorted[j], sorted[i]
			}
		}
	}
	return group{[3]string{sorted[0], sorted[1], sorted[2]}}
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day23/input.txt")
	input := string(bytes)
	split := strings.Split(input, "\n")
	fmt.Println(part1(split))
}

func part1(input []string) int {
	connections := make(map[string][]string)
	for _, line := range input {
		split := strings.Split(line, "-")
		connections[split[0]] = append(connections[split[0]], split[1])
		connections[split[1]] = append(connections[split[1]], split[0])
	}

	groups := make(map[group]struct{})
	for key, value := range connections {
		if key[0] != 't' {
			continue
		}
		for _, candidate := range getPairs(value) {
			if contains(connections[candidate.a], candidate.b) {
				g := group{[3]string{key, candidate.a, candidate.b}}.sort()
				groups[g] = struct{}{}
			}
		}
	}
	return len(groups)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type pair struct {
	a, b string
}

func getPairs(s []string) []pair {
	var pairs []pair
	for i := 0; i < len(s); i++ {
		for j := i + 1; j < len(s); j++ {
			pairs = append(pairs, pair{s[i], s[j]})
		}
	}
	return pairs
}
