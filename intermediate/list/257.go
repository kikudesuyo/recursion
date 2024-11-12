package list

func ShufflePositions(arr []int32, shuffledArr []int32) []int32 {
	var idxes []int32
	for _, elem := range arr {
		for i, shuffledElem := range shuffledArr {
			if elem == shuffledElem {
				idxes = append(idxes, int32(i))
			}
		}
	}
	return idxes
}
