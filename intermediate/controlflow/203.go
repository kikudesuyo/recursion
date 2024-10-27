package controlflow

func isPrime(n int32) bool {
	if n < 2 {
		return false
	}
	for i := int32(2); i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func SumOfAllPrimes(n int32) int32 {
	var sum int32
	for i := int32(2); i <= n; i++ {
		if isPrime(i) {
			sum += i
		}
	}
	return sum
}
