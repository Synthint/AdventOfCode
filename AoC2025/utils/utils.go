package utils

import "strconv"

func Atoi(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return num
}

func StringsToInts(input []string) ([]int, error) {
	result := make([]int, 0, len(input))

	for _, s := range input {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		result = append(result, n)
	}

	return result, nil
}
