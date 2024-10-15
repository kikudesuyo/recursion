package intermediate

func DivideBy3Count(n int32) int32 {
	if n == 0 {
		return 0
	}
	if n%3 == 0 {
		return 1 + DivideBy3Count(n/3)
	}
	return DivideBy3Count(n / 3)
}
