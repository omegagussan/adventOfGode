package main

import (
	"adventOfGode/common"
	"os"
	"strings"
)

type point struct {
	x int
	y int
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2018/day6/" + "input.txt")
	var input = string(bytes)
	println(part1(input))
	println(part2(input))
}

func part1(input string) int64 {
	points := parsePoints(input)
	minX, maxX, minY, maxY := findRange(points)

	//count of points closest to the edge (are infinite)
	isInfinite := getMapOfClosestPointsToEdge(minX, maxX, minY, maxY, points)

	// count of points in each area
	areaByPoint := getMapOfClosesPointsInArea(minX, maxX, minY, maxY, points)

	return getLargestFiniteArea(points, isInfinite, areaByPoint)
}

func part2(input string) int64 {
	points := parsePoints(input)
	minX, maxX, minY, maxY := findRange(points)

	//find the sum of the distances to all points
	d := make(map[point]int)
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			p := point{i, j}
			for _, x := range points {
				d[p] += manhattanDistance(p, x)
			}
		}
	}
	keep := func(p common.Pair[point, int]) bool {
		if p.Value < 10000 {
			return true
		}
		return false
	}
	getKey := func(p common.Pair[point, int]) point {
		return p.Key
	}

	filteredPointsWithScores := common.Filter(common.ToPairList(d), keep)
	filteredPoints := common.Map(filteredPointsWithScores, getKey)
	return int64(len(filteredPoints))
}

func getLargestFiniteArea(points []point, isInfinite map[point]int, areaByPoint map[point]int) int64 {
	maxArea := 0
	for _, p := range points {
		if isInfinite[p] > 0 {
			//this means its infinite
			continue
		}
		if areaByPoint[p] > maxArea {
			maxArea = areaByPoint[p]
		}
	}
	return int64(maxArea)
}

func getMapOfClosesPointsInArea(minX int, maxX int, minY int, maxY int, points []point) map[point]int {
	d := make(map[point]int)
	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			p := point{i, j}
			xp := point{0, 0}
			m := common.MaxInt
			// If a point is equally far from two or more coordinates, it doesn't count as being closest to any.
			tie := false
			for _, x := range points {
				dist := manhattanDistance(p, x)
				if dist < m {
					tie = false
					m = dist
					xp = x
				} else if dist == m {
					tie = true
				}
			}
			if !tie {
				d[xp]++
			}
		}
	}
	return d
}

func getMapOfClosestPointsToEdge(minX int, maxX int, minY int, maxY int, points []point) map[point]int {
	i := make(map[point]int)
	edges := getEdges(minX, maxX, minY, maxY)

	for _, p := range edges {
		m := common.MaxInt
		px := point{0, 0}
		for _, x := range points {
			dist := manhattanDistance(p, x)
			if dist < m {
				m = dist
				px = x
			}
		}
		i[px]++
	}
	return i
}

func getEdges(minX int, maxX int, minY int, maxY int) []point {
	edges := make(map[point]bool)
	for i := minX; i <= maxX; i++ {
		edges[point{i, minY}] = true
		edges[point{i, maxY}] = true
	}

	for i := minY; i <= maxY; i++ {
		edges[point{minX, i}] = true
		edges[point{maxX, i}] = true
	}
	return common.Keys(edges)
}

func findRange(points []point) (int, int, int, int) {
	minX, maxX, minY, maxY := points[0].x, points[0].x, points[0].y, points[0].y
	for _, p := range points {
		if p.x < minX {
			minX = p.x
		}
		if p.x > maxX {
			maxX = p.x
		}
		if p.y < minY {
			minY = p.y
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	return minX, maxX, minY, maxY

}

func parsePoints(input string) []point {
	points := make([]point, 0)
	for _, line := range strings.Split(input, "\n") {
		xy := strings.Split(line, ", ")
		xyInt := common.Map(xy, common.ToInt)
		points = append(points, point{xyInt[0], xyInt[1]})
	}
	return points
}

func manhattanDistance(p1, p2 point) int {
	return common.AbsInt(p1.x-p2.x) + common.AbsInt(p1.y-p2.y)
}

// 5186
// 3998 your answer is too high
// 3015 too low
// 3989
