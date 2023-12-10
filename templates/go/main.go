package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func SolvePartOne(file_slice []string) int {
	var score = 0
	for i, _ := range file_slice {
		score += i
	}
	return score
}

func SolvePartTwo(file_slice []string) int {
	var score = 0
	for i, _ := range file_slice {
		score += i
	}
	return score
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var file_slice []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		file_slice = append(file_slice, scanner.Text())
	}

	fmt.Println("Part 1 Answer -> ", SolvePartOne(file_slice))
	fmt.Println("Part 2 Answer -> ", SolvePartTwo(file_slice))
}
