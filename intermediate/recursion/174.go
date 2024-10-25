package recursion

import "strconv"

func Divisor(number int32) string {
	var divisors string
	for i := int(number); i > 0; i-- {
		if i == 1 {
			divisors += strconv.Itoa(int(number) / i)
		} else if int(number)%i == 0 {
			divisors += strconv.Itoa(int(number)/i) + "-"
		}
	}
	return divisors
}
