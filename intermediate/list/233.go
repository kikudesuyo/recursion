package list

func AddEveryOtherElement(intArr []int32) int32 {
	var oddSum int32
	for i, v := range intArr {
		if i%2 == 0 {
			oddSum += v
		}
	}
	return oddSum
}
