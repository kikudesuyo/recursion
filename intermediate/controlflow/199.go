package controlflow

import (
	"strconv"
	"strings"
)

func NotDivisibleNumbers(x int32, y int32) string {
	var arr []string
	for i := 1; i <= int(x); i++ {
		if i%int(y) != 0 {
			arr = append(arr, strconv.Itoa(i))
		}
	}
	return strings.Join(arr, "-")
}
