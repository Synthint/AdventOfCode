package template

import (
	"AoC2025/utils"
	"fmt"
	"log"
	"strings"
)

func SolvePartOne(file_slice []string) string {
	return solveForNumDigits(file_slice, 2)
}

func SolvePartTwo(file_slice []string) string {
	return solveForNumDigits(file_slice, 12)
}

func getLargestDigitBeforeIndex(bank string, index int) (int, int, error) {
	if index < 0 || index > len(bank) {
		return 0, 0, fmt.Errorf("index out of bounds: %d", index)
	}

	bank_before_index := bank[:len(bank)-index]

	max := 0
	max_index := 0
	for i, value := range strings.Split(bank_before_index, "") {
		if utils.Atoi(value) > max {
			max = utils.Atoi(value)
			max_index = i
		}
	}

	return max_index, max, nil
}

func solveForNumDigits(file_slice []string, num_digits_wanted int) string {
	grand_total := 0
	for _, bank := range file_slice {
		cutoff_index := num_digits_wanted - 1
		bank_total_str := ""
		remaining_bank := bank
		for cutoff_index > 0 {
			max_index, max, err := getLargestDigitBeforeIndex(remaining_bank, cutoff_index)
			if err != nil {
				log.Fatal(err)
			}
			removal := max_index + 1
			remaining_bank = remaining_bank[removal:]
			cutoff_index -= 1
			bank_total_str += fmt.Sprintf("%d", max)
		}
		_, final_digit, err := getLargestDigitBeforeIndex(remaining_bank, 0)
		if err != nil {
			log.Fatal(err)
		}
		bank_total_str += fmt.Sprintf("%d", final_digit)

		// fmt.Printf("Bank %s had a joltage of %s \n\n\n", bank, bank_total_str)
		grand_total += utils.Atoi(bank_total_str)
	}
	return fmt.Sprintf("%d", grand_total)
}

func Solve(input []string, arg int) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
