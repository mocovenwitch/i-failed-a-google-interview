package questions

import (
	"fmt"
	"math"
)

// IsPowerOfThree is
// log3 x = log10 x / log 10 3
//
// Result (not ideal):
// Runtime: 32 ms, faster than 41.60% of Go online submissions for Power of Three.
// Memory Usage: 6.1 MB, less than 12.80% of Go online submissions for Power of Three.
func IsPowerOfThree(n int) bool {
	a := math.Log10(float64(n))
	fmt.Println(a)

	b := math.Log10(3.0)
	fmt.Println(b)

	c := a / b
	fmt.Println(c)

	// handle float64 precision
	result := math.Round(c*100000000000000) / 100000000000000
	fmt.Println(result)

	return math.Mod(float64(result), 1.0) == 0
}
