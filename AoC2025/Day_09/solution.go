package template

import (
	"AoC2025/utils"
	"fmt"
	"sort"
	"strings"
)

type point struct {
	x int
	y int
}

type rect struct {
	corner1 point
	corner2 point
	area    int
}

func twoPointsToRect(p1, p2 point) rect {
	x1 := min(p1.x, p2.x) // Get top left bottom right corners
	x2 := max(p1.x, p2.x)
	y1 := min(p1.y, p2.y)
	y2 := max(p1.y, p2.y)

	return rect{
		corner1: point{x1, y1},
		corner2: point{x2, y2},
		area:    (x2 - x1 + 1) * (y2 - y1 + 1),
	}
}

func pointsToRects(corners []point) []rect {
	var rectangles []rect
	for x := 0; x < len(corners)-1; x++ {
		for y := x + 1; y < len(corners); y++ {
			rectangles = append(rectangles, twoPointsToRect(corners[x], corners[y]))
		}
	}

	return rectangles
}

func stringToPoint(in string) point {
	vals := strings.Split(in, ",")
	return point{
		x: utils.Atoi(vals[0]),
		y: utils.Atoi(vals[1]),
	}
}

func SolvePartOne(file_slice []string) string {
	corners := make([]point, len(file_slice))
	for x, line := range file_slice {
		corners[x] = stringToPoint(line)
	}

	rects := pointsToRects(corners)

	sort.Slice(rects, func(i, j int) bool {
		return rects[i].area < rects[j].area
	})

	return fmt.Sprintf("%d", rects[len(rects)-1].area)
}

func makeInsideGrid(points []point) ([][]bool, map[int]int, map[int]int) {
	n := len(points)

	all_edge_tiles := make(map[point]struct{})

	for i := 0; i < n; i++ {
		from := points[i]
		to := points[(i+1)%n]

		if from.x == to.x {
			step := 1
			if to.y < from.y {
				step = -1
			}
			for y := from.y; y != to.y+step; y += step {
				all_edge_tiles[point{from.x, y}] = struct{}{}
			}
		} else {
			step := 1
			if to.x < from.x {
				step = -1
			}
			for x := from.x; x != to.x+step; x += step {
				all_edge_tiles[point{x, from.y}] = struct{}{}
			}
		}
	}

	x_set := map[int]struct{}{}
	y_set := map[int]struct{}{}
	for _, p := range points {
		x_set[p.x] = struct{}{}
		y_set[p.y] = struct{}{}
	}

	shrink_x := keysInt(x_set)
	shrink_y := keysInt(y_set)
	sort.Ints(shrink_x)
	sort.Ints(shrink_y)

	x_index := map[int]int{}
	y_index := map[int]int{}
	for i, x := range shrink_x {
		x_index[x] = i
	}
	for i, y := range shrink_y {
		y_index[y] = i
	}

	w, h := len(shrink_x), len(shrink_y)
	insides := make([][]bool, h)
	for i := range insides {
		insides[i] = make([]bool, w)
	}

	vertical_edges_by_y := map[int][]int{}
	prev := points[n-1]

	for _, cur := range points {
		if prev.x == cur.x {
			y_from := min(prev.y, cur.y)
			y_to := max(prev.y, cur.y)
			for _, y := range shrink_y {
				if y_from < y && y <= y_to {
					vertical_edges_by_y[y] = append(vertical_edges_by_y[y], prev.x)
				}
			}
		}
		prev = cur
	}

	is_inside_polygon := func(x, y int) bool {
		crossings := 0
		for _, edge_x := range vertical_edges_by_y[y] {
			if x <= edge_x {
				crossings++
			}
		}
		return crossings%2 == 1
	}

	for yi, y := range shrink_y {
		for xi, x := range shrink_x {
			if _, ok := all_edge_tiles[point{x, y}]; ok {
				insides[yi][xi] = true
			} else if is_inside_polygon(x, y) {
				insides[yi][xi] = true
			}
		}
	}

	return insides, x_index, y_index
}

func isRectValid(r rect, insides [][]bool, x_index, y_index map[int]int) bool {
	x1 := x_index[r.corner1.x]
	x2 := x_index[r.corner2.x]
	y1 := y_index[r.corner1.y]
	y2 := y_index[r.corner2.y]

	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			if !insides[y][x] {
				return false
			}
		}
	}
	return true
}

func keysInt(m map[int]struct{}) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func SolvePartTwo(file_slice []string) string {
	poly_points := make([]point, len(file_slice))
	for i, line := range file_slice {
		poly_points[i] = stringToPoint(line)
	}

	rects := pointsToRects(poly_points)
	insides, x_index, y_index := makeInsideGrid(poly_points)

	var valid_rects []rect
	for _, r := range rects {
		if isRectValid(r, insides, x_index, y_index) {
			valid_rects = append(valid_rects, r)
		}
	}

	sort.Slice(valid_rects, func(i, j int) bool {
		return valid_rects[i].area < valid_rects[j].area
	})

	return fmt.Sprintf("%d", valid_rects[len(valid_rects)-1].area)

}

func Solve(input []string) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
