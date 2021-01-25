package questions

// This solution is too SLOW
func countPrimes1(n int) int {
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
		// if math.Mod(float64(n), float64(p)) == 0 {
		if n%p == 0 {
			return false
		}
	}
	return true
}

// Runtime: 8 ms, faster than 86.88% of Go online submissions for Count Primes.
// Memory Usage: 4.9 MB, less than 33.33% of Go online submissions for Count Primes.
func countPrimes2(n int) int {
	primes := make([]bool, n)
	// set default value
	for i := 0; i < n; i++ {
		primes[i] = true
	}
	// set false to non primes
	for i := 2; i*i < n; i++ {
		if primes[i] {
			for j := i; j*i < n; j++ {
				primes[i*j] = false
			}
		}
	}
	// then count them out
	count := 0
	for i := 2; i < n; i++ {
		if primes[i] {
			count++
		}
	}
	return count
}

// Runtime: 4 ms, faster than 98.94% of Go online submissions for Count Primes.
// Memory Usage: 4.9 MB, less than 33.33% of Go online submissions for Count Primes.
// Compare to the 3rd approach, it save the step that set the default value to true, and
// then pick the opposite flag in notprimes array
func countPrimes3(n int) int {
	notprimes := make([]bool, n)

	// set false to non primes
	for i := 2; i*i < n; i++ {
		if !notprimes[i] {
			for j := i; j*i < n; j++ {
				notprimes[i*j] = true
			}
		}
	}
	// then count them out
	count := 0
	for i := 2; i < n; i++ {
		if !notprimes[i] {
			count++
		}
	}
	return count
}

// Runtime: 1632 ms, faster than 7.09% of Go online submissions for Count Primes.
// Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Count Primes.
func countPrimes(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	var count int
	for i := 2; i < n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

func isPrime(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
