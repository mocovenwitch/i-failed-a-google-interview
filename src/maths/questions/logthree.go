package questions

import (
	"math"
)

// IsPowerOfThree is
// log3 x = log10 x / log 10 3
//
// Result:
// Runtime: 16 ms, faster than 95.20% of Go online submissions for Power of Three.
// Memory Usage: 6 MB, less than 92.80% of Go online submissions for Power of Three.
func IsPowerOfThree(n int) bool {

	if n == 0 {
		return false
	}

	a := math.Log10(float64(n))
	// fmt.Println(a)

	b := math.Log10(3.0)
	// fmt.Println(b)

	c := a / b
	// fmt.Println(c)

	// handle float64 precision
	result := math.Round(c*100000000000000) / 100000000000000
	// result := float64(int(c*100)) / 100
	// fmt.Println(result)

	return math.Mod(float64(result), 1.0) == 0
	// return math.Mod(result, 1.0) == 0
}
