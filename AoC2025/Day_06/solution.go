package template

import (
	"AoC2025/utils"
	"fmt"
	"strings"
)

func SolvePartOne(file_slice []string) string {
	cephalapod_homework := make([][]string, len(file_slice))
	for ind, line := range file_slice {
		cephalapod_homework[ind] = strings.Split(strings.Join(strings.Fields(line), " "), " ")
	}

	homework_nums := make([][]int, len(cephalapod_homework[0]))

	// 2. Loop through and initialize each inner slice with the desired number of columns
	for i := 0; i < len(cephalapod_homework[0]); i++ {
		homework_nums[i] = make([]int, len(cephalapod_homework))
	}

	homework_ops := make([]rune, len(cephalapod_homework[0]))

	for line_num, line := range cephalapod_homework {
		for problem_num, number_or_op := range line {
			if number_or_op == "*" || number_or_op == "+" {
				homework_ops[problem_num] = []rune(number_or_op)[0]
				if number_or_op == "*" {
					homework_nums[problem_num][line_num] = 1
					// Default 0 to fill in spot left by operation messes up multiplication,
					// replace with 1 rather than fix the lists :)
				}
			} else {
				homework_nums[problem_num][line_num] = utils.Atoi(number_or_op)
			}
		}
	}

	grand_total := 0
	for problem_num, number_list := range homework_nums {
		total := 0
		if homework_ops[problem_num] == '*' {
			total = 1
		}
		for _, number := range number_list {
			if homework_ops[problem_num] == '*' {
				total *= number
			} else if homework_ops[problem_num] == '+' {
				total += number
			}
		}
		grand_total += total
	}

	// fmt.Println("\n\n")
	return fmt.Sprintf("%d", grand_total)
}

func SolvePartTwo(file_slice []string) string {
	chars := make([][]rune, len(file_slice))
	for x, line := range file_slice {
		line += " "
		chars[x] = []rune(line)
	}

	temp_total := 0
	grand_total := 0
	curr_op := ' '
	for col := range len(chars[0]) {
		num_str := ""
		for row := range len(chars) - 1 {
			num_str += string(chars[row][col])
		}
		blank_column := strings.TrimSpace(num_str) == ""
		if !blank_column {
			if chars[len(chars)-1][col] != ' ' {
				curr_op = chars[len(chars)-1][col]
			}
			col_value := utils.Atoi(strings.TrimSpace(num_str))

			if curr_op == '+' {
				temp_total += col_value
			} else if curr_op == '*' {
				if temp_total == 0 {
					temp_total = 1
				}
				temp_total *= col_value
			}
		} else {
			// fmt.Println("-----")
			// fmt.Println(temp_total, " ", string(curr_op), "\n\n")
			grand_total += temp_total
			// fmt.Println("----- TOT: ", grand_total)
			temp_total = 0
		}
	}

	return fmt.Sprintf("%d", grand_total)
}

func Solve(input []string) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
