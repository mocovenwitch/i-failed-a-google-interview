package reversestring

import "fmt"

// ReverseString is O(n)
func ReverseString(s []byte) {

	r := len(s) - 1
	for l := 0; l < r; l++ {
		t := s[l]
		s[l] = s[r]
		s[r] = t
		r--
	}
	fmt.Println(s)
}

// ReverseString1 use Golang multi-valued assignment
// what's the strategy underneeth? the memory is observed similar to function above, and slower.
func ReverseString1(s []byte) {
	r := len(s) - 1
	for l := 0; l < r; l++ {
		s[l], s[r] = s[r], s[l] // exchange s[l] and s[r]
		r--
	}
	fmt.Println(s)
}

// ReverseString2 use XOR algorithm
// The result is the same as the #0 approach. I was expecting it can save some memory...
func ReverseString2(s []byte) {
	r := len(s) - 1
	for l := 0; l < r; l++ {
		s[l] = s[l] ^ s[r]
		s[r] = s[l] ^ s[r]
		s[l] = s[l] ^ s[r]
		r--
	}
	fmt.Println(s)
}
