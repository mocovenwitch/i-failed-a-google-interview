package questions

import (
	"fmt"
	"testing"
)

func TestCountPrimes(test *testing.T) {
	// n := 10000 //e:1229

	n := 499979

	// n := 10 //e:1229
	fmt.Printf("%d has %d primes", n, countPrimes(n))
}
