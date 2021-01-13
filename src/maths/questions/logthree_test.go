package questions

import (
	"fmt"
	"testing"
)

func TestLogThree(test *testing.T) {
	x := 531440
	// x = 9
	// x = 4782968
	if IsPowerOfThree(x) {
		fmt.Printf("%d is power of 3", x)
	} else {
		fmt.Printf("%d is NOT power of 3", x)
	}
}
