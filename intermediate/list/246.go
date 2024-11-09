package list

func RotateByTimes(ids []int32, n int32) []int32 {
	rotatedIds := make([]int32, len(ids))
	for i, id := range ids {
		rotatedIds[(i+int(n))%len(ids)] = id
	}
	return rotatedIds
}
