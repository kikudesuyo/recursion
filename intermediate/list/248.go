package list

func HasPenalty(records []int32) bool {
	penalty := false
	for i := 0; i < len(records)-1; i++ {
		if records[i] > records[i+1] {
			penalty = true
			break
		}
	}
	return penalty
}
