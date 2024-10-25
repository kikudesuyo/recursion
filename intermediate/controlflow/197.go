package controlflow

func IsPrime(number int32) bool {
	if number < 2 {
		return false
	}
	for i := 2; i < int(number); i++ {
		if int(number)%i == 0 {
			return false
		}
	}
	return true
}
