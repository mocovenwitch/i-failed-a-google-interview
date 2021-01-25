package questions

import "math"

// This solution is too SLOW
func countPrimes(n int) int {
	// edge case
	if n <= 2 {
		return 0
	}

	if n == 3 {
		return 1
	}

	count := 2
	primes := []int{2, 3}
	// loop and count
	for i := 4; i < n; i++ {
		if isPrimes(i, primes) {
			primes = append(primes, i)
			count++
		}
	}

	return count

}

func isPrimes(n int, primes []int) bool {
	for _, p := range primes {
		if math.Mod(float64(n), float64(p)) == 0 {
			return false
		}
	}
	return true
}
