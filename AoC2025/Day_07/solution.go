package template

import (
	"fmt"
	"strings"
)

func calculateTachyonStream(file_slice []string) (int, [][]string) {
	has_beam := make([]bool, len(file_slice[0]))
	beam_start_index := strings.Index(file_slice[0], "S")
	has_beam[beam_start_index] = true
	splits := 0
	calculated_stream := make([][]string, len(file_slice))
	for x, line := range file_slice {
		calculated_stream[x] = make([]string, len(line))
	}

	// for ind, line := range file_slice {
	calculated_stream[0] = strings.Split(file_slice[0], "")
	// }

	for x, line := range file_slice[1:] {
		line_slice := strings.Split(line, "")
		new_beam := make([]bool, len(has_beam))
		// fmt.Println(has_beam)
		for x := range len(line_slice) {
			if has_beam[x] && line_slice[x] == "^" {
				splits += 1
				if x > 0 {
					new_beam[x-1] = true
				}
				if x < len(new_beam)-1 {
					new_beam[x+1] = true
				}
			}
			if has_beam[x] && line_slice[x] == "." {
				new_beam[x] = true
				line_slice[x] = "|"
			}
		}
		calculated_stream[x+1] = line_slice
		// fmt.Println(line_slice)
		copy(has_beam, new_beam)
	}

	return splits, calculated_stream
}

func SolvePartOne(file_slice []string) string {
	splits, _ := calculateTachyonStream(file_slice)

	// for _, line := range calculated_stream {
	// 	fmt.Println(line)
	// }

	return fmt.Sprintf("%d", splits)
}

func walkTachyonStream(pos int, depth int, stream [][]string, already_traveled [][]int) int {

	if depth+2 >= len(stream) {
		// fmt.Println("End Ahead")
		return 1
	}
	if stream[depth+2][pos] == "^" {
		// fmt.Println("depth: ", depth)
		if already_traveled[depth+2][pos] > -1 {
			// fmt.Println("I SAVED TIME")
			return already_traveled[depth+2][pos]
		} else {
			already_traveled[depth+2][pos] = walkTachyonStream(pos-1, depth+2, stream, already_traveled) + walkTachyonStream(pos+1, depth+2, stream, already_traveled)
			return already_traveled[depth+2][pos]
		}

	}
	return walkTachyonStream(pos, depth+2, stream, already_traveled)
}

func SolvePartTwo(file_slice []string) string {
	_, calculated_stream := calculateTachyonStream(file_slice)

	start := 0
	for x, _ := range calculated_stream[0] {
		if calculated_stream[0][x] == "S" {
			start = x
		}
	}

	already_traveled := make([][]int, len(calculated_stream))
	for x := range len(calculated_stream) {
		already_traveled[x] = make([]int, len(calculated_stream[x]))
		for y := range already_traveled[x] {
			already_traveled[x][y] = -1
		}
	}
	return fmt.Sprintf("%d", walkTachyonStream(start, 0, calculated_stream, already_traveled))
}

func Solve(input []string, arg int) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
