package controlflow

import (
	"strconv"
	"strings"
)

func isPerfectNumber(n int32) bool {
	sum := 0
	for i := 1; i < int(n); i++ {
		if int(n)%i == 0 {
			sum += i
		}
	}
	return sum == int(n)
}

func PerfectNumberList(n int32) string {
	arr := make([]string, 0)
	for i := 1; i <= int(n); i++ {
		if isPerfectNumber(int32(i)) {
			arr = append(arr, strconv.Itoa(i))
		}
	}
	if len(arr) == 0 {
		return "none"
	}
	return strings.Join(arr, "-")
}
