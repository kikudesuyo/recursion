package list

func PrimesUpToNCount(n int32) int32 {
	if n <= 2 {
		return 0
	}
	var primes []int
	for i := 2; i < int(n); i++ {
		flag := true
		for _, prime := range primes {
			if i%prime == 0 {
				flag = false
				break
			}
		}
		if flag {
			primes = append(primes, i)
		}
	}
	return int32(len(primes))
}
