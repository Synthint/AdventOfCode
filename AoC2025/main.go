package main

import (
	day01 "AoC2025/Day_01"
	day02 "AoC2025/Day_02"
	day03 "AoC2025/Day_03"
	day04 "AoC2025/Day_04"

	// day05 "AoC2025/Day_05"
	// day06 "AoC2025/Day_06"
	// day07 "AoC2025/Day_07"
	// day08 "AoC2025/Day_08"
	// day09 "AoC2025/Day_09"
	// day10 "AoC2025/Day_10"
	// day11 "AoC2025/Day_11"
	// day12 "AoC2025/Day_12"
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
	3: day03.Solve,
	4: day04.Solve,
	// 5:  day05.Solve,
	// 6:  day06.Solve,
	// 7:  day07.Solve,
	// 8:  day08.Solve,
	// 9:  day09.Solve,
	// 10: day10.Solve,
	// 11: day11.Solve,
	// 12: day12.Solve,
}

func main() {
	var day int

	fmt.Print("Enter the day number (1-12): ")

	fmt.Scanf("%d", &day)

	fmt.Printf("Running Day %d: \n", day)

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
