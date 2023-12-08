package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func SolvePartOne(file_slice []string) int {
	var rpsMap = make(map[string]int)
	// A and X are rock Adds 1
	// B and Y are paper Adds 2
	// C and Z are Scissors Adds 3
	// Win Adds 6, Tie Adds 3, Lose Adds 0
	var Rock int = 1
	var Paper int = 2
	var Scissors int = 3
	var Win int = 6
	var Tie int = 3
	var Lose int = 0
	rpsMap["A X"] = Rock + Tie
	rpsMap["A Y"] = Paper + Win
	rpsMap["A Z"] = Scissors + Lose
	rpsMap["B X"] = Rock + Lose
	rpsMap["B Y"] = Paper + Tie
	rpsMap["B Z"] = Scissors + Win
	rpsMap["C X"] = Rock + Win
	rpsMap["C Y"] = Paper + Lose
	rpsMap["C Z"] = Scissors + Tie

	var score = 0
	for _, line := range file_slice {
		line = strings.TrimSpace(line)
		score += rpsMap[line]
	}
	return score
}

func SolvePartTwo(file_slice []string) int {
	var rpsMap = make(map[string]int)
	// X means Win -> 6
	// Y means Tie -> 3
	// Z means Lose -> 0
	var Rock int = 1
	var Paper int = 2
	var Scissors int = 3
	var Win int = 6
	var Tie int = 3
	var Lose int = 0
	rpsMap["A X"] = Lose + Scissors
	rpsMap["A Y"] = Tie + Rock
	rpsMap["A Z"] = Win + Paper
	rpsMap["B X"] = Lose + Rock
	rpsMap["B Y"] = Tie + Paper
	rpsMap["B Z"] = Win + Scissors
	rpsMap["C X"] = Lose + Paper
	rpsMap["C Y"] = Tie + Scissors
	rpsMap["C Z"] = Win + Rock

	var score = 0
	for _, line := range file_slice {
		line = strings.TrimSpace(line)
		score += rpsMap[line]
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
