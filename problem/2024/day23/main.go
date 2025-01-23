package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

type group struct {
	connections [3]string
}

func (g group) sort() group {
	sorted := g.connections[:]
	sort.Strings(sorted)
	return group{[3]string{sorted[0], sorted[1], sorted[2]}}
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day23/input.txt")
	input := string(bytes)
	split := strings.Split(input, "\n")
	fmt.Println(part1(split))
	fmt.Println(part2(split))
}

func part1(input []string) int {
	connections := parseConnections(input)
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

func part2(input []string) string {
	connections := parseConnections(input)
	pools := make(map[string]struct{})
	for key := range connections {
		pools[key] = struct{}{}
	}
	old := make(map[string]struct{})

	for hash(old) != hash(pools) {
		for current := range connections {
			for pool := range pools {
				if containz(connections, pool, current) {
					pool2 := sortKey(pool + "," + current)
					delete(pools, pool)
					pools[pool2] = struct{}{}
				}
			}
		}
		old = pools
	}

	longest := ""
	for pool := range pools {
		if len(pool) > len(longest) {
			longest = pool
		}
	}
	return longest
}

func parseConnections(input []string) map[string][]string {
	connections := make(map[string][]string)
	for _, line := range input {
		split := strings.Split(line, "-")
		connections[split[0]] = append(connections[split[0]], split[1])
		connections[split[1]] = append(connections[split[1]], split[0])
	}
	return connections
}

func hash(m map[string]struct{}) int {
	h := 0
	for key := range m {
		h += len(key)
	}
	return h
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func getKeys(s string) []string {
	return strings.Split(s, ",")
}

func sortKey(s string) string {
	keys := getKeys(s)
	sort.Strings(keys)
	return strings.Join(keys, ",")
}

func containz(connections map[string][]string, s string, e string) bool {
	for _, k := range getKeys(s) {
		if !contains(connections[k], e) {
			return false
		}
	}
	return true
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
