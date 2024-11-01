package controlflow

import "strings"

func isEqual(arr []string, val string) bool {
	for _, element := range arr {
		if element != val {
			return false
		}
	}
	return true
}

func TwosComplement(bits string) string {
	var complement []string
	for _, bit := range bits {
		if bit == '0' {
			complement = append(complement, "1")
		} else {
			complement = append(complement, "0")
		}
	}
	if isEqual(complement, "1") {
		resutts := []string{"1"}
		for i := 0; i < len(complement); i++ {
			resutts = append(resutts, "0")
		}
		return strings.Join(resutts, "")
	}
	for i := len(complement) - 1; i >= 0; i-- {
		if complement[i] == "1" {
			complement[i] = "0"
		} else {
			complement[i] = "1"
			break
		}
	}
	return strings.Join(complement, "")
}
