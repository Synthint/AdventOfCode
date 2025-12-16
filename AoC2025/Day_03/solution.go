package template

import (
	"AoC2025/utils"
	"fmt"
	"strings"
)

func SolvePartOne(file_slice []string) string {
	grand_total := 0

	for _, bank := range file_slice {
		bank_total := 0
		max := 0
		for _, jolt := range strings.Split(bank, "") {
			if utils.Atoi(jolt) > max {
				max = utils.Atoi(jolt)
			}
		}

		bank_split := strings.Split(bank, fmt.Sprintf("%d", max))
		if len(bank_split) == 2 {
			mini_bank := bank_split[1]
			max_second_digit := false
			if len(bank_split[1]) == 0 {
				max_second_digit = true
				mini_bank = bank_split[0]
			}

			max_2 := 0
			for _, jolt := range strings.Split(mini_bank, "") {
				if utils.Atoi(jolt) > max_2 {
					max_2 = utils.Atoi(jolt)
				}
			}

			if max_second_digit {
				bank_total = utils.Atoi(fmt.Sprintf("%d%d", max_2, max))
			} else {
				bank_total = utils.Atoi(fmt.Sprintf("%d%d", max, max_2))
			}
		} else {
			bank_total = utils.Atoi(fmt.Sprintf("%d%d", max, max))
		}

		fmt.Printf("Bank %s had a joltage of %d \n", bank, bank_total)
		grand_total += bank_total
	}

	return fmt.Sprintf("%d", grand_total)
}

func SolvePartTwo(file_slice []string) string {
	return "Not Implemented, first line of input to confirm it was read: " + file_slice[0]
}

func Solve(input []string) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
