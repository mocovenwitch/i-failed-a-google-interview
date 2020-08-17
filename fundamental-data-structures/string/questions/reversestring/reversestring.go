package reversestring

import "fmt"

// ReverseString is O(n)
func ReverseString(s []byte) {

	// only process when size is larger than 2
	size := len(s)
	r := size - 1
	if size > 1 {
		for l := 0; l < r; l++ {
			t := s[l]
			s[l] = s[r]
			s[r] = t
			r--
		}
	}
	fmt.Println(s)
}
