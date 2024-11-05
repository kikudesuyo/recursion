package list

import "strings"

func IsFirstStringLarger(s1 string, s2 string) bool {
	var asciiS1Sum int
	lowerS1 := strings.ToLower(s1)
	for _, char := range lowerS1 {
		asciiS1Sum += int(char)
	}
	var asciiS2Sum int
	lowerS2 := strings.ToLower(s2)
	for _, char := range lowerS2 {
		asciiS2Sum += int(char)
	}
	if asciiS1Sum > asciiS2Sum {
		return true
	} else {
		return false
	}
}
