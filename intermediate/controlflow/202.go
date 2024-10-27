package controlflow

import (
	"reflect"
	"strconv"
)

func reverse(arr []string) []string {
	var reversed []string
	for i := len(arr) - 1; i >= 0; i-- {
		reversed = append(reversed, arr[i])
	}
	return reversed
}

func IsPalindromeInteger(n int32) bool {
	strN := strconv.Itoa(int(n))
	arr := []string{}
	for _, v := range strN {
		arr = append(arr, string(v))
	}
	reversed := reverse(arr)
	if reflect.DeepEqual(arr, reversed) {
		return true
	} else {
		return false
	}
}
