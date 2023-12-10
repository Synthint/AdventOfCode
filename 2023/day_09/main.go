package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func find_sub_history(nums []int) []int {
	var sub_nums []int
	for i := 0; i < len(nums)-1; i++ {
		sub_nums = append(sub_nums, nums[i+1]-nums[i])
	}
	return sub_nums
}

func recur_predictions(nums []int) []int {
	var sub_nums = find_sub_history(nums)
	var all_zero = true
	var nums_future []int = make([]int, len(nums))
	copy(nums_future, nums)
	for _, x := range sub_nums {
		if x != 0 {
			all_zero = false
			break
		}
	}
	if all_zero {
		nums_future = append(nums_future, nums[len(nums)-1])
	} else {
		sub_nums = recur_predictions(sub_nums)
		nums_future = append(nums_future, nums[len(nums)-1]+sub_nums[len(sub_nums)-1])
	}
	return nums_future
}

func SolvePartOne(file_slice []string) int {
	var new_total = 0
	for _, line := range file_slice {
		var nums_as_strings = strings.Split(line, " ")
		nums := make([]int, len(nums_as_strings))

		for i, s := range nums_as_strings {
			nums[i], _ = strconv.Atoi(s)
		}

		var new_nums = recur_predictions(nums)
		new_total += new_nums[len(new_nums)-1]
	}
	return new_total
}

func recur_history(nums []int) []int {
	var sub_nums = find_sub_history(nums)
	var all_zero = true
	var nums_future []int = make([]int, len(nums))
	copy(nums_future, nums)
	for _, x := range sub_nums {
		if x != 0 {
			all_zero = false
			break
		}
	}
	if all_zero {
		nums_future = append([]int{nums[0]}, nums_future...)
	} else {
		sub_nums = recur_history(sub_nums)
		nums_future = append([]int{nums[0] - sub_nums[0]}, nums_future...)
	}
	return nums_future
}

func SolvePartTwo(file_slice []string) int {
	var new_total = 0
	for _, line := range file_slice {
		var nums_as_strings = strings.Split(line, " ")
		nums := make([]int, len(nums_as_strings))

		for i, s := range nums_as_strings {
			nums[i], _ = strconv.Atoi(s)
		}

		var new_nums = recur_history(nums)
		new_total += new_nums[0]

	}
	return new_total
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
