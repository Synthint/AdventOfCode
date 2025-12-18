package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

func TestSolutions(t *testing.T) {
	// Deterministic order
	var days []int
	for day := range solvers {
		days = append(days, day)
	}
	sort.Ints(days)
	fmt.Println(strings.Repeat("-", 69))
	for _, day := range days {
		day := day // capture loop variable

		t.Run(fmt.Sprintf("Day %02d", day), func(t *testing.T) {
			testPath := fmt.Sprintf("Day_%02d/test.txt", day)

			file, err := os.Open(testPath)
			if err != nil {
				t.Fatalf("cannot open test input %s: %v", testPath, err)
			}
			defer file.Close()

			var test_slice []string
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				test_slice = append(test_slice, scanner.Text())
			}

			part_one_expected := strings.TrimSpace(test_slice[0])
			part_two_expected := strings.TrimSpace(test_slice[1])

			input_slice := test_slice[2:]

			part_one_solution, part_two_solution := solvers[day](input_slice)

			day_result_str := "\u2705"
			part_one_result_str := "\u2705"
			if part_one_solution != part_one_expected {
				part_one_result_str = "\u274C"
				day_result_str = "\u274C"
			}
			part_two_result_str := "\u2705"
			if part_two_solution != part_two_expected {
				part_two_result_str = "\u274C"
				day_result_str = "\u274C"
			}

			fmt.Printf("|%s %s %-2d | %s %-14s %s | %s %-15s %s |\n",
				day_result_str, "Day", day,
				"Part 1:", part_one_solution, part_one_result_str,
				"Part 2:", part_two_solution, part_two_result_str)
			fmt.Println(strings.Repeat("-", 69))
			t.Run("First Problem", func(t *testing.T) {
				if part_one_solution != part_one_expected {
					t.Errorf("Day %d Part 1 expected %q, got %q", day, part_one_expected, part_one_solution)
				}
			})

			t.Run("Second Problem", func(t *testing.T) {
				if part_two_solution != part_two_expected {
					t.Errorf("Day %d Part 2 expected %q, got %q", day, part_two_expected, part_two_solution)
				}
			})
		})

	}
	fmt.Println("")
}
