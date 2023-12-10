package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func AssertEQ(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got %d; want %d", got, want)
	}
}

func ReadFile(file_str string) []string {
	file, err := os.Open(file_str)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var file_slice []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		file_slice = append(file_slice, scanner.Text())
	}
	return file_slice
}

func TestSolvePartOne(t *testing.T) {
	var file_slice = ReadFile("./test_input.txt")
	AssertEQ(SolvePartOne(file_slice), 114, t)
}

func TestSolvePartTwo(t *testing.T) {

	var file_slice = ReadFile("./test_input.txt")
	AssertEQ(SolvePartTwo(file_slice), 2, t)
}
