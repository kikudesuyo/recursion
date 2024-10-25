package recursion

func TotalSquareArea(x int32) int32 {
	var totalArea int32
	for row := int32(1); row <= x; row++ {
		totalArea += row * row * row
	}
	return totalArea

}
