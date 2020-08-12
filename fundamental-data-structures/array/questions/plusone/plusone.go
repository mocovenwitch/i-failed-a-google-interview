package plusone

import "fmt"

// PlusOne is the best
func PlusOne(digits []int) []int {
	// skip empty slice
	len := len(digits)

	if len == 0 {
		return digits
	}

	var result []int = digits
	var carry int = 1
	for i := len - 1; i >= 0; i-- {
		value := digits[i] + carry
		if value >= 10 {
			carry = 1
			result[i] = value % 10
		} else {
			carry = 0
			result[i] = value
		}
	}
	if carry > 0 {
		// go specific way of handle this
		result = append(result, 0)
		copy(result[2:], result[1:])
		result[0] = 1
		fmt.Println(result)
	}
	return result
}
