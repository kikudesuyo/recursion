package list

import (
	"sort"
)

func min(a int32, b int32) int32 {
	if a >= b {
		return b
	} else {
		return a
	}
}

func IntersectionOfArraysRepeats(intList1 []int32, intList2 []int32) []int32 {
	map1 := map[int32]int32{}
	for _, v := range intList1 {
		map1[v]++
	}
	map2 := map[int32]int32{}
	for _, v := range intList2 {
		map2[v]++
	}
	arr := make([]int32, 0)
	for key, map1Val := range map1 {
		if _, exists := map2[key]; exists {
			cnt := int(min(map1Val, map2[key]))
			for i := 0; i < cnt; i++ {
				arr = append(arr, key)
			}
		}
	}
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })
	return arr
}
