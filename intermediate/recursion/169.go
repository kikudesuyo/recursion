package recursion

import (
	"strconv"
	"strings"
)

func Sheeps(count int32) string {
	var builder strings.Builder
	for i := int32(1); i <= count; i++ {
		builder.WriteString(strconv.Itoa(int(i)))
		builder.WriteString(" sheep ~ ")
	}
	return builder.String()
}
