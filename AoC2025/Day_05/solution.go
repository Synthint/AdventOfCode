package template

import (
	"AoC2025/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

type Range struct {
	min int
	max int
}

func SolvePartOne(file_slice []string) string {
	ranges := []string{}
	ids := []int{}
	blankLineHit := false
	for _, line := range file_slice {
		if line == "" {
			blankLineHit = true
			continue
		}
		if !blankLineHit {
			ranges = append(ranges, line)
		} else {
			ids = append(ids, utils.Atoi(line))
		}
	}

	count := 0
	for _, id := range ids {
		_, is_fresh := numberInAnyOfMultipleRanges(id, ranges, "")
		if is_fresh {
			count += 1
		}
	}

	return fmt.Sprintf("%d", count)
}

func numberInAnyOfMultipleRanges(number int, ranges []string, exclude string) (int, bool) {
	for x, r := range ranges {
		if r == exclude || r == "-NONE-" {
			continue
		}
		var min, max int
		min_max := strings.Split(r, "-")
		min = utils.Atoi(min_max[0])
		max = utils.Atoi(min_max[1])

		if number >= min && number <= max {
			return x, true
		}
	}
	return -1, false
}

func condenseRanges(ranges []string) []string {
	parsed := []Range{}

	for _, r := range ranges {
		parts := strings.Split(r, "-")
		parsed = append(parsed, Range{
			min: utils.Atoi(parts[0]),
			max: utils.Atoi(parts[1]),
		})
	}

	sort.Slice(parsed, func(i, j int) bool {
		return parsed[i].min < parsed[j].min
	})

	merged := []Range{parsed[0]}

	for _, curr := range parsed[1:] {
		last := &merged[len(merged)-1]

		if curr.min <= last.max+1 {
			last.min = int(math.Min(float64(last.min), float64(curr.min)))
			last.max = int(math.Max(float64(last.max), float64(curr.max)))
		} else {
			merged = append(merged, curr)
		}
	}

	result := []string{}
	for _, r := range merged {
		result = append(result, fmt.Sprintf("%d-%d", r.min, r.max))
	}

	return result
}

func SolvePartTwo(file_slice []string) string {
	ranges := []string{}
	blankLineHit := false
	for _, line := range file_slice {
		if line == "" {
			blankLineHit = true
			continue
		}
		if !blankLineHit {
			ranges = append(ranges, line)
		} else {
			break
		}
	}

	new_ranges := condenseRanges(ranges)
	// fmt.Println(new_ranges)
	// fmt.Println(ranges)
	total := 0
	for _, r := range new_ranges {
		if r == "-NONE-" {
			continue
		}
		var min, max int
		min_max := strings.Split(r, "-")
		min = utils.Atoi(min_max[0])
		max = utils.Atoi(min_max[1])

		range_count := (max - min) + 1

		total += range_count

		// fmt.Printf("Range: %s , count: %d \n", r, range_count)
	}

	return fmt.Sprintf("%d", total)

}

func Solve(input []string) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
