package template

import (
	utils "AoC2025/utils"
	"fmt"
)

func SolvePartOne(file_slice []string) string {
	start := 50
	pass_0 := 0

	current_position := start
	for _, line := range file_slice {
		clicks := line[1:]
		clicks_int := utils.Atoi(clicks)
		if line[:1] == "L" {
			clicks_int *= -1
		}
		current_position += clicks_int
		if current_position == 0 || current_position%100 == 0 {
			current_position = 0
			pass_0++
		}
	}
	return fmt.Sprintf("%d", pass_0)
}

func SolvePartTwo(file_slice []string) string {
	start := 50
	pass_0 := 0

	current_position := start
	for _, line := range file_slice {
		clicks := line[1:]
		clicks_int := utils.Atoi(clicks)

		if clicks_int > 100 {
			pass_0 += clicks_int / 100
			clicks_int = clicks_int % 100
		}

		if line[:1] == "L" {
			clicks_int *= -1
		}
		old_position := current_position
		current_position += clicks_int
		passed_0 := false
		if current_position == 0 {
			pass_0++
			passed_0 = true
		} else if current_position/100 > 0 {
			pass_0 += current_position / 100
			passed_0 = true
		}

		if current_position < 0 {
			current_position += 100
		}
		current_position = current_position % 100

		if line[:1] == "L" && old_position < current_position && old_position != 0 && !passed_0 {
			pass_0++
			passed_0 = true
		} else if line[:1] == "R" && old_position > current_position && old_position != 0 && !passed_0 {
			pass_0++
			passed_0 = true
		}
	}
	return fmt.Sprintf("%d", pass_0)
}

func Solve(input []string) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
