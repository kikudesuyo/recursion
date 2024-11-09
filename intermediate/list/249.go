package list

func IsMountain(height []int32) bool {
	if len(height) < 3 {
		return false
	}
	if height[0] >= height[1] {
		return false
	}
	isAscending := true
	for i := 0; i < len(height)-1; i++ {
		if height[i] == height[i+1] {
			return false
		}
		if isAscending && height[i] > height[i+1] {
			isAscending = false
			continue
		}
		if !isAscending && height[i] < height[i+1] {
			return false
		}
	}
	return !isAscending
}
