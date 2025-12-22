package template

import (
	"AoC2025/utils"
	"fmt"
	"maps"
	"math"
	"regexp"
	"slices"
	"strings"

	"github.com/lanl/clp"
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

func DetermineCoefficients(mach machine) [][]int {
	matrix := make([][]int, len(mach.joltage_reqs))
	matrix_width := len(mach.buttons) + 1

	for x, req := range mach.joltage_reqs {
		next_line := make([]int, matrix_width)
		next_line[matrix_width-1] = req
		for ind, button := range mach.buttons {
			if slices.Contains(button, x) {
				next_line[ind] = 1
			}
		}
		matrix[x] = next_line
	}

	return matrix
}

func PackCLPMatrix(coeff_matrix [][]int) *clp.PackedMatrix { // My coeff matrix wasn't good enough I guess?
	num_rows := len(coeff_matrix)
	num_vars := len(coeff_matrix[0]) - 1

	packed_matrix := clp.NewPackedMatrix()

	for x := 0; x < num_vars; x++ {
		column := []clp.Nonzero{}

		for y := 0; y < num_rows; y++ {
			coef := coeff_matrix[y][x]
			if coef != 0 {
				column = append(column, clp.Nonzero{
					Index: y,
					Value: float64(coef),
				})
			}
		}

		packed_matrix.AppendColumn(column)
	}

	return packed_matrix
}

func SolveCLP(coeff_matrix [][]int, extra_var_bounds map[int]clp.Bounds) ([]float64, float64, error) {

	num_rows := len(coeff_matrix)
	num_vars := len(coeff_matrix[0]) - 1

	matrix := PackCLPMatrix(coeff_matrix)

	var_bounds := make([]clp.Bounds, num_vars)
	for index := 0; index < num_vars; index++ {
		if bound, ok := extra_var_bounds[index]; ok {
			var_bounds[index] = bound
		} else {
			var_bounds[index] = clp.Bounds{
				Lower: 0,
				Upper: math.Inf(1),
			}
		}
	}

	row_bounds := make([]clp.Bounds, num_rows)
	for row_index := 0; row_index < num_rows; row_index++ {
		joltage_answer := float64(coeff_matrix[row_index][num_vars])
		row_bounds[row_index] = clp.Bounds{
			Lower: joltage_answer,
			Upper: joltage_answer,
		}
	}

	objective := make([]float64, num_vars)
	for i := range objective {
		objective[i] = 1
	}

	lp := clp.NewSimplex()
	lp.LoadProblem(matrix, var_bounds, objective, row_bounds, nil)
	lp.SetOptimizationDirection(clp.Minimize)

	if lp.Primal(clp.NoValuesPass, clp.NoStartFinishOptions) != clp.Optimal {
		return nil, 0, fmt.Errorf("CLP err")
		// Still gives an answer without this, and a lower answer too but that one is wrong ¯\_(ツ)_/¯
	}

	solution := lp.PrimalColumnSolution()
	objective_value := lp.ObjectiveValue()

	return solution, objective_value, nil
}

//

func IsPracticallyAnInteger(num float64) bool {
	return math.Abs(num-math.Round(num)) > 1e-6
}

func BranchAndBound(matrix [][]int, current_bounds map[int]clp.Bounds, best_solution *int) {
	solution, lp_value, err := SolveCLP(matrix, current_bounds)
	if err != nil {
		return
	}
	if int(math.Ceil(lp_value)) >= *best_solution {
		return
	}

	for var_index, value := range solution {
		if IsPracticallyAnInteger(value) {
			floor_val := math.Floor(value)
			ceil_val := math.Ceil(value)

			left_bounds := maps.Clone(current_bounds)
			left_bounds[var_index] = clp.Bounds{
				Lower: 0,
				Upper: floor_val,
			}
			BranchAndBound(matrix, left_bounds, best_solution)

			right_bounds := maps.Clone(current_bounds)
			right_bounds[var_index] = clp.Bounds{
				Lower: ceil_val,
				Upper: math.Inf(1),
			}
			BranchAndBound(matrix, right_bounds, best_solution)

			return
		}
	}

	// All variables are integer — update best solution
	sum := 0
	for _, v := range solution {
		sum += int(math.Round(v))
	}

	if sum < *best_solution {
		*best_solution = sum
	}

}

func SolveMatrix(coeff_matrix [][]int) (int, error) {
	best_solution := math.MaxInt32

	BranchAndBound(
		coeff_matrix,
		map[int]clp.Bounds{},
		&best_solution,
	)

	if best_solution == math.MaxInt32 {
		return 0, fmt.Errorf("no solution")
	}

	return best_solution, nil
}

func SolvePartTwo(file_slice []string) string {
	machines := parse_machine_input(file_slice)
	total := 0

	for _, mach := range machines {
		// fmt.Println("---")
		coeffMatrix := DetermineCoefficients(mach)
		value, _ := SolveMatrix(coeffMatrix)
		total += value
		// fmt.Println(total)
	}

	return fmt.Sprintf("%d", total)
}

func Solve(input []string, arg int) (string, string) {

	partOneChan := make(chan string)
	partTwoChan := make(chan string)

	go func() {
		result := SolvePartOne(input)
		partOneChan <- result
	}()

	go func() {
		result := SolvePartTwo(input)
		partTwoChan <- result
	}()

	part_one_solution := <-partOneChan
	part_two_solution := <-partTwoChan

	return part_one_solution, part_two_solution
}
