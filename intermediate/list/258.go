package list

func ShuffleSuccessRate(arr []int32, shuffledArr []int32) int32 {
	var cnt int
	for i, elem := range arr {
		for j, shuffledElem := range shuffledArr {
			if elem == shuffledElem && i != j {
				cnt++
				break
			}
		}
	}
	return int32((cnt * 100 / len(arr)))
}
