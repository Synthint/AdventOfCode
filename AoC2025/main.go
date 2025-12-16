package main

import (
	day01 "AoC2025/Day_01"
	day02 "AoC2025/Day_02"
	"bufio"
	"fmt"
	"log"
	"os"
)

type Solver func([]string) (string, string)

// Solver array
var solvers = map[int]Solver{
	1: day01.Solve,
	2: day02.Solve,
}

func main() {
	var day int
	var mode string

	fmt.Print("Enter the day number (1-12) then the mode (test,run,both): ")

	fmt.Scanf("%d %s", &day, &mode)

	fmt.Printf("Running Day %d with in %s mode: \n", day, mode)

	if mode != "test" && mode != "run" && mode != "both" && mode != "" {
		log.Fatalf("Invalid mode: %s. Must be 'test', 'run', or 'both'.", mode)
	}

	if mode == "test" || mode == "both" || mode == "" {
		test_for_day := fmt.Sprintf("Day_%02d/test.txt", day)
		fmt.Printf("Using test file: %s\n", test_for_day)
		test_file, err := os.Open(test_for_day)
		if err != nil {
			log.Fatal(err)
		}
		defer test_file.Close()

		var test_file_slice []string
		scanner := bufio.NewScanner(test_file)
		for scanner.Scan() {
			test_file_slice = append(test_file_slice, scanner.Text())
		}

		part_one_solution, part_two_solution := solvers[day](test_file_slice)
		fmt.Println("Test Solutions:")
		fmt.Printf("\tDay %d - Part One Solution: %s\n", day, part_one_solution)
		fmt.Printf("")
		fmt.Printf("\tDay %d - Part Two Solution: %s\n", day, part_two_solution)
	}

	if mode == "run" || mode == "both" || mode == "" {
		input_file := fmt.Sprintf("Day_%02d/input.txt", day)
		fmt.Printf("Using test file: %s\n", input_file)
		file, err := os.Open(input_file)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		var file_slice []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			file_slice = append(file_slice, scanner.Text())
		}

		part_one_solution, part_two_solution := solvers[day](file_slice)
		fmt.Println("Run Solutions:")
		fmt.Printf("\tDay %d - Part One Solution: %s\n", day, part_one_solution)
		fmt.Printf("")
		fmt.Printf("\tDay %d - Part Two Solution: %s\n", day, part_two_solution)
	}

}
