package anagram

import (
	"fmt"
	"testing"
)

func TestAnagram(t *testing.T) {
	case1a := "anagram"
	case1b := "nagaram"

	fmt.Println(isAnagram(case1a, case1b))

	case2a := "rat"
	case2b := "car"

	fmt.Println(isAnagram(case2a, case2b))
}
