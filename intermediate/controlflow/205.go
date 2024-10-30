package controlflow

import "fmt"

func DecimalToHexadecimal(n int32) string {
	s := fmt.Sprintf("%X", n)
	return s
}
