package main

import (
	day01 "AoC2025/Day_01"
	day02 "AoC2025/Day_02"
	day03 "AoC2025/Day_03"
	day04 "AoC2025/Day_04"
	day05 "AoC2025/Day_05"
	day06 "AoC2025/Day_06"
	day07 "AoC2025/Day_07"
	"AoC2025/utils"
	"flag"
	"maps"
	"slices"
	"sort"
	"strings"

	day08 "AoC2025/Day_08"
	day09 "AoC2025/Day_09"

	// day10 "AoC2025/Day_10"

	// day11 "AoC2025/Day_11"
	// day12 "AoC2025/Day_12"

	"bufio"
	"fmt"
	"log"
	"os"
	// "os"
)

type Solver func([]string, int) (string, string)

// Solver array
var solvers = map[int]Solver{
	1: day01.Solve,
	2: day02.Solve,
	3: day03.Solve,
	4: day04.Solve,
	5: day05.Solve,
	6: day06.Solve,
	7: day07.Solve,
	8: day08.Solve,
	9: day09.Solve,
	// 10: day10.Solve,
	// 11: day11.Solve,
	// 12: day12.Solve,
}

func main() {
	days := slices.Collect(maps.Keys(solvers))
	sort.Slice(days, func(i, j int) bool {
		return days[i] < days[j]
	})

	DEFAULT := ""
	DEFAULT_PATH := "Day_%02d/input.txt" // Searches for file differently than either option below

	allFlagPtr := flag.Bool("all", false, "Run all problems, cannot be used with -day or -file. \nBy default, without -path option will look for input.txt in \n./Day_xx/ where xx is the 2 digit day including a leading 0 on single digit days")
	dayFlagPtr := flag.Int("day", -1, "The day to run, do not include a leading 0 on single digit days. \nBy default, without -path or -file options will look for input.txt in \n./Day_xx/ where xx is the 2 digit day including a leading 0 on single digit days")
	fileFlagPtr := flag.String("file", DEFAULT, "A single file to use as input, cannot be used with -path or -all.")
	pathFlagPtr := flag.String("path", DEFAULT, "Path to find input files, cannot be used with -file. \nFiles must be named input_xx.txt where xx is the 2 digit day. \ninclude leading 0 on single digit days.")

	if *fileFlagPtr != DEFAULT && *pathFlagPtr != DEFAULT {
		fmt.Println("Incompatible Options: file and path. specify the full path in the file option if needed.Exiting")
		os.Exit(1)
	}
	if *allFlagPtr && *dayFlagPtr != -1 {
		fmt.Println("Incompatible Options: all and day. Select either a single day or all days. Exiting")
		os.Exit(1)
	}
	if *allFlagPtr && *fileFlagPtr != DEFAULT {
		fmt.Println("Incompatible Options: all and file. Cannot run all days with single file. Exiting")
		os.Exit(1)
	}

	flag.Parse()

	var results [][]string
	day_args := flag.Args()
	if *allFlagPtr || len(day_args) > 0 {
		if len(day_args) > 0 {
			day_args_int := make([]int, len(day_args))
			for x, d := range day_args {
				day_args_int[x] = utils.Atoi(d)
			}
			days = day_args_int
		}

		for _, day := range days {
			day_path := fmt.Sprintf(DEFAULT_PATH, day)
			if *fileFlagPtr != DEFAULT {
				os.Exit(1)
			} else if *pathFlagPtr != DEFAULT {
				day_path = *pathFlagPtr + fmt.Sprintf("input_%02d.txt", day)
			}
			fmt.Println("Running Day ", day, " With File ", day_path, "...")
			res1, res2 := runDay(day, day_path)
			results = append(results, []string{fmt.Sprintf("Day %02d", day), res1, res2})
		}
	} else if *dayFlagPtr != -1 {
		day := *dayFlagPtr
		day_path := fmt.Sprintf(DEFAULT_PATH, day)
		if *fileFlagPtr != DEFAULT {
			day_path = fmt.Sprintf(*fileFlagPtr, day)
		} else if *pathFlagPtr != DEFAULT {
			day_path = *fileFlagPtr + fmt.Sprintf("input_%02d.txt", day)
		}
		fmt.Println("Running Day ", day, " With File ", day_path, "...")
		res1, res2 := runDay(day, day_path)
		results = append(results, []string{fmt.Sprintf("Day %02d", day), res1, res2})
	}
	fmt.Println("")
	fmt.Println(strings.Repeat("-", 67))
	fmt.Println("|", strings.Repeat("~ ", 8), "Advent of Code 2025 SOLUTIONS", strings.Repeat(" ~", 8), "|")
	fmt.Println(strings.Repeat("-", 67))
	for _, res := range results {
		day := res[0]
		part_one_solution := res[1]
		part_two_solution := res[2]
		fmt.Printf("| %s | %s %-17s | %s %-18s |\n",
			day,
			"Part 1:", part_one_solution,
			"Part 2:", part_two_solution)
		fmt.Println(strings.Repeat("-", 67))
	}
}

func runDay(day int, input_file string) (string, string) {
	if !slices.Contains(slices.Collect(maps.Keys(solvers)), day) {
		panic(fmt.Sprintf("ERROR: Day %d Not Found", day))
	}

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

	arg := -1
	if day == 8 {
		arg = 1000
	}

	return solvers[day](file_slice, arg)
}
