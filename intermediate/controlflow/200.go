package controlflow

import (
	"strconv"
	"strings"
)

func FizzBuzz(n int32) string {
	arr := make([]string, 0)
	for i := 1; i <= int(n); i++ {
		if i%15 == 0 {
			arr = append(arr, "FizzBuzz")
		} else if i%3 == 0 {
			arr = append(arr, "Fizz")
		} else if i%5 == 0 {
			arr = append(arr, "Buzz")
		} else {
			arr = append(arr, strconv.Itoa(i))
		}
	}
	return strings.Join(arr, "-")
}
