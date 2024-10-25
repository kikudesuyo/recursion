package recursion

func Pascal(x int32) int32 {
	var pascalNum int32
	for row := int32(1); row <= x; row++ {
		pascalNum += row
	}
	return pascalNum
}
