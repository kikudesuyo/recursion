package list

func maxIndex(intList []int) int {
	maxIndex := 0
	for i, n := range intList {
		if n > intList[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func MaxAsciiString(stringList []string) int32 {
	var asciiStrings []int
	for _, str := range stringList {
		n := 0
		for _, char := range str {
			n += int(char)
		}
		asciiStrings = append(asciiStrings, n)
	}
	return int32(maxIndex(asciiStrings))

}
