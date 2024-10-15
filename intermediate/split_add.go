package intermediate

import "strconv"

func SplitAndAdd(digits int32) int32 {
	stringDigits := strconv.Itoa(int(digits))
	digitSum := int32(0)
	for _, digit := range stringDigits {
		digitSum += int32(digit - '0')
	}
	return digitSum
}
