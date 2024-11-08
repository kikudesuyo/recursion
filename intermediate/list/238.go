package list

import (
	"strings"
)

var letters = []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

func GenerateAlphabet(youngAlphabet byte, oldAlphabet byte) []byte {

	var result []byte
	if strings.ToLower(string(youngAlphabet)) > strings.ToLower(string(oldAlphabet)) {
		youngAlphabet, oldAlphabet = oldAlphabet, youngAlphabet
	}
	firstIndex := strings.Index(string(letters), strings.ToLower(string(youngAlphabet)))
	secondIndex := strings.Index(string(letters), strings.ToLower(string(oldAlphabet)))
	for i := firstIndex; i <= secondIndex; i++ {
		result = append(result, letters[i])
	}
	return result
}
