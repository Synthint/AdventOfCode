package template

import (
	utils "AoC2025/utils"
	"fmt"
	"math"
)

func SolvePartOne(file_slice []string) string {
	start := 50
	pass_0 := 0

	current_position := start
	fmt.Printf("aa")
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
	fmt.Printf("aa")
	for _, line := range file_slice {
		clicks := line[1:]
		clicks_int := utils.Atoi(clicks)
		if line[:1] == "L" {
			clicks_int *= -1
		}
		old_position := current_position
		current_position += clicks_int
		if current_position == 0 || current_position%100 == 0 || math.Abs(float64(current_position-old_position)) > 100 {
			current_position = 0
			pass_0++
		}
	}

	return fmt.Sprintf("%d", pass_0)
}

func Solve(input []string) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
