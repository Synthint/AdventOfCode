package template

import (
	"AoC2025/utils"
	"fmt"
	"maps"
	"math"
	"slices"
	"sort"
	"strings"
)

type junctionBox struct {
	x, y, z int
}

func (jb *junctionBox) toString() string {
	return fmt.Sprintf("%d,%d,%d", jb.x, jb.y, jb.z)
}

func (jb1 *junctionBox) distanceTo(jb2 junctionBox) float64 {
	return math.Sqrt(math.Pow(float64(jb1.x-jb2.x), 2) + math.Pow(float64(jb1.y-jb2.y), 2) + math.Pow(float64(jb1.z-jb2.z), 2))
}

func StrToJunc(s string) junctionBox {
	nums := strings.Split(s, ",")
	return junctionBox{
		x: utils.Atoi(nums[0]),
		y: utils.Atoi(nums[1]),
		z: utils.Atoi(nums[2]),
	}
}

func parseStrings(s []string) []junctionBox {
	jbs := make([]junctionBox, len(s))
	for x, item := range s {
		jbs[x] = StrToJunc(item)
	}
	return jbs
}

func key(jb1, jb2 junctionBox) string {
	return fmt.Sprintf("%s-%s", jb1.toString(), jb2.toString())
}

func decodeKey(str string) (junctionBox, junctionBox) {
	juncs := strings.Split(str, "-")
	return StrToJunc(juncs[0]), StrToJunc(juncs[1])
}

func genDistanceMap(jbs []junctionBox) map[string]float64 {
	distMap := make(map[string]float64)

	for x := 0; x < len(jbs); x++ {
		for y := x + 1; y < len(jbs); y++ {
			distMap[key(jbs[x], jbs[y])] = jbs[x].distanceTo(jbs[y])
		}
	}
	return distMap
}

func SolvePartOne(file_slice []string, connections_to_count int) string {
	junctions := parseStrings(file_slice)

	distMap := genDistanceMap(junctions)
	mapKeys := slices.Collect(maps.Keys(distMap))

	sort.Slice(mapKeys, func(i, j int) bool {
		return distMap[mapKeys[i]] < distMap[mapKeys[j]]
	})

	connections := 0
	var circuits [][]junctionBox
	for _, key := range mapKeys {
		jb1, jb2 := decodeKey(key)

		contains_1, contains_2 := -1, -1
		for x, circ := range circuits {
			if slices.Contains(circ, jb1) {
				contains_1 = x
			}
			if slices.Contains(circ, jb2) {
				contains_2 = x
			}
			if contains_1 != -1 && contains_2 != -1 {
				break
			}
		}

		if contains_1 == contains_2 {
			if contains_1 == -1 { // Neither Circuit exists
				circuits = append(circuits, []junctionBox{
					jb1, jb2,
				})
			}
			// Do nothing if already in the same circuit and not missing from circuits
		} else {
			if contains_1 != -1 && contains_2 != -1 { // Join 2 existing circuits
				circuits[contains_1] = append(circuits[contains_1], circuits[contains_2]...)
				circuits = slices.Delete(circuits, contains_2, contains_2+1)
			} else if contains_1 != -1 { // Add missing connection to existing circuit
				// 2 missing
				circuits[contains_1] = append(circuits[contains_1], jb2)
			} else if contains_2 != -1 {
				// 1 missing
				circuits[contains_2] = append(circuits[contains_2], jb1)
			}
		}

		connections += 1
		if connections == connections_to_count {
			break
		}
	}

	sort.Slice(circuits, func(i, j int) bool {
		return len(circuits[i]) > len(circuits[j])
	})

	// for _, line := range circuits {
	// 	fmt.Println(line)
	// }

	return fmt.Sprintf("%d", len(circuits[0])*len(circuits[1])*len(circuits[2]))
}

func SolvePartTwo(file_slice []string) string {
	junctions := parseStrings(file_slice)

	distMap := genDistanceMap(junctions)
	mapKeys := slices.Collect(maps.Keys(distMap))

	sort.Slice(mapKeys, func(i, j int) bool {
		return distMap[mapKeys[i]] < distMap[mapKeys[j]]
	})

	var circuits [][]junctionBox
	var res1, res2 junctionBox
	for _, key := range mapKeys {
		jb1, jb2 := decodeKey(key)

		contains_1, contains_2 := -1, -1
		for x, circ := range circuits {
			if slices.Contains(circ, jb1) {
				contains_1 = x
			}
			if slices.Contains(circ, jb2) {
				contains_2 = x
			}
			if contains_1 != -1 && contains_2 != -1 {
				break
			}
		}

		if contains_1 == contains_2 {
			if contains_1 == -1 { // Neither Circuit exists
				circuits = append(circuits, []junctionBox{
					jb1, jb2,
				})
			}
			// Do nothing if already in the same circuit and not missing from circuits
		} else {
			if contains_1 != -1 && contains_2 != -1 { // Join 2 existing circuits
				circuits[contains_1] = append(circuits[contains_1], circuits[contains_2]...)
				circuits = slices.Delete(circuits, contains_2, contains_2+1)
			} else if contains_1 != -1 { // Add missing connection to existing circuit
				// 2 missing
				circuits[contains_1] = append(circuits[contains_1], jb2)
			} else if contains_2 != -1 {
				// 1 missing
				circuits[contains_2] = append(circuits[contains_2], jb1)
			}
		}

		if len(circuits) == 1 && len(circuits[0]) == len(junctions) {
			res1 = jb1
			res2 = jb2
			break
		}
	}

	// for _, line := range circuits {
	// 	fmt.Println(line)
	// }

	return fmt.Sprintf("%d", res1.x*res2.x)
}

func Solve(input []string, arg int) (string, string) {
	part_one_solution := SolvePartOne(input, arg)
	part_two_solution := SolvePartTwo(input)
	return part_one_solution, part_two_solution
}
