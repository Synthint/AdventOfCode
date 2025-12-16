package day02

import (
	utils "AoC2025/utils"
	"fmt"
	"strings"
)

func SolvePartOne(file_slice []string) string {
	// repeat_regex = "^(\\d+)\\1$" No backreferences in Go regex :(

	if len(file_slice) > 1 {
		fmt.Println("Day 2, Part 1 shoud contain only 1 line")
	}

	line := file_slice[0]

	line_ranges := strings.Split(line, ",")

	var solution_int int = 0

	for _, line_range := range line_ranges {
		bounds := strings.Split(line_range, "-")
		if len(bounds) != 2 {
			fmt.Printf("Invalid range: %s\n", line_range)
			continue
		}

		for i := utils.Atoi(bounds[0]); i <= utils.Atoi(bounds[1]); i++ {
			num_str := fmt.Sprintf("%d", i)
			if isRepeatedTwice(num_str) {
				solution_int += i
			}
		}

	}

	return fmt.Sprintf("%d", solution_int)

}

func isRepeatedTwice(s string) bool {
	n := len(s)
	if n%2 != 0 {
		return false
	}
	half := n / 2
	return s[:half] == s[half:]
}

func isRepeatedAny(s string) bool {
	n := len(s)
	for size := 1; size <= n/2; size++ {
		if n%size != 0 {
			continue
		}
		repeated := true
		pattern := s[:size]
		for i := size; i < n; i += size {
			if s[i:i+size] != pattern {
				repeated = false
				break
			}
		}
		if repeated {
			return true
		}
	}
	return false
}

func SolvePartTwo(file_slice []string) string {
	// repeat_regex = "^(\\d+)\\1+$" No backreferences in Go regex :(

	if len(file_slice) > 1 {
		fmt.Println("Day 2, Part 1 shoud contain only 1 line")
	}

	line := file_slice[0]

	line_ranges := strings.Split(line, ",")

	var solution_int int = 0

	for _, line_range := range line_ranges {
		bounds := strings.Split(line_range, "-")
		if len(bounds) != 2 {
			fmt.Printf("Invalid range: %s\n", line_range)
			continue
		}

		for i := utils.Atoi(bounds[0]); i <= utils.Atoi(bounds[1]); i++ {
			num_str := fmt.Sprintf("%d", i)
			if isRepeatedAny(num_str) {
				solution_int += i
			}
		}

	}

	return fmt.Sprintf("%d", solution_int)
}

func Solve(input []string) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
