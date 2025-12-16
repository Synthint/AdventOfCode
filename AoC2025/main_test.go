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

			t.Run("First Problem", func(t *testing.T) {
				if part_one_solution != part_one_expected {
					t.Errorf("expected %q, got %q", part_one_expected, part_one_solution)
				}
			})

			t.Run("Second Problem", func(t *testing.T) {
				if part_two_solution != part_two_expected {
					t.Errorf("expected %q, got %q", part_two_expected, part_two_solution)
				}
			})
		})
	}
}
