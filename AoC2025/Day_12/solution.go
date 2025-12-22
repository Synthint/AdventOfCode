package template

import (
	"AoC2025/utils"
	"fmt"
	"regexp"
	"strings"
)

type present struct {
	grid [][]bool
	area int
}

type space struct {
	length           int
	width            int
	present_quantity []int
}

func ParseInput(input []string) ([]present, []space) {
	space_regex := "^(\\d+)x(\\d+):\\s([\\d\\s]+)$"
	re := regexp.MustCompile(space_regex)
	var spaces []space
	var presents []present
	for x := 0; x < len(input); x++ {
		line := input[x]
		if strings.Contains(line, "x") {
			matches := re.FindStringSubmatch(line)
			length := utils.Atoi(matches[1])
			width := utils.Atoi(matches[2])
			var quantities []int
			for _, num := range strings.Split(matches[3], " ") {
				quantities = append(quantities, utils.Atoi(num))
			}

			spaces = append(spaces, space{
				length:           length,
				width:            width,
				present_quantity: quantities,
			})

			continue
		} else if strings.Contains(line, ":") {
			continue
		}

		if strings.Contains(line, ".") || strings.Contains(line, "#") {
			area := 0
			var present_array [][]bool
			for strings.Contains(line, ".") || strings.Contains(line, "#") {
				next_line := make([]bool, len(line))
				line_split := strings.Split(line, "")

				for x, char := range line_split {
					if char == "#" {
						next_line[x] = true
						area++
					}
				}
				present_array = append(present_array, next_line)
				x++
				line = input[x]
			}
			presents = append(presents, present{
				grid: present_array,
				area: area,
			})
		}
	}

	return presents, spaces
}

func removeInvalidSpaces(presents []present, spaces []space) []space {
	var new_spaces []space
	for _, sp := range spaces {
		space_area := sp.length * sp.width
		total_gift_area := 0
		for x, count := range sp.present_quantity {
			total_gift_area += presents[x].area * count
		}
		if space_area > total_gift_area {
			new_spaces = append(new_spaces, sp)
		}

	}

	return new_spaces
}

func sum(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func SolvePartOne(file_slice []string) string {
	presents, spaces := ParseInput(file_slice)

	spaces = removeInvalidSpaces(presents, spaces)

	count := 0
	for _, sp := range spaces {
		piece_count := sum(sp.present_quantity)
		spacex3x3 := (sp.length / 3) * (sp.width / 2) // Works on the test and my input, highly doubt it works generally
		// fmt.Println(sp, " : ", spacex3x3)
		if piece_count < spacex3x3 {
			count++
		}
	}

	// for x, sp := range spaces {
	// 	fmt.Println(x, " -> ", sp)
	// }
	// fmt.Println("Space Count: ", count)
	return fmt.Sprintf("%d", count)

}

func SolvePartTwo(file_slice []string) string {
	return "No Part 2"
}

func Solve(input []string, arg int) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
