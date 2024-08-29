package main

import (
	"adventOfGode/common"
	"os"
	"regexp"
	"slices"
	"strings"
)

var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// grab from, to
var re = regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z]) can begin.$`)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2018/day7/" + "input.txt")
	var input = string(bytes)
	println(part1(input))
}

func pop(source []string) (string, []string) {
	x, source := source[0], source[1:]
	return x, source
}

func part1(input string) string {
	var sArr = strings.Split(input, "\n")

	downstream := make(map[string][]string)
	upstream := make(map[string][]string)
	for _, s := range sArr {
		m := re.FindStringSubmatch(s)
		from, to := m[1], m[2]
		downstream[from] = append(downstream[from], to)
		upstream[to] = append(upstream[to], from)
	}

	seen := make([]string, 0)
	startingNodes := getStartingNodes(upstream)
	for _, i := range startingNodes {
		curr := []string{i}
		for len(curr) > 0 {
			c, t := pop(curr)
			curr = t
			seen = append(seen, c)
			//only append if all prerequisites are met
			for _, v := range downstream[c] {
				if isAcceptedDownstream(upstream, v, seen) {
					curr = append(curr, v)
				}
			}
			slices.Sort(curr)
		}
		if len(seen) == len(common.Keys(upstream)) {
			break
		}
	}
	return strings.Join(seen, "")
}

func isAcceptedDownstream(upstream map[string][]string, v string, seen []string) bool {
	accepted := true
	for _, a := range upstream[v] {
		if !slices.Contains(seen, a) {
			accepted = false
			break
		}
	}
	return accepted
}

func getStartingNodes(upstream map[string][]string) []string {
	starts := make([]string, 0)
	for _, a := range alphabet {
		if _, ok := upstream[string(a)]; !ok {
			starts = append(starts, string(a))
		}
	}
	return starts
}
