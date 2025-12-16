package template

import "fmt"

func SolvePartOne(file_slice []string) string {
	width := len(file_slice[0]) + 2
	height := len(file_slice) + 2

	paperRolls := make([][]rune, height)
	for i := range paperRolls {
		paperRolls[i] = make([]rune, width)
		for j := range paperRolls[i] {
			paperRolls[i][j] = '.'
		}
	}

	for i, line := range file_slice {
		for j, char := range line {
			paperRolls[i+1][j+1] = char
		}
	}

	neighborMatrix := createNeighborMatrix(paperRolls)
	count := 0
	for h := 0; h < len(neighborMatrix); h++ {
		for w := 0; w < len(neighborMatrix[0]); w++ {
			// fmt.Printf("%d", neighborMatrix[h][w])
			if neighborMatrix[h][w] < 4 && paperRolls[h+1][w+1] == '@' {
				count++
				paperRolls[h+1][w+1] = 'X'
			}
		}
		// fmt.Printf("\n")
	}

	// fmt.Printf("Final Matrix:\n")
	// for h := 1; h < height-1; h++ {
	// 	for w := 1; w < width-1; w++ {
	// 		fmt.Printf("%c", paperRolls[h][w])
	// 	}
	// 	fmt.Printf("\n")
	// }

	return fmt.Sprintf("%d", count)
}

func createNeighborMatrix(matrix [][]rune) [][]int {
	height := len(matrix) - 2
	width := len(matrix[0]) - 2 // Assume a border of 1 around the input matrix
	neighborMatrix := make([][]int, height)
	for i := range neighborMatrix {
		neighborMatrix[i] = make([]int, width)
	}
	for nh := 1; nh <= height; nh++ {
		for nw := 1; nw <= width; nw++ {
			if matrix[nh][nw] != '@' {
				continue
			}
			count := 0
			for dh := -1; dh <= 1; dh++ {
				for dw := -1; dw <= 1; dw++ {
					if dh == 0 && dw == 0 {
						continue
					}
					if matrix[nh+dh][nw+dw] == '@' {
						count++
					}
				}
			}
			neighborMatrix[nh-1][nw-1] = count
		}
	}
	return neighborMatrix

}

func removeRolls(paperRolls [][]rune) (int, [][]rune) {
	neighborMatrix := createNeighborMatrix(paperRolls)
	count := 0
	for h := 0; h < len(neighborMatrix); h++ {
		for w := 0; w < len(neighborMatrix[0]); w++ {
			// fmt.Printf("%d", neighborMatrix[h][w])
			if neighborMatrix[h][w] < 4 && paperRolls[h+1][w+1] == '@' {
				count++
				paperRolls[h+1][w+1] = '.'
			}
		}
		// fmt.Printf("\n")
	}
	return count, paperRolls
}

func SolvePartTwo(file_slice []string) string {
	totalRemoved := 0
	width := len(file_slice[0]) + 2
	height := len(file_slice) + 2

	paperRolls := make([][]rune, height)
	for i := range paperRolls {
		paperRolls[i] = make([]rune, width)
		for j := range paperRolls[i] {
			paperRolls[i][j] = '.'
		}
	}

	for i, line := range file_slice {
		for j, char := range line {
			paperRolls[i+1][j+1] = char
		}
	}

	lastRemoved := -1
	for lastRemoved != 0 {
		lastRemoved, paperRolls = removeRolls(paperRolls)
		totalRemoved += lastRemoved
	}

	return fmt.Sprintf("%d", totalRemoved)

}

func Solve(input []string) (string, string) {
	part_one_solution := SolvePartOne(input)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
