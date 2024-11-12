package list

import "strings"

func IsPangram(str string) bool {
	lowerStr := strings.ToLower(str)
	isExistAlphabet := [26]bool{}
	for _, c := range lowerStr {
		if 'a' <= c && c <= 'z' {
			isExistAlphabet[c-'a'] = true
		}
	}
	for _, exist := range isExistAlphabet {
		if !exist {
			return false
		}
	}
	return true
}
