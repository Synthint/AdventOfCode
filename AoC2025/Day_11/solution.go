package template

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"strings"
)

func deleteItem(input []string, item string) []string {
	for x, val := range input {
		if val == item {
			return append(input[:x], input[x+1:]...)
		}
	}
	return input
}

func parseInput(input []string) map[string][]string {
	connection_nodes := make(map[string][]string)
	connection_regex := "^(\\w+):\\s([\\w\\s]+)$"
	re := regexp.MustCompile(connection_regex)
	for _, line := range input {
		matches := re.FindStringSubmatch(line)
		connection_nodes[matches[1]] = strings.Split(matches[2], " ")
	}
	return connection_nodes
}

func cullTheOrphans(nodes map[string][]string, keep_svr bool) map[string][]string { // Think of a less dark name
	parents_and_potential_orphans := slices.Collect(maps.Keys(nodes))

	for _, children := range nodes {
		for _, child := range children {
			parents_and_potential_orphans = deleteItem(parents_and_potential_orphans, child) // remove parents that are also children
		}
	}

	nodes_with_parent := make(map[string][]string)
	maps.Copy(nodes_with_parent, nodes)
	for _, orphan := range parents_and_potential_orphans {
		if !(keep_svr && orphan == "svr") {
			delete(nodes_with_parent, orphan)
		}
	}
	return nodes_with_parent
}

func removeYouChildren(nodes map[string][]string) map[string][]string {
	for key, children := range nodes {
		for _, child := range children {
			if child == "you" {
				nodes[key] = deleteItem(nodes[key], "you")
				break
			}
		}
	}
	return nodes
}

func cullTheChildless(nodes map[string][]string) map[string][]string { // Welp if we'r already culling children
	var childless []string

	for k, item := range nodes {
		if len(item) == 0 {
			childless = append(childless, k)
		}
	}

	for _, key := range childless {
		delete(nodes, key)
	}
	return nodes
}

func traceNetwork(nodes map[string][]string, origin string) int {
	// fmt.Println(origin, " -> ", nodes[origin])
	if nodes[origin][0] == "out" {
		return 1
	}
	total := 0
	for _, child := range nodes[origin] {
		total += traceNetwork(nodes, child)
	}
	return total
}

func SolvePartOne(file_slice []string) string {
	connection_nodes := parseInput(file_slice)
	connection_nodes = cullTheOrphans(connection_nodes, false)
	connection_nodes = removeYouChildren(connection_nodes)
	connection_nodes = cullTheChildless(connection_nodes)

	return fmt.Sprintf("%d", traceNetwork(connection_nodes, "you"))
}

type memo struct {
	total int
	dac   bool
	fft   bool
}

type memoKey struct {
	node string
	dac  bool
	fft  bool
}

var memos map[memoKey]int

func traceNetworkWithTag(
	nodes map[string][]string,
	origin string,
	seenDAC bool,
	seenFFT bool,
) int {

	if origin == "dac" {
		seenDAC = true
	}
	if origin == "fft" {
		seenFFT = true
	}

	key := memoKey{origin, seenDAC, seenFFT}
	if v, ok := memos[key]; ok {
		return v
	}

	if nodes[origin][0] == "out" {
		if seenDAC && seenFFT {
			memos[key] = 1
			return 1
		}
		memos[key] = 0
		return 0
	}

	total := 0
	for _, child := range nodes[origin] {
		total += traceNetworkWithTag(nodes, child, seenDAC, seenFFT)
	}

	memos[key] = total
	return total
}

func SolvePartTwo(file_slice []string) string {
	memos = make(map[memoKey]int)

	connection_nodes := parseInput(file_slice)
	connection_nodes = cullTheOrphans(connection_nodes, true)
	connection_nodes = cullTheChildless(connection_nodes)

	total := traceNetworkWithTag(connection_nodes, "svr", false, false)
	return fmt.Sprintf("%d", total)
}

func Solve(input []string, arg int) (string, string) {
	TEST_SPLIT_TERM := "SPLIT"
	part_one_solution := ""
	part_two_solution := ""
	if slices.Contains(input, TEST_SPLIT_TERM) { //Annoyingly, the test input is different for part 1 and part 2 and the input from one breaks the other
		ind := slices.Index(input, TEST_SPLIT_TERM)
		part_one_solution = SolvePartOne(input[:ind])
		part_two_solution = SolvePartTwo(input[ind+1:])
	} else {
		part_one_solution = SolvePartOne(input)
		part_two_solution = SolvePartTwo(input)

	}
	return part_one_solution, part_two_solution
}
