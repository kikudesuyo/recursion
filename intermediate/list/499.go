package list

func MissingItems(listA []int32, listB []int32) []int32 {
	m := map[int32]bool{}
	for _, v := range listB {
		m[v] = true
	}
	arr := make([]int32, 0)
	for _, v := range listA {
		if _, exists := m[v]; !exists {
			arr = append(arr, v)
		}
	}
	if len(arr) == 0 {
		return listA
	}
	return arr
}
