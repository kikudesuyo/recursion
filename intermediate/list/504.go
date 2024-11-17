package list

func FirstNonRepeating(s string) int32 {
	problemMap := make(map[string]int32)
	for _, r := range s {
		problemMap[string(r)]++
	}
	for i, r := range s {
		if problemMap[string(r)] == 1 {
			return int32(i)
		}
	}
	return -1
}
