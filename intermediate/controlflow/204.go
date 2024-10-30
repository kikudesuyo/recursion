package controlflow

import (
	"fmt"
)

func DecimalToBinary(n int32) string {
	s := ""
	s = fmt.Sprintf("%b", n)
	return s
}
