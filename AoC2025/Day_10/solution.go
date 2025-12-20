package template

import (
	"AoC2025/utils"
	"fmt"

	// "math"
	"regexp"
	"slices"
	"strings"
	// "github.com/lanl/clp"
)

type machine struct {
	desired_indicator []bool
	buttons           [][]int
	joltage_reqs      []int
}

func parse_machine_input(input []string) []machine {
	machine_regex := "^\\[([.#]+)\\]\\s(\\([\\d\\(\\),\\s]+\\))\\s{([\\d,]+)}$"
	re := regexp.MustCompile(machine_regex)

	var machines []machine
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		// fmt.Println(line)
		// for x, match := range matches {
		// 	fmt.Println("\t", x, ": ", match)
		// }
		if len(matches) == 0 {
			continue
		}
		// matches[0] is full
		indicator_str := matches[1]
		button_strs := strings.Split(matches[2], " ")
		jolt_reqs_str := strings.Split(matches[3], ",")

		des_ind := make([]bool, len(indicator_str))

		for x, char := range strings.Split(indicator_str, "") {
			if char == "#" {
				des_ind[x] = true
			} else {
				des_ind[x] = false
			}
		}

		jolt_reqs := make([]int, len(jolt_reqs_str))
		for x, num_str := range jolt_reqs_str {
			jolt_reqs[x] = utils.Atoi(num_str)
		}

		button_nums := make([][]int, len(button_strs))
		for x, str := range button_strs {
			nums := strings.Split(strings.ReplaceAll(strings.ReplaceAll(str, "(", ""), ")", ""), ",")
			button_nums[x] = make([]int, len(nums))
			for y, num := range nums {
				button_nums[x][y] = utils.Atoi(num)
			}
		}

		machines = append(machines, machine{
			desired_indicator: des_ind,
			joltage_reqs:      jolt_reqs,
			buttons:           button_nums,
		})
	}
	return machines
}

func branchMachinePathsIndicators(start []bool, mach machine) [][]bool {
	if len(start) != len(mach.desired_indicator) {
		fmt.Println("Something not right")
	}

	output := make([][]bool, len(mach.buttons))
	for x, button := range mach.buttons {
		output[x] = make([]bool, len(start))
		copy(output[x], start)
		for _, num := range button {
			output[x][num] = !output[x][num]
		}
	}
	return output
}

func fewestToGetIndicators(mach machine) int {
	type queue_block struct {
		items [][]bool
		depth int
	}
	start_block := queue_block{
		items: [][]bool{
			make([]bool, len(mach.desired_indicator)),
		},
		depth: 0,
	}

	queue := []queue_block{start_block}

	for len(queue) != 0 {
		current_depth := queue[0].depth

		for _, block_input := range queue[0].items {
			next_block := queue_block{
				depth: current_depth + 1,
			}
			if slices.Equal(block_input, mach.desired_indicator) {
				return current_depth
			}
			next_block.items = branchMachinePathsIndicators(block_input, mach)

			queue = append(queue, next_block)
		}
		queue = slices.Delete(queue, 0, 1)
	}
	return -1
}

func SolvePartOne(file_slice []string) string {
	machines := parse_machine_input(file_slice)
	total := 0
	for _, machine := range machines {
		total += fewestToGetIndicators(machine)
	}
	// fmt.Println("Day 10 Part 1: ", total)
	return fmt.Sprintf("%d", total)

}

// func addButtonToJoltages(button []int, joltage []int) []int {
// 	for _, x := range button {
// 		if x > len(joltage) {
// 			panic("No")
// 		}
// 		joltage[x]++
// 	}
// 	return joltage
// }

// func bKey(button []int) string {
// 	return fmt.Sprint(button)
// }

// func fewestToGetJoltages(mach machine) int {
// 	var memo map[string][]int

// 	return 2
// }

func SolvePartTwo(file_slice []string) string {
	// machines := parse_machine_input(file_slice)
	total := 0

	// for _, machine := range machines {
	// 	total += fewestToGetJoltages(machine)
	// }

	return fmt.Sprintf("%d", total)
}

func Solve(input []string, arg int) (string, string) {
	// Channels to receive solutions from goroutines
	partOneChan := make(chan string)
	partTwoChan := make(chan string)

	// Start goroutine for SolvePartOne
	go func() {
		result := SolvePartOne(input)
		partOneChan <- result // Send result to channel
	}()

	// Start goroutine for SolvePartTwo
	go func() {
		result := SolvePartTwo(input)
		partTwoChan <- result // Send result to channel
	}()

	// Wait for both results to be received from the channels
	part_one_solution := <-partOneChan
	part_two_solution := <-partTwoChan

	// Return both solutions
	return part_one_solution, part_two_solution
}
