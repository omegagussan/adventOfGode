package common

import "math"

func isPrimeOptimized(n int) bool {
	if n <= 1 {
		return false
	} else if n == 2 {
		return true
	} else if n%2 == 0 {
		return false
	}
	sqrt := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrt; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func GetFirstNPrime(n int) []int {
	var primes []int
	var i = 2
	for len(primes) < n {
		if isPrimeOptimized(i) {
			primes = append(primes, i)
		}
		i += 1
	}
	return primes
}
