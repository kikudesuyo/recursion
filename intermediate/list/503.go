package list

import "sort"

func FindPairs(numbers []int32) []int32 {
	mapNumbers := make(map[int32]int)
	for _, number := range numbers {
		mapNumbers[number]++
	}
	pairs := make([]int32, 0)
	for key, value := range mapNumbers {
		if value == 2 {
			pairs = append(pairs, key)
		}
	}
	sort.Slice(pairs, func(i, j int) bool { return pairs[i] < pairs[j] })
	return pairs
}
